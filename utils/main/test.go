package main

import (
	"fmt"
	"github.com/Anteoy/liongo/utils"
	"time"
)

func main() {
	token, err := utils.GenToken(1)
	fmt.Printf("%s,%+v\n", token, err)

	result := utils.ValidateToken(token)
	fmt.Println(result)

	time.Sleep(15 * time.Second)
	fmt.Println("===")
	result = utils.ValidateToken(token)
	fmt.Println(result)
}
