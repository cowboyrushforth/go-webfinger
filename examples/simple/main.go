// Minimal example, used in the README
package main

import (
	"fmt"
	"github.com/ant0ine/go-webfinger"
	"os"
)

func main() {
	email := os.Args[1]

	client := webfinger.NewClient(nil)
	client.AllowHTTP = true

	jrd, err := client.Lookup(email, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("JRD: %+v", jrd)
}
