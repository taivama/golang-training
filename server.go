package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/taivama/golang-training/proto"
	"github.com/taivama/golang-training/server"
	"github.com/taivama/golang-training/services"
	"github.com/taivama/golang-training/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func RunServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := utils.ConnectDB(ctx)
	if err != nil {
		log.Fatal("connection to database failed")
	}
	defer c.Disconnect(ctx)

	db := c.Database("Matti")
	userSvc := services.InitUserService(db.Collection("Users"))
	us := server.InitUserServer(userSvc)
	prodSvc := services.InitProductService(db.Collection("Products"))
	ps := server.InitProductServer(prodSvc)

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("listening port tcp 9090 failed: %s", err.Error())
		return
	}
	defer listen.Close()

	tlsCredentials, err := LoadServerCredentials()
	if err != nil {
		log.Fatalf("cannot load TLS credentials: %s", err.Error())
	}

	s := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	proto.RegisterUserServiceServer(s, us)
	proto.RegisterProductServiceServer(s, ps)
	fmt.Println("Server listening on port tcp 9090")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("grpc server failed: %s", err.Error())
	}
}

func LoadServerCredentials() (credentials.TransportCredentials, error) {
	serverCert, err := tls.LoadX509KeyPair("certs/localhost/cert.pem", "certs/localhost/key.pem")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}
