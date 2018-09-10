package main

import (
	"fmt"
	"log"

	"roboime-atlas/protos"

	"github.com/golang/protobuf/proto"
)

func main() {
	proto.
	test := &sslpkg.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
	}
	data, err := proto.Marshal(test)
	fmt.Println(data)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &sslpkg.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// etc.
}
