package models

type User struct {
	Name     string   `json:"name" bson:"name"`
	Email    string   `json:"email" bson:"email"`
	Password string   `json:"password" bson:"password"`
	Dob      string   `json:"dob" bson:"dob"`
	Phone    int64    `json:"phone" bson:"phone"`
	Role     []string `json:"role" bson:"role"`
	Status   string   `json:"status" bson:"status"`
}

type Rolerequest struct{
	Name string `json:"name" bson:"name"`
	Role []string `json:"role" bson:"role"`
}
