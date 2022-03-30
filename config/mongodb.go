package config

import (
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/shinhagunn/shop-auth/config/collection"
	"github.com/shinhagunn/shop-auth/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mgm.SetDefaultConfig(nil, "authDB", options.Client().ApplyURI("mongodb://root:123456@localhost:27017"))

	log.Println("Connected to authDB!")

	collection.Code = mgm.Coll(&models.Code{})
}
