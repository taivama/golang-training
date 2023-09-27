package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/controllers"
	"github.com/taivama/golang-training/routes"
	"github.com/taivama/golang-training/services"
	"github.com/taivama/golang-training/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := utils.ConnectDB(ctx)
	if err != nil {
		log.Fatal("connection to database failed")
	}
	defer c.Disconnect(ctx)

	server := gin.Default()
	db := c.Database("Matti")
	InitUsers(db, server)
	InitProducts(db, server)
	server.Run(":8080")
}

func InitUsers(db *mongo.Database, g *gin.Engine) {
	users := db.Collection("Users")
	userSvc := services.InitUserService(users)
	userCtrl := controllers.InitUserController(userSvc)
	routes.AddUnSecuredRoutes(g, userCtrl)
	routes.AddSecuredRoutes(g, userCtrl)
}

func InitProducts(db *mongo.Database, g *gin.Engine) {
	products := db.Collection("Products")
	productSvc := services.InitProductService(products)
	productCtrl := controllers.InitProductController(productSvc)
	routes.AddProductRoutes(g, productCtrl)
}
