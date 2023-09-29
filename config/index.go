package config

import (
	"context"
	"fmt"
	"log"
	"new_project/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func Connectdatabase()(*mongo.Client,error){
ctx,_ := context.WithTimeout(context.Background(), 10*time.Second)
mongoConnection :=options.Client().ApplyURI(constants.ConnectionString)
mongoClient,err:= mongo.Connect(ctx,mongoConnection)
if err!=nil{
	fmt.Println("An error has been encountered when trying to connecting to db")
	log.Fatal(err.Error())
	return nil,err
}
if err:=mongoClient.Ping(ctx,readpref.Primary());err!=nil{
         log.Fatal(err.Error())
		 return nil,err
}

return mongoClient,nil
}

func Getcollection(client *mongo.Client,dbname string,collectionname string)(*mongo.Collection){
     collection:=client.Database(dbname).Collection(collectionname)
	 return collection
}

