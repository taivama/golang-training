package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/taivama/golang-training/proto"
	"github.com/taivama/golang-training/server"
	"github.com/taivama/golang-training/services"
	"github.com/taivama/golang-training/utils"
	"google.golang.org/grpc"
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

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, us)
	proto.RegisterProductServiceServer(s, ps)
	fmt.Println("Server listening on port tcp 9090")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("grpc server failed: %s", err.Error())
	}
}
