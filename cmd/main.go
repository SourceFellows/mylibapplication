package main

import (
	"fmt"
	"log"

	"github.com/sourcefellows/mylib"
)

func main() {

	result, err := mylib.HideDigits("meine12345Nummer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

}
