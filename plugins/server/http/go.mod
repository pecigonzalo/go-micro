module go-micro.dev/plugins/server/http/v4

go 1.16

require (
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
)

replace (
	go-micro.dev/plugins/registry/memory/v4 => ../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../go-micro
)
