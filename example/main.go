package main

import (
	"fmt"
	pb "gen-go-tarantool-demo/api/v1/api_tnt"

	tnt "github.com/tarantool/go-tarantool"
)

func main() {
	opts := tnt.Opts{User: "test", Pass: "test"}
	conn, err := tnt.Connect("127.0.0.1:3401", opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	x := pb.RequestPutAttrValue{AttrValues: []*pb.AttrValue{&pb.AttrValue{
		UnitGuid:  "xx-1",
		ObjectId:  "yy-1",
		AttrId:    "atr-1",
		Timestamp: 01234,
		Value:     345,
		Action:    pb.Action_ACTION_ACTION_CLICKHOUSE,
	}}}

	var y []pb.RequestPutAttrValue

	err = conn.Call17Typed("push_messages", []interface{}{&x}, &y)
	if err != nil {
		panic(err)
	}

	for _, msg := range y[0].GetAttrValues() {
		fmt.Printf("%#v", msg)
	}

}
