package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-echo-test/common/config"
	"go-echo-test/common/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

var localStorage *model.TodoList

func init() {
	localStorage = new(model.TodoList)
	localStorage.List = make([]*model.Todo, 0)
}

type TodosServer struct{}

func (TodosServer) Register(ctx context.Context, param *model.Todo) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, param)
	log.Println("Registering Todo", param.String())
	return new(empty.Empty), nil
}

func (TodosServer) List(ctx context.Context, void *empty.Empty) (*model.TodoList, error) {
	return localStorage, nil
}

func (TodosServer) Remove(ctx context.Context, param *model.Todo) (*empty.Empty, error) {
	log.Println("Removing Todo", param.String())
	index := -1

	for i, t := range localStorage.List {
		if t.Id == param.Id {
			index = i
		}
	}

	if index == -1 {
		log.Println("NOT FOUND !")
	}

	localStorage.List = append(localStorage.List[:index], localStorage.List[index+1:]...)

	return new(empty.Empty), nil
}

func (TodosServer) Edit(ctx context.Context, param *model.Todo) (*empty.Empty, error) {
	log.Println("Edit Todo", param.String())

	index := -1
	for i, t := range localStorage.List {
		if t.Id == param.Id {
			index = i
		}
	}

	if index == -1 {
		log.Println("NOT FOUND !")
	}

	localStorage.List[index].Name = param.Name

	return new(empty.Empty), nil
}

func main() {
	srv := grpc.NewServer()
	var todoSrv TodosServer
	model.RegisterTodosServer(srv, todoSrv)

	log.Println("Starting RPC server at", config.SERVICE_TODO_PORT)

	l, err := net.Listen("tcp", config.SERVICE_TODO_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_TODO_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
