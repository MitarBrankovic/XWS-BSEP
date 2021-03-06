package startup

import (
	connection "dislinkt/common/proto/connection_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/connection_service/application"
	"dislinkt/connection_service/domain"
	"dislinkt/connection_service/infrastructure/api"
	"dislinkt/connection_service/infrastructure/persistence"
	"dislinkt/connection_service/startup/config"
	"dislinkt/user_service/auth"
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

const (
	QueueGroup = "connection_service"
)

func accessibleRoles() map[string][]string {

	const connectionServicePath = "/connection.ConnectionService/"

	return map[string][]string{
		connectionServicePath + "Get": {"user"},
	}
}

func (server *Server) Start() {

	jwtManager := auth.NewJWTManager("secretKey", 15*time.Minute)

	mongoClient := server.initMongoClient()
	connectionStore := server.initConnectionStore(mongoClient)

	connectionService := server.initConnectionService(connectionStore)

	commandSubscriber := server.initSubscriber(server.config.UpdateUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateUserReplySubject)
	server.initUpdateUserHandler(connectionService, replyPublisher, commandSubscriber)

	connectionHandler := server.initConnectionHandler(connectionService)

	server.startGrpcServer(connectionHandler, jwtManager)
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

func (server *Server) initUpdateUserHandler(service *application.ConnectionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.ConnectionDBHost, server.config.ConnectionDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initConnectionStore(client *mongo.Client) domain.ConnectionStore {
	store := persistence.NewConnectionMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Connection := range connections {
		_, err := store.Create(Connection)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler, jwtManager *auth.JWTManager) {
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	//tlsCredentials, err := https.LoadTLSServerCredentials()
	/*if err != nil {
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
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
