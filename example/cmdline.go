package main

import (
	"fmt"
	"github.com/ant0ine/go-webfinger"
	"os"
)

func main() {
	email := os.Args[1]

	client := webfinger.NewClient(nil)

	resource, err := webfinger.Parse(email)
	if err != nil {
		panic(err)
	}

	jrd, err := client.GetJRD(resource)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("JRD: %+v", jrd)
}
