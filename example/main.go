package main

import (
	"fmt"
	tnt2 "gen-go-tarantool-demo/api/v1/api_tnt"
	tnt "github.com/tarantool/go-tarantool"
)

func main() {
	opts := tnt.Opts{User: "test", Pass: "test"}
	conn, err := tnt.Connect("127.0.0.1:3401", opts)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	vl := tnt2.UnitStatus{
		UnitGuid:      "x-11",
		ConnStates:    map[string]int32{"hello": 23, "world": 24},
		ActiveMessage: map[string]int64{"hello": 23, "world": 24},
		Timestamp:     0,
	}

	fmt.Println(conn.Call17("push_messages", []interface{}{&vl}))

	var ret []tnt2.UnitStatus
	fmt.Println(conn.Call17Typed("push_messages", []interface{}{&vl}, &ret))
	fmt.Println(ret)
}
