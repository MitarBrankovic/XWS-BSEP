package startup

import (
	connection "dislinkt/common/proto/connection_service"
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
	messageStore := server.initMessageStore(mongoClient)

	connectionService := server.initConnectionService(connectionStore)
	messageService := server.initMessageService(messageStore)

	connectionHandler := server.initConnectionHandler(connectionService, messageService)

	server.startGrpcServer(connectionHandler, jwtManager)
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

func (server *Server) initMessageStore(client *mongo.Client) domain.MessageStore {
	store := persistence.NewMessageMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Message := range messages {
		_, err := store.Create(Message)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initConnectionService(store domain.ConnectionStore) *application.ConnectionService {
	return application.NewConnectionService(store)
}

func (server *Server) initMessageService(store domain.MessageStore) *application.MessageService {
	return application.NewMessageService(store)
}

func (server *Server) initConnectionHandler(service *application.ConnectionService, messageService *application.MessageService) *api.ConnectionHandler {
	return api.NewConnectionHandler(service, messageService)
}

func (server *Server) startGrpcServer(connectionHandler *api.ConnectionHandler, jwtManager *auth.JWTManager) {
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
	connection.RegisterConnectionServiceServer(grpcServer, connectionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
