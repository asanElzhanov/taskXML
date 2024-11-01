package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type DatabaseRepository struct {
	DB *sql.DB
}

type UserRepository interface {
	ValidateCredentials(login, password string) bool
	RetrieveAddress(login string) (string, error)
}

func NewDatabaseRepository(dbName string) *DatabaseRepository {
	connStr := fmt.Sprintf("user=postgres password=123456Aa dbname=%s sslmode=disable", dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return &DatabaseRepository{db}
}

func (r *DatabaseRepository) ValidateCredentials(login, password string) bool {
	var storedPassword string
	query := "SELECT pass FROM users WHERE login = $1"
	err := r.DB.QueryRow(query, login).Scan(&storedPassword)
	if err != nil {
		fmt.Println("Error while finding user: ", err)
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)) == nil
}

func (r *DatabaseRepository) RetrieveAddress(login string) (string, error) {
	var storedAddress string
	query := "SELECT address FROM users WHERE login = $1"
	err := r.DB.QueryRow(query, login).Scan(&storedAddress)
	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}
	return storedAddress, nil
}
