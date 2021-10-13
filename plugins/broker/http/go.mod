module go-micro.dev/plugins/broker/http/v4

go 1.16

require (
	github.com/google/uuid v1.2.0
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
)

replace (
	go-micro.dev/plugins/registry/memory/v4 => ../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../go-micro
)
