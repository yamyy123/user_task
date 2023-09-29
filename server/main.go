package main

import (
	"context"
	"fmt"
	"net"
	"new_project/config"
	"new_project/constants"
	"new_project/controller"
	"new_project/service"
     pb "new_project/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func inituser(client *mongo.Client)(){
        ctx:=context.TODO()
		 usercollection :=config.Getcollection(client,constants.DatabaseName,"users")
        controller.UserService = service.NewUserServiceInit(ctx, usercollection)
}

func main(){
	mongoclient, err := config.Connectdatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	inituser(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	//It creates a new gRPC server instance s using grpc.NewServer().

	pb.RegisterUserServer(s, &controller.RPCserver{})
     //this line connect the  comntroller and server 
	 //It registers the controllers.RPCServer{} as the gRPC server implementation for the pro package's CustomerServiceServer.
	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}