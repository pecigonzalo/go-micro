module go-micro.dev/plugins/wrapper/breaker/gobreaker/v4

go 1.16

require (
	github.com/sony/gobreaker v0.4.1
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
)

replace (
	go-micro.dev/plugins/registry/memory/v4 => ../../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../../go-micro
)
