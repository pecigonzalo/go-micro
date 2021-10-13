module go-micro.dev/plugins/client/grpc/v4

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20210902184326-c93e472777b9
)

replace (
	go-micro.dev/plugins/registry/memory/v4 => ../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../go-micro
)
