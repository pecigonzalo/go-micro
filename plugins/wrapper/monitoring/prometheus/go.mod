module go-micro.dev/plugins/wrapper/monitoring/prometheus/v4

go 1.16

require (
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/client_model v0.2.0
	github.com/stretchr/testify v1.7.0
	go-micro.dev/plugins/broker/memory/v4 v4.1.0
	go-micro.dev/plugins/registry/memory/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
)

replace (
	go-micro.dev/plugins/broker/memory/v4 => ../../../../plugins/broker/memory
	go-micro.dev/plugins/registry/memory/v4 => ../../../../plugins/registry/memory
	go-micro.dev/v4 => ../../../../../go-micro
)
