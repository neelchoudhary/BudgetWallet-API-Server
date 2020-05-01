package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neelchoudhary/budgetwallet-api-server/config"
	"github.com/neelchoudhary/budgetwallet-api-server/services/auth"
	"github.com/neelchoudhary/budgetwallet-api-server/services/plaidfinances"
	"github.com/neelchoudhary/budgetwallet-api-server/services/userfinances"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server holds the dependencies for a HTTP server.
type Server struct {
	serverConfig         *config.ServerConfig
	jwtManager           *utils.JWTManager
	authService          *auth.AuthServiceServer
	userFinancesService  *userfinances.UserFinancesServiceServer
	plaidFinancesService *plaidfinances.PlaidFinancesServiceServer
}

// NewServer construct a new server with service dependencies
func NewServer(serverConfig *config.ServerConfig,
	jwtManager *utils.JWTManager,
	authService *auth.AuthServiceServer,
	userFinancesService *userfinances.UserFinancesServiceServer,
	plaidFinancesService *plaidfinances.PlaidFinancesServiceServer) *Server {
	return &Server{
		serverConfig:         serverConfig,
		jwtManager:           jwtManager,
		authService:          authService,
		userFinancesService:  userFinancesService,
		plaidFinancesService: plaidFinancesService,
	}
}

// RunGRPCServer Starts the gRPC server
func (s *Server) runGRPCServer() error {
	listen, err := net.Listen("tcp", s.serverConfig.Host+":"+s.serverConfig.APIServerPort)
	if err != nil {
		return err
	}

	// TLS setup
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		creds, err := credentials.NewServerTLSFromFile(s.serverConfig.TLSCertPath, s.serverConfig.TLSKeyPath)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	}

	// Add interceptor
	opts = append(opts, grpc.UnaryInterceptor(s.serverInterceptor))
	grpcServer := grpc.NewServer(opts...)

	// Register services
	auth.RegisterAuthServiceServer(grpcServer, *s.authService)
	userfinances.RegisterUserFinancesServiceServer(grpcServer, *s.userFinancesService)
	plaidfinances.RegisterPlaidFinancesServiceServer(grpcServer, *s.plaidFinancesService)

	// Start gRPC server
	log.Println("starting gRPC server...")
	return grpcServer.Serve(listen)
}

// RunHTTPServer Starts the HTTP server
func (s *Server) runHTTPServer() error {
	// Init router
	r := mux.NewRouter()

	// Webhooks
	//r.HandleFunc("/webhook/users/{user_id}", controllers.ReceiveWebhooks).Methods("POST")

	// Start server
	log.Println("starting http server...")
	return http.ListenAndServe(":"+s.serverConfig.WebhookServerPort, r)
}

// Authorization unary interceptor function to handle authorize per RPC call
func (s *Server) serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authorize when signing up for logging in
	if info.FullMethod != "/auth.AuthService/Signup" && info.FullMethod != "/auth.AuthService/Login" {
		userID, err := s.jwtManager.AuthorizeToken(ctx)
		if err != nil {
			return nil, err
		}

		// Calls the handler with new context
		h, err := handler(utils.PassUserIDMetadata(ctx, userID), req)
		return h, err
	}

	// Calls the handler
	h, err := handler(ctx, req)
	return h, err
}
