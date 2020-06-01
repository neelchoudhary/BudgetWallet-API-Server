package main

import (
	"context"
	"log"
	"net"

	"github.com/neelchoudhary/budgetwallet-api-server/services/dataprocessing"
	"github.com/neelchoudhary/budgetwallet-api-server/services/webhooks"

	"github.com/neelchoudhary/budgetwallet-api-server/services/financialcategories"

	"github.com/neelchoudhary/budgetwallet-api-server/config"
	"github.com/neelchoudhary/budgetwallet-api-server/services/auth"
	"github.com/neelchoudhary/budgetwallet-api-server/services/plaidfinances"
	"github.com/neelchoudhary/budgetwallet-api-server/services/userfinances"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server holds the dependencies for a gRPC server.
type Server struct {
	serverConfig               *config.ServerConfig
	jwtManager                 *utils.JWTManager
	authService                *auth.AuthServiceServer
	userFinancesService        *userfinances.UserFinancesServiceServer
	plaidFinancesService       *plaidfinances.PlaidFinancesServiceServer
	financialCategoriesService *financialcategories.FinancialCategoryServiceServer
	dataProcessingService      *dataprocessing.DataProcessingServiceServer
	webhooksService            *webhooks.WebhooksServiceServer
}

// NewServer construct a new server with service dependencies
func NewServer(
	serverConfig *config.ServerConfig,
	jwtManager *utils.JWTManager,
	authService *auth.AuthServiceServer,
	userFinancesService *userfinances.UserFinancesServiceServer,
	plaidFinancesService *plaidfinances.PlaidFinancesServiceServer,
	financialCategoriesService *financialcategories.FinancialCategoryServiceServer,
	dataProcessingService *dataprocessing.DataProcessingServiceServer,
	webhooksService *webhooks.WebhooksServiceServer,
) *Server {
	return &Server{
		serverConfig:               serverConfig,
		jwtManager:                 jwtManager,
		authService:                authService,
		userFinancesService:        userFinancesService,
		plaidFinancesService:       plaidFinancesService,
		financialCategoriesService: financialCategoriesService,
		dataProcessingService:      dataProcessingService,
		webhooksService:            webhooksService,
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
	financialcategories.RegisterFinancialCategoryServiceServer(grpcServer, *s.financialCategoriesService)
	dataprocessing.RegisterDataProcessingServiceServer(grpcServer, *s.dataProcessingService)
	webhooks.RegisterWebhooksServiceServer(grpcServer, *s.webhooksService)

	// Start gRPC server
	log.Println("starting gRPC server...")
	return grpcServer.Serve(listen)
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
