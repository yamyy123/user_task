package controller

import (
	"context"
	"new_project/interfaces"
	"new_project/models"
	pb "new_project/proto"
)

type RPCserver struct{
	pb.UnimplementedUserServer
}

var (
	UserService interfaces.Iuser
)

func(u *RPCserver)Adduser(ctx context.Context,req *pb.AddRequest)(*pb.AddResponse,error){
	  user:=&models.User{
	  	Name:     req.Name,
	  	Email:    req.Email,
	  	Password: req.Password,
	  	Dob:      req.Dob,
	  	Phone:    req.Phone,
	  	Role:     req.Role,
	  	Status:   req.Status,
	  }
	  result,err:= UserService.Adduser(user)
	  if err!=nil{
		return &pb.AddResponse{
			Response: "failure",
		},err
	  }
	
	return &pb.AddResponse{
		Response: result,
	},nil
}