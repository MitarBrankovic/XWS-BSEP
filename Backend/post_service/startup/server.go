package startup

import (
	"dislinkt/common/clients"
	post "dislinkt/common/proto/post_service"
	pbUser "dislinkt/common/proto/user_service"
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

	userClient, err := clients.NewUserClient(fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort))
	if err != nil {
		log.Fatal(err)
	}

	postHandler := server.initPostHandler(postService, reactionService, commentService, userClient)
	server.startGrpcServer(postHandler, jwtManager)
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

func (server *Server) initPostHandler(service *application.PostService, reactionService *application.ReactionService, commentService *application.CommentService, userClient pbUser.UserServiceClient) *api.PostHandler {
	return api.NewPostHandler(service, reactionService, commentService, userClient)
}

func (server *Server) startGrpcServer(postHandler *api.PostHandler, jwtManager *auth.JWTManager) {
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
	post.RegisterPostServiceServer(grpcServer, postHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
