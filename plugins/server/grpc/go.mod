module go-micro.dev/plugins/server/grpc/v4

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	go-micro.dev/plugins/broker/memory/v4 v4.1.0
	go-micro.dev/plugins/client/grpc/v4 v4.1.0
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/plugins/transport/grpc/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.38.0
)

replace (
	go-micro.dev/plugins/broker/memory/v4 => ../../../plugins/broker/memory
	go-micro.dev/plugins/client/grpc/v4 => ../../../plugins/client/grpc
	go-micro.dev/plugins/registry/memory/v4 => ../../../plugins/registry/memory
	go-micro.dev/plugins/transport/grpc/v4 => ../../../plugins/transport/grpc
	go-micro.dev/v4 => ../../../../go-micro
)
