package main

import (
	"fmt"
	"log"

	"github.com/sourcefellows/mylib"
	newver "github.com/sourcefellows/mylib/v2"
)

func main() {

	result, err := mylib.HideDigits("meine12345Nummer")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	result = newver.HideDigits("123")
	fmt.Println(result)

}
