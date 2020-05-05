package models

import (
	"fmt"

	"github.com/neelchoudhary/budgetwallet-api-server/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User ...
type User struct {
	ID        string `json:"id"`
	FullName  string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
}

// Login ...
func (u *User) Login(password string, jwtManager *utils.JWTManager) (string, error) {
	// Check if user exists in db
	if u.ID != "" {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err != nil {
			// Error, incorrect password
			return "", status.Errorf(codes.PermissionDenied, fmt.Sprintf("Invalid Login Credientials: %s", err.Error()))
		}

		tokenString, err := jwtManager.CreateToken(u.ID)
		if err != nil {
			return "", utils.InternalServerError
		}

		return tokenString, nil
	}
	// User with the given email does not exist
	return "", status.Errorf(codes.PermissionDenied, fmt.Sprintf("Invalid Login Credientials"))
}

// UserRepository ...
type UserRepository interface {
	GetUserByEmail(email string) (User, error)
	CreateUser(user User) error
}
