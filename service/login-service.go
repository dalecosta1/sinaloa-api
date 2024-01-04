package service

import (
	"os"
	"fmt"

	"github.com/dalecosta1/sinaloa-api/helpers/config"
)

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	// Set basic auth username and password
	return &loginService{
		authorizedUsername: config.AppConfig.BASIC_AUTH_USER,
		authorizedPassword: config.AppConfig.BASIC_AUTH_PASSWORD,
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
