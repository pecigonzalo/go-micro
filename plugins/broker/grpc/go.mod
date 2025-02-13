module go-micro.dev/plugins/broker/grpc/v4

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	google.golang.org/grpc v1.38.0
)

replace (
	go-micro.dev/plugins/registry/memory/v4 => ../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../go-micro
)
