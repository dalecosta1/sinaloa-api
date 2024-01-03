package service

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[ERROR] Error loading .env file:", err)
	}
	// Set basic auth username and password
	return &loginService{
		authorizedUsername: os.Getenv("BASIC_AUTH_USER"),
		authorizedPassword: os.Getenv("BASIC_AUTH_PASSWORD"),
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
