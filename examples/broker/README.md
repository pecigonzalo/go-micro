# Broker

The [broker](https://godoc.org/go-micro.dev/broker#Broker) is an interface for PubSub.

## Contents

- main.go - uns pub-sub as two go routines for 10 seconds.
- producer - publishes messages to the broker every second
- consumer - consumes any messages sent by the producer
