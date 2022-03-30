package main

import (
	"github.com/shinhagunn/shop-email/config"
	"github.com/shinhagunn/shop-email/services"
)

type Service interface {
	Process()
}

func main() {
	config.InitMongoDB()
	service := services.NewSendEmail()

	service.Process()
}
