module go-micro.dev/plugins/wrapper/ratelimiter/ratelimit/v4

go 1.16

require (
	github.com/juju/ratelimit v1.0.1
	go-micro.dev/plugins/broker/memory/v4 v4.1.0
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/plugins/transport/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
)

replace (
	go-micro.dev/plugins/broker/memory/v4 => ../../../../plugins/broker/memory
	go-micro.dev/plugins/registry/memory/v4 => ../../../../plugins/registry/memory
	go-micro.dev/plugins/transport/memory/v4 => ../../../../plugins/transport/memory
	go-micro.dev/v4 => ../../../../../go-micro
)
