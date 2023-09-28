package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/controllers"
	"github.com/taivama/golang-training/proto"
	"github.com/taivama/golang-training/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunClient() {
	connection, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connecting to grpc server failed: %s", err.Error())
	}
	defer connection.Close()

	server := gin.Default()
	InitUsers(connection, server)
	InitProducts(connection, server)
	server.Run(":8080")
}

func InitUsers(c *grpc.ClientConn, g *gin.Engine) {
	users := proto.NewUserServiceClient(c)
	userCtrl := controllers.InitUserController(users)
	routes.AddUnSecuredRoutes(g, userCtrl)
	routes.AddSecuredRoutes(g, userCtrl)
}

func InitProducts(c *grpc.ClientConn, g *gin.Engine) {
	products := proto.NewProductServiceClient(c)
	productCtrl := controllers.InitProductController(products)
	routes.AddProductRoutes(g, productCtrl)
}
