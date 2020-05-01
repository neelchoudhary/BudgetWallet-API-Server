package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service ...
type Service struct {
	userRepo   models.UserRepository
	jwtManager *utils.JWTManager
}

// NewAuthServiceServer contructor to assign repo
func NewAuthServiceServer(repository *models.UserRepository, jwtManager *utils.JWTManager) AuthServiceServer {
	return &Service{userRepo: *repository, jwtManager: jwtManager}
}

// Signup ...
func (s *Service) Signup(ctx context.Context, req *SignupRequest) (*SignupResponse, error) {
	// Check if email already exists in db
	signUpUser := req.GetSignUpUser()
	if signUpUser.GetFullname() == "" || signUpUser.GetEmail() == "" || signUpUser.GetPassword() == "" {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("All fields are required"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpUser.GetPassword()), 10)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to hash password: %s", err.Error()))
	}
	//uniqueID := uuid.NewV4()
	newUser := User{
		Email:     signUpUser.Email,
		Password:  string(hash),
		Fullname:  signUpUser.Fullname,
		CreatedOn: time.Now().Format("2006-01-02T15:04:05"),
	}
	err = s.userRepo.CreateUser(*signUpPbToData(newUser))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Repo error creating user: %s", err.Error()))
	}

	res := &SignupResponse{
		Success: true,
	}

	return res, nil
}

// Login ...
func (s *Service) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	loginUser := req.GetLoginUser()
	email := loginUser.GetEmail()
	password := loginUser.GetPassword()
	userToLogIn, err := s.userRepo.GetUserByEmail(email)
	tokenString, err := userToLogIn.Login(password, s.jwtManager)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("User error logging in: %s", err.Error()))
	}

	res := &LoginResponse{
		Success: true,
		Token:   tokenString,
	}
	return res, nil
}

func signUpPbToData(data User) *models.User {
	return &models.User{
		Email:     data.Email,
		Password:  data.Password,
		FullName:  data.Fullname,
		CreatedOn: data.CreatedOn,
	}
}
