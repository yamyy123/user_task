package service

import (
	"context"
	"fmt"
	"log"
	"new_project/interfaces"
	"new_project/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx            context.Context
	usercollection *mongo.Collection
}

func NewUserServiceInit(ctx context.Context, usercollection *mongo.Collection) interfaces.Iuser {     //dependency injection method
	return &UserService{ctx, usercollection}
}

func (u *UserService) Adduser(user *models.User) (string, error) {
    // Hash the user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return "", err // Return the error instead of a string
    }
    
    user.Password = string(hashedPassword)
    fmt.Println("i came here")
    // Insert the user into the database
    _, err = u.usercollection.InsertOne(u.ctx, &user) // Pass 'user' instead of '&user'
    if err != nil {
        return "", err // Return the error instead of a string
    }
    
    return "User has been added successfully", nil
}


func (u *UserService) UpdateRole(role *models.Rolerequest)(string,error){
    filter := bson.D{{Key: "name",Value: role.Name},{Key: "status",Value: "enabled"}}
    var result *models.User
     err:=u.usercollection.FindOne(u.ctx,filter).Decode(&result)
     if err!=nil{
      log.Fatal(err.Error())
     }
     filter2:=bson.D{{Key: "$set",Value: bson.D{{Key: "role",Value: role.Role}}}}
     _, err =u.usercollection.UpdateOne(u.ctx,filter,filter2)
     if err!=nil{
        log.Fatal(err.Error())
     }
     return "Role updated successfully",nil

}
