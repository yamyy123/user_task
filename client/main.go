package main

import (
	"context"
	"fmt"
	"net/http"
	pb "new_project/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting client...")
	r := gin.Default()
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	//r.Static("/static", "./")
	r.POST("/add", func(c *gin.Context) {
		var request pb.AddRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.Adduser(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Response})
	})

	r.POST("/update",func(c *gin.Context){
		var request pb.UpdateRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		res, err := client.UpdateRole(context.TODO(), &request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Response})
	})
	r.Run(":2000")
}