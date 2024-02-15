package service

// import (
// 	"github.com/dalecosta1/sinaloa-api/helpers"
// )

type HaProxyService interface {
	RenewCerts() error
}

func RenewCerts() bool {
	return true
}
