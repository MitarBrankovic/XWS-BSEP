package startup

import (
	"dislinkt/common/auth"
	user "dislinkt/common/proto/user_service"
	"dislinkt/user_service/application"
	"dislinkt/user_service/domain"
	"dislinkt/user_service/infrastructure/api"
	"dislinkt/user_service/infrastructure/persistence"
	"dislinkt/user_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {

	const userServicePath = "/user.UserService/"

	return map[string][]string{}
}

func (server *Server) Start() {

	jwtManager := auth.NewJWTManager("secretKey", 15*time.Minute)

	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)
	userService := server.initUserService(userStore)
	userHandler := server.initUserHandler(userService)
	server.startGrpcServer(userHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, User := range users {
		err := store.Create(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(userHandler *api.UserHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	reflection.Register(grpcServer)
	user.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
