package natstest

import (
	"github.com/nats-io/gnatsd/server"
	gnatsd "github.com/nats-io/gnatsd/test"
	"github.com/nats-io/nats.go"
)

var DefaultTestOptions = server.Options{
	Host:           "127.0.0.1",
	Port:           server.RANDOM_PORT,
	NoLog:          true,
	NoSigs:         true,
	MaxControlLine: 256,
}

func RunDefaultServer() *server.Server {
	return gnatsd.RunServer(&DefaultTestOptions)
}

func serverAddr(s *server.Server) string {
	u := "nats://" + s.Addr().String()
	return u
}

func main() {
	server := RunDefaultServer()
	nc, err := nats.Connect(serverAddr(server), nats.Name("test_server"))
	if err != nil {
		panic(err)
	}
	_, err = nc.Subscribe("topic", func(msg *nats.Msg) {})
	if err != nil {
		panic(err)
	}
}
