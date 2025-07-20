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
	apiKey := creds["desafio05"]["api"]

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{"text": "As inscrições para a São Silvestre de 2025 já estão abertas?"},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(body)
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respData))
}
