package service

import (
	"github.com/dalecosta1/sinaloa-api/helpers"
)

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	// Load the config
	helpers.LoadConfig()
	// Set basic auth username and password
	return &loginService{
		authorizedUsername: helpers.AppConfig.BASIC_AUTH_USER,
		authorizedPassword: helpers.AppConfig.BASIC_AUTH_PASSWORD,
	}
}

func (service *loginService) Login(username string, password string) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
