package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/controllers"
	"github.com/taivama/golang-training/proto"
	"github.com/taivama/golang-training/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunClient() {
	tlsCredentials, err := LoadCACredentials()
	if err != nil {
		log.Fatalf("cannot load CA credentials: %s", err.Error())
	}

	connection, err := grpc.Dial(":9090", grpc.WithTransportCredentials(tlsCredentials))
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
	//routes.AddSecuredRoutes(g, userCtrl)
}

func InitProducts(c *grpc.ClientConn, g *gin.Engine) {
	products := proto.NewProductServiceClient(c)
	productCtrl := controllers.InitProductController(products)
	routes.AddProductRoutes(g, productCtrl)
}

func LoadCACredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := os.ReadFile("certs/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}
