package service

import (
	"context"
	"new_project/interfaces"
	"new_project/models"

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
    
    // Insert the user into the database
    _, err = u.usercollection.InsertOne(u.ctx, user) // Pass 'user' instead of '&user'
    if err != nil {
        return "", err // Return the error instead of a string
    }
    
    return "User has been added successfully", nil
}
