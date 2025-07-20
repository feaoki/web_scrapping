package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiKey := "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"

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
