package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NYTimes/gizmo/pubsub/http"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("hello")
	elliot := &Person{
		Name: "Elliot",
		Age:  24,
	}

	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)
	p := http.NewPublisher("http://localhost:8900", nil)

	// msg := []proto.Message{text}

	err = p.Publish(context.Background(), "", elliot)
	// if err != nil {
	// 	// handle error
	// }
}
