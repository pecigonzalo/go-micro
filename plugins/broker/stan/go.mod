module go-micro.dev/plugins/broker/stan/v4

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/nats-io/nats-server/v2 v2.3.0 // indirect
	github.com/nats-io/nats-streaming-server v0.22.0 // indirect
	github.com/nats-io/stan.go v0.9.0
	go-micro.dev/v4 v4.1.0
)

replace go-micro.dev/v4 => ../../../../go-micro
