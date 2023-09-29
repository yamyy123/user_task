package interfaces

import "new_project/models"

type Iuser interface{
	Adduser(user *models.User)(string,error)
}