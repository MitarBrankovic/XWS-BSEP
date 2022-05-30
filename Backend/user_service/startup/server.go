package startup

import (
	"context"
	"dislinkt/common/clients"
	pbPost "dislinkt/common/proto/post_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/user_service/application"
	"dislinkt/user_service/auth"
	"dislinkt/user_service/domain"
	"dislinkt/user_service/infrastructure/api"
	"dislinkt/user_service/infrastructure/persistence"
	"dislinkt/user_service/startup/config"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"
)

type Server struct {
	config     *config.Config
	userStore  domain.UserStore
	jwtManager *auth.JWTManager
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {

	const userServicePath = "/user.UserService/"

	return map[string][]string{
		userServicePath + "Get": {"user", "admin"},
		//userServicePath + "FindByUsername": {"user", "admin"},
		userServicePath + "Create":         {"admin"},
		userServicePath + "Update":         {"admin", "user"},
		userServicePath + "Delete":         {"admin"},
		userServicePath + "ChangePassword": {"user"},
	}
}

func (server *Server) Start() {

	jwtManager := auth.NewJWTManager("secretKey", 15*time.Minute)

	userClient, err := clients.NewUserClient(fmt.Sprintf("%s:%s", server.config.Host, server.config.Port))
	if err != nil {
		log.Fatal(err)
	}

	postClient, err := clients.NewPostClient(fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort))
	if err != nil {
		log.Fatal(err)
	}

	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)
	userService := server.initUserService(userStore)
	mailService := server.initMailService()
	userHandler := server.initUserHandler(userService, mailService, jwtManager, userClient, postClient)
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

func (server *Server) initMailService() *application.MailService {
	return application.NewMailService()
}

func (server *Server) initUserHandler(service *application.UserService, mailService *application.MailService, jwtManager *auth.JWTManager,
	userClient pbUser.UserServiceClient, postClient pbPost.PostServiceClient) *api.UserHandler {
	return api.NewUserHandler(service, mailService, jwtManager, userClient, postClient)
}

func (server *Server) Login(ctx context.Context, req *pbUser.LoginRequest) (*pbUser.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pbUser.LoginResponse{AccessToken: token}
	return res, nil
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
	pbUser.RegisterUserServiceServer(grpcServer, userHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
