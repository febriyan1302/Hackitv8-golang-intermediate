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

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	userSvc := serviceUser()
	ctx := context.Background()

	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "for the horde",
		Gender:   model.UserGender_FEMALE,
	}
	log.Printf("Hit userSvc.Register %+v\n", user1)
	_, _ = userSvc.Register(ctx, user1)

	user2 := &model.User{
		Id:       "u002",
		Name:     "John Doe",
		Password: "for the horde",
		Gender:   model.UserGender_MALE,
	}
	log.Printf("Hit userSvc.Register %+v\n", user2)
	_, _ = userSvc.Register(ctx, user2)

	log.Printf("Hit userSvc.List\n")
	users, _ := userSvc.List(ctx, new(empty.Empty))
	log.Printf("List Users %+v\n", users.GetList())
}
