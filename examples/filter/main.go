package main

import (
	"context"
	"fmt"

	"go-micro.dev/examples/v4/filter/version"
	proto "go-micro.dev/examples/v4/service/proto"
	"go-micro.dev/v4"
)

func main() {
	service := micro.NewService()
	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(
		// provide a context
		context.TODO(),
		// provide the request
		&proto.Request{Name: "John"},
		// set the filter
		version.Filter("latest"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
