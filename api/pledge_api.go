package main

import (
	"fmt"
	"pledge-backend-practise/config"
)

func main() {
	fmt.Println("pledge api")

	fmt.Println(config.Config.Env.Port)
}
