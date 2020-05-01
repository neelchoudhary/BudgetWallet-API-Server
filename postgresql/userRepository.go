package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository returns a new instance of a postgresql user repository.
func NewUserRepository(db *sql.DB) models.UserRepository {
	return &userRepository{db: db}
}

// GetUserByEmail tries to find user with email
func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	row := r.db.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Password, &user.CreatedOn)
	if err == sql.ErrNoRows {
		return user, nil
	} else if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser create new user for signup
func (r *userRepository) CreateUser(user models.User) error {
	var userID string
	statement := "INSERT INTO users (name, email, password, created_on) VALUES ($1, $2, $3, $4) RETURNING id;"
	err := r.db.QueryRow(statement, user.FullName, user.Email, user.Password, user.CreatedOn).Scan(&userID)

	if err != nil {
		return err
	}

	return nil
}
