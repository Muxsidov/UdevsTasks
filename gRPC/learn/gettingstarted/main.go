// https://www.youtube.com/watch?v=NoDRq6Twkts
// https://github.com/evilsocket/opensnitch/issues/373

package main

import (
	"fmt"
	"log"
	
	proto "github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("Hello World!")
	
	peter := &Contact{
		Id: 0,
		Name: "Peter",
		Gender: "male",
		Number: 5550001,
		mail: "petergr@gmail.com",
	}
	
	data, err := proto.Marshal(peter)
	if err != nil {
		log.Fatal("Marshaling error", err)
	}
	
	fmt.Println(data)
}
