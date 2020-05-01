package utils

import (
	"context"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// TokenAuth token auth for client side
type TokenAuth struct {
	tokenString string
}

// GetRequestMetadata ...
func (t TokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.tokenString,
	}, nil
}

// RequireTransportSecurity require TLS
func (TokenAuth) RequireTransportSecurity() bool {
	return true
}

// GetTokenAuth gets auth token to pass into each rpc
func GetTokenAuth(data string) TokenAuth {
	return TokenAuth{tokenString: data}
}

// JWTManager struct
type JWTManager struct {
	expiryMinutes int
	secret        string
}

// NewJWTManager creates new JWTManager
func NewJWTManager(jwtExpiryMinutes int, jwtSecret string) *JWTManager {
	return &JWTManager{expiryMinutes: jwtExpiryMinutes, secret: jwtSecret}
}

type claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// CreateToken creates a new token with a claim, returns grpc errors
func (j *JWTManager) CreateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(j.expiryMinutes) * time.Minute)
	claims := &claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Login and get the encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", status.Errorf(codes.Internal, "Failed to sign token: "+err.Error())

	}
	return tokenString, nil
}

// AuthorizeToken authorizes the token received from metadata, returns grpc errors
func (j *JWTManager) AuthorizeToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	token := authHeader[0]

	const prefix = "Bearer "
	if !strings.HasPrefix(token, prefix) {
		return "", status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
	}

	token = strings.TrimPrefix(token, prefix)

	// validateToken function validates the token
	userID, err := j.validateToken(token, j.secret)

	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, err.Error())
	}
	return userID, nil
}

// ValidateToken validates the provided token string
func (j *JWTManager) validateToken(tokenString string, jwtSecret string) (string, error) {
	claims := &claims{}
	// Parse the JWT string and store the result in claims.
	// This method will return an error if the token is invalid/expired,
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
