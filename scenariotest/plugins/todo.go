package main

import (
	todov1 "github.com/matumoto1234/connect-go-todo/gen/todo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn

func newConn(target string) *grpc.ClientConn {
	if conn != nil {
		return conn
	}

	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	return conn
}

func NewToDoServiceServiceClient() todov1.ToDoServiceClient {
	conn := newConn("localhost:8080")
	return todov1.NewToDoServiceClient(conn)
}
