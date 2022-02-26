package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:123456@127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to MongoDB!")

	collection := client.Database("temp").Collection("aaa")
	//fmt.Println(collection)

	filter := bson.D{{"name", "Ash"}}

	// 查询文档
	var result Trainer
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// 查询多个文档
	one, _ := collection.Find(context.TODO(), filter)
	defer func() {
		if err := one.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	c := []Trainer{}
	_ = one.All(context.TODO(), &c)
	for _, r := range c {
		fmt.Println(r)
	}

	//删除文档
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	fmt.Println(deleteResult.DeletedCount)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")

}
