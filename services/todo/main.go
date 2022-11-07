package main

import (
	"context"
	"flag"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-echo-test/common/config"
	"go-echo-test/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

var localStorage *model.TodoList

func init() {
	localStorage = new(model.TodoList)
	localStorage.Data = make([]*model.Todo, 0)
}

type TodosServer struct {
	model.UnimplementedTodosServer
}

func (TodosServer) Register(ctx context.Context, param *model.Todo) (*model.MutationResponse, error) {
	localStorage.Data = append(localStorage.Data, param)
	log.Println("Registering Todo", param.String())

	msg := param.GetId() + " successfully appended"
	return &model.MutationResponse{Success: msg}, nil
}

func (TodosServer) List(ctx context.Context, void *empty.Empty) (*model.TodoList, error) {
	return localStorage, nil
}

func (TodosServer) GetById(ctx context.Context, param *model.GetByIDRequest) (*model.GetByIDResponse, error) {
	log.Println("Getting Todo ", param.Id)
	index := -1

	for i, t := range localStorage.Data {
		if t.Id == param.Id {
			index = i
		}
	}

	if index == -1 {
		log.Println("NOT FOUND !")
		return &model.GetByIDResponse{Data: nil}, nil
	}

	return &model.GetByIDResponse{Data: localStorage.Data[index]}, nil

}

func (TodosServer) Remove(ctx context.Context, param *model.Todo) (*model.MutationResponse, error) {
	log.Println("Removing Todo", param.String())
	index := -1

	for i, t := range localStorage.Data {
		if t.Id == param.Id {
			index = i
		}
	}

	if index == -1 {
		log.Println("NOT FOUND !")
	}

	localStorage.Data = append(localStorage.Data[:index], localStorage.Data[index+1:]...)

	msg := param.GetId() + " successfully removed"
	return &model.MutationResponse{Success: msg}, nil
}

func (TodosServer) Edit(ctx context.Context, param *model.Todo) (*model.MutationResponse, error) {
	log.Println("Edit Todo", param.String())

	index := -1
	for i, t := range localStorage.Data {
		if t.Id == param.Id {
			index = i
		}
	}

	if index == -1 {
		log.Println("NOT FOUND !")
	}

	localStorage.Data[index].Name = param.Name

	msg := param.GetId() + " successfully edited"
	return &model.MutationResponse{Success: msg}, nil
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

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_TODO_PORT, "gRPC server endpoint")
		_ = model.RegisterTodosHandlerFromEndpoint(context.Background(), mux, *grpcServerEndpoint, opts)
		log.Println("Starting Todo Server HTTP at 9001 ")

		http.ListenAndServe(":9001", mux)
	}()

	log.Fatal(srv.Serve(l))
}
