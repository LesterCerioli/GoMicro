package main

import (
	"fmt"

	"github.com/myodc/go-micro/client"
	"github.com/myodc/go-micro/cmd"
	c "github.com/myodc/go-micro/context"
	example "github.com/myodc/go-micro/examples/server/proto/example"
	"golang.org/x/net/context"
)

func pub() {
	msg := client.NewPublication("topic.go.micro.srv.example", &example.Message{
		Say: "This is a publication",
	})

	// create context with metadata
	ctx := c.WithMetadata(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	// publish message
	if err := client.Publish(ctx, msg); err != nil {
		fmt.Println("pub err: ", err)
		return
	}

	fmt.Printf("Published: %v\n", msg)
}

func call(i int) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Call", &example.Request{
		Name: "John",
	})

	// create context with metadata
	ctx := c.WithMetadata(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &example.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp.Msg)
}

func stream() {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("go.micro.srv.example", "Example.Stream", &example.StreamingRequest{
		Count: int64(10),
	})

	rspChan := make(chan *example.StreamingResponse, 10)

	stream, err := client.Stream(context.Background(), req, rspChan)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for rsp := range rspChan {
		fmt.Println("Stream: rsp:", rsp.Count)
	}

	if stream.Error() != nil {
		fmt.Println("stream err:", err)
		return
	}

	stream.Close()
}

func main() {
	cmd.Init()

	for i := 0; i < 10; i++ {
		call(i)
	}

	stream()
	pub()
}
