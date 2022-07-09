package startup

import (
	post "dislinkt/common/proto/post_service"
	saga "dislinkt/common/saga/messaging"
	"dislinkt/common/saga/messaging/nats"
	"dislinkt/post_service/application"
	"dislinkt/post_service/domain"
	"dislinkt/post_service/infrastructure/api"
	"dislinkt/post_service/infrastructure/persistence"
	"dislinkt/post_service/startup/config"
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
	QueueGroup = "post_service"
)

func accessibleRoles() map[string][]string {

	const postServicePath = "/post.PostService/"

	return map[string][]string{}
}

func (server *Server) Start() {

	jwtManager := auth.NewJWTManager("secretKey", 15*time.Minute)

	mongoClient := server.initMongoClient()
	postStore := server.initPostStore(mongoClient)
	postService := server.initPostService(postStore)
	reactionService := server.initReactionService(postStore)
	commentService := server.initCommentService(postStore)

	/*userClient, err := clients.NewUserClient(fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort))
	if err != nil {
		log.Fatal(err)
	}*/

	commandSubscriber := server.initSubscriber(server.config.UpdateUserCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.UpdateUserReplySubject)
	server.initUpdateUserHandler(postService, replyPublisher, commandSubscriber)

	postHandler := server.initPostHandler(postService, reactionService, commentService)
	server.startGrpcServer(postHandler, jwtManager)
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

func (server *Server) initUpdateUserHandler(service *application.PostService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUpdateUserCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initPostStore(client *mongo.Client) domain.PostStore {
	store := persistence.NewPostMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Post := range posts {
		err := store.Create(Post)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initPostService(store domain.PostStore) *application.PostService {
	return application.NewPostService(store)
}

func (server *Server) initReactionService(store domain.PostStore) *application.ReactionService {
	return application.NewReactionService(store)
}

func (server *Server) initCommentService(store domain.PostStore) *application.CommentService {
	return application.NewCommentService(store)
}

func (server *Server) initPostHandler(service *application.PostService, reactionService *application.ReactionService, commentService *application.CommentService) *api.PostHandler {
	return api.NewPostHandler(service, reactionService, commentService)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler, jwtManager *auth.JWTManager) {
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
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
