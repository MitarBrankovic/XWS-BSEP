package startup

import (
	"dislinkt/common/clients"
	"dislinkt/common/https"
	pbOffer "dislinkt/common/proto/offer_service"
	pbUser "dislinkt/common/proto/user_service"
	"dislinkt/offer_service/application"
	"dislinkt/offer_service/domain"
	"dislinkt/offer_service/infrastructure/api"
	"dislinkt/offer_service/infrastructure/persistence"
	"dislinkt/offer_service/startup/config"
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
	config     *config.Config
	offerStore domain.OfferStore
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func accessibleRoles() map[string][]string {

	const offerServicePath = "/offer.OfferService/"

	return map[string][]string{}
}

func (server *Server) Start() {

	jwtManager := auth.NewJWTManager("secretKey", 15*time.Minute)
	mongoClient := server.initMongoClient()
	offerStore := server.initOfferStore(mongoClient)
	offerService := server.initOfferService(offerStore)

	userClient, err := clients.NewUserClient(fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort))
	if err != nil {
		log.Fatal(err)
	}

	offerHandler := server.initOfferHandler(offerService, userClient)
	server.startGrpcServer(offerHandler, jwtManager)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initOfferStore(client *mongo.Client) domain.OfferStore {
	store := persistence.NewOfferMongoDBStore(client)
	err := store.DeleteAll()
	if err != nil {
		return nil
	}
	for _, Offer := range offers {
		err := store.Create(Offer)
		if err != nil {
			log.Fatal(err)
		}
	}
	return store
}

func (server *Server) initOfferService(store domain.OfferStore) *application.OfferService {
	return application.NewOfferService(store)
}

func (server *Server) initOfferHandler(service *application.OfferService, userClient pbUser.UserServiceClient) *api.OfferHandler {
	return api.NewOfferHandler(service, userClient)
}

func (server *Server) startGrpcServer(offerHandler *api.OfferHandler, jwtManager *auth.JWTManager) {

	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	tlsCredentials, err := https.LoadTLSServerCredentials()
	if err != nil {
		panic("cannot load TLS credentials: %w")
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serverOptions := []grpc.ServerOption{
		grpc.Creds(tlsCredentials),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	reflection.Register(grpcServer)
	pbOffer.RegisterOfferServiceServer(grpcServer, offerHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
