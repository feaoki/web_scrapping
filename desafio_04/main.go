package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	credFile := "../.credenciais/credenciais.json"
	credData, err := os.ReadFile(credFile)
	if err != nil {
		panic(fmt.Sprintf("Erro ao ler o arquivo de credenciais: %v", err))
	}
	var creds map[string]map[string]string
	if err := json.Unmarshal(credData, &creds); err != nil {
		panic(fmt.Sprintf("Erro ao fazer unmarshal das credenciais: %v", err))
	}
	apiKey := creds["desafio04"]["api_key"]

	url := "https://api.openai.com/v1/chat/completions"
	body := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": "as inscrições para a são silvestre 2025 estão abertas?"},
		},
	}

	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+apiKey)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(responseData))
}
