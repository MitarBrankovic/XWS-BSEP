package startup

import (
	"context"
	pbUser "dislinkt/common/proto/user_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
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

const (
	QueueGroup = "user_service"
)

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

	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	commandPublisher := server.initPublisher(server.config.UpdateUserCommandSubject)
	replySubscriber := server.initSubscriber(server.config.UpdateUserReplySubject, QueueGroup)
	updateUserOrchestrator := server.initUpdateUserOrchestrator(commandPublisher, replySubscriber)

	userService := server.initUserService(userStore, updateUserOrchestrator)
	mailService := server.initMailService()

	commandSubscriber := server.initSubscriber(server.config.UpdateUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateUserReplySubject)
	server.initUpdateUserHandler(userService, replyPublisher, commandSubscriber)

	userHandler := server.initUserHandler(userService, mailService, jwtManager)
	server.startGrpcServer(userHandler, jwtManager)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initUpdateUserOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UpdateUserOrchestrator {
	orchestrator, err := application.NewUpdateUserOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initUpdateUserHandler(service *application.UserService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
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
	err = store.DeleteAllBlocks()
	if err != nil {
		return nil
	}
	for _, User := range users {
		err := store.Create(User)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, Block := range blocks {
		err := store.Block(Block)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initUserService(store domain.UserStore, orchestartor *application.UpdateUserOrchestrator) *application.UserService {
	return application.NewUserService(store, orchestartor)
}

func (server *Server) initMailService() *application.MailService {
	return application.NewMailService()
}

func (server *Server) initUserHandler(service *application.UserService, mailService *application.MailService, jwtManager *auth.JWTManager) *api.UserHandler {
	return api.NewUserHandler(service, mailService, jwtManager)
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
	/*tlsCredentials, err := https.LoadTLSServerCredentials()
	if err != nil {
		panic("cannot load TLS credentials: %w")
	}*/
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
