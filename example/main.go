package main

import (
	"fmt"
	pb "gen-go-tarantool-demo/demo/v1"
	tnt "github.com/tarantool/go-tarantool"
)

func main() {
	opts := tnt.Opts{User: "test", Pass: "test"}
	conn, err := tnt.Connect("127.0.0.1:3401", opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var x = pb.Ex{
		IdNum: "1",
		Value: &pb.Ex_IData{IData: 1},
	}

	var y []pb.Ex

	err = conn.Call17Typed("push_messages", []interface{}{&x}, &y)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", y[0])

}
