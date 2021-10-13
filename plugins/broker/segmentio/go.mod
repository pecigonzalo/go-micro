module go-micro.dev/plugins/broker/segmentio/v4

go 1.16

require (
	github.com/google/uuid v1.2.0
	github.com/segmentio/kafka-go v0.4.16
	go-micro.dev/plugins/broker/kafka/v4 v4.1.0
	go-micro.dev/plugins/codec/segmentio/v4 v4.1.0
	go-micro.dev/v4 v4.1.0
)

replace (
	go-micro.dev/plugins/broker/kafka/v4 => ../kafka
	go-micro.dev/plugins/codec/segmentio/v4 => ../../codec/segmentio
	go-micro.dev/v4 => ../../../../go-micro
)
