package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/madisonsmartin/projHopExp/handler"
	"github.com/madisonsmartin/projHopExp/subscriber"

	example "github.com/madisonsmartin/projHopExp/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.projHopExp"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.projHopExp", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.projHopExp", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
