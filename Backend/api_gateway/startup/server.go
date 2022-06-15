package startup

import (
	"context"
	cfg "dislinkt/api_gateway/startup/config"
	"dislinkt/common/https"
	"dislinkt/common/loggers"
	_ "dislinkt/common/loggers"
	connectionPb "dislinkt/common/proto/connection_service"
	offerPb "dislinkt/common/proto/offer_service"
	postPb "dislinkt/common/proto/post_service"
	userPb "dislinkt/common/proto/user_service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"strings"
)

var errorLog = loggers.NewErrorLogger()
var customLog = loggers.NewCustomLogger()

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()

		customLog.WithFields(logrus.Fields{
			"method":     r.Method,
			"url":        r.URL.String(),
			"origin":     r.Header.Get("Origin"),
			"user-agent": r.Header.Get("User-Agent"),
		}).Info("CORS filter")

		h.Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", strings.Join(
				[]string{
					http.MethodOptions,
					http.MethodGet,
					http.MethodPut,
					http.MethodHead,
					http.MethodPost,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodTrace,
				}, ", ",
			))

			h.Set("Access-Control-Allow-Headers", strings.Join(
				[]string{
					"Access-Control-Allow-Headers",
					"Origin",
					"X-Requested-With",
					"Content-Type",
					"Accept",
					"Authorization",
					"Location",
				}, ", ",
			))

			return
		}

		next.ServeHTTP(w, r)
	})
}

func (server *Server) initHandlers() {
	tlsCredentials, err := https.LoadTLSClientCredentials()
	if err != nil {
		log.Fatalf("failed to load tls credentials: %v", err)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(tlsCredentials)}
	//opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	userEndpoint := fmt.Sprintf("%s:%s", server.config.UserHost, server.config.UserPort)
	err = userPb.RegisterUserServiceHandlerFromEndpoint(context.TODO(), server.mux, userEndpoint, opts)
	if err != nil {
		panic(err)
	}

	postEndpoint := fmt.Sprintf("%s:%s", server.config.PostHost, server.config.PostPort)
	errPost := postPb.RegisterPostServiceHandlerFromEndpoint(context.TODO(), server.mux, postEndpoint, opts)
	if errPost != nil {
		panic(errPost)
	}

	connectionEndpoint := fmt.Sprintf("%s:%s", server.config.ConnectionHost, server.config.ConnectionPort)
	errConnection := connectionPb.RegisterConnectionServiceHandlerFromEndpoint(context.TODO(), server.mux, connectionEndpoint, opts)
	if errConnection != nil {
		panic(errConnection)
	}

	offerEndpoint := fmt.Sprintf("%s:%s", server.config.OfferHost, server.config.OfferPort)
	errOffer := offerPb.RegisterOfferServiceHandlerFromEndpoint(context.TODO(), server.mux, offerEndpoint, opts)
	if errOffer != nil {
		panic(errOffer)
	}
}

func (server *Server) Start() {
	serverCertFile := getCertPath() + "cert/server-cert.pem"
	serverKeyFile := getCertPath() + "cert/server-key.pem"
	errorLog.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%s", server.config.Port), serverCertFile, serverKeyFile, cors(server.mux)))
}

func getCertPath() string {
	if os.Getenv("OS_ENV") != "docker" {
		return "../"
	}
	return ""
}
