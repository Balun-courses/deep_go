package main

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	msg := &anypb.Any{
		TypeUrl: "http://google.ru",
	}
	opts := protojson.UnmarshalOptions{DiscardUnknown: true}
	var res []byte
	_ = protojson.UnmarshalOptions.Unmarshal(opts, res, msg)

	fmt.Println(res)

	//func (o UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
	//	return o.unmarshal(b, m)
	//}
}
