package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"go-echo-test/common/config"
	"go-echo-test/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func serviceTodo() model.TodosClient {
	port := config.SERVICE_TODO_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewTodosClient(conn)
}

func main() {
	todoSvc := serviceTodo()
	ctx := context.Background()

	todo1 := &model.Todo{
		Id:   "t001",
		Name: "Test gan 1",
	}
	log.Printf("Hit todoSvc.Register %+v\n", todo1)
	_, _ = todoSvc.Register(ctx, todo1)

	todo2 := &model.Todo{
		Id:   "t002",
		Name: "Test gan 2",
	}
	log.Printf("Hit todoSvc.Register %+v\n", todo2)
	_, _ = todoSvc.Register(ctx, todo2)

	log.Printf("Hit todoSvc.List\n")
	todo, _ := todoSvc.List(ctx, new(empty.Empty))
	log.Printf("List Todo %+v\n", todo.GetData())

	log.Printf("Hit todoSvc.Remove %+v\n", todo1)
	_, _ = todoSvc.Remove(ctx, todo1)

	log.Printf("Hit todoSvc.List after remove\n")
	todoAfterDelete, _ := todoSvc.List(ctx, new(empty.Empty))
	log.Printf("List Todo after remove %+v\n", todoAfterDelete.GetData())

	log.Printf("Hit todoSvc.Edit %+v\n", todo2)
	todo2.Name = "Udah di edit deh !"
	_, _ = todoSvc.Edit(ctx, todo2)

	log.Printf("Hit todoSvc.List after edit\n")
	todoAfterEdit, _ := todoSvc.List(ctx, new(empty.Empty))
	log.Printf("List Todo after edit %+v\n", todoAfterEdit.GetData())
}
