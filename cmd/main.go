package main

import (
	"fmt"
	"log"

	youtube "github.com/burravlev/goytdl"
)

func main() {
	client := youtube.Client{}
	video, err := client.GetVideo("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(video)
}
