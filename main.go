package main

import (
	"fmt"
	//"github.com/marcosvillanueva9/worldplay/github.com/protocolbuffers/protobuf/examples/go/tutorialpb"
	//"google.golang.org/protobuf/proto"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/integration"
)

func main() {
	fmt.Println("Hello WorldPlay")

	//p := tutorialpb.Person{
	//	Id: 1234,
	//	Name: "Santiago",
	//	Email: "mananetsantiago@gmail.com",
	//	Phones: []*tutorialpb.Person_PhoneNumber{
	//		{Number: "2612533134", Type: tutorialpb.Person_MOBILE},
	//		{Number: "38016654", Type: tutorialpb.Person_HOME},
	//	},
	//}

	//book := &tutorialpb.AddressBook{}

	//out, err := proto.Marshal(book)
	//if err != nil {
	//
	//}

	integration.Run()
}
