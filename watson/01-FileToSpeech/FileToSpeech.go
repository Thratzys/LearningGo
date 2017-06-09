package main

import (
"fmt"
"io"
"net/http"
"encoding/json"
"os"
"bytes"
)

const (
	speech_user = "{{ USER }}"
	speech_pass = "{{ PASS }}"
)

type T struct {
	Text string `json:"text"`
}

func main() {

	if(len(os.Args) < 2) {
		fmt.Println("Requer arquivo texto a ser lido.")
		return
	}

	file, err := os.Open(string(os.Args[1]))
	if err != nil {
		//	
		return
	}
	defer file.Close()


	stat, err := file.Stat()
	if err != nil {
		//
		return
	}

	file_bytes := make([]byte, stat.Size())

	_, err = file.Read(file_bytes)

	file_text := string(file_bytes)

	data := T{Text: file_text}

	requestBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	text_body := bytes.NewReader(requestBytes)

	file_output, err := os.Create(os.Args[1]+".wav")
	if err != nil {
		//
		return
	}
	defer file_output.Close()

	request, err := http.NewRequest("POST", "https://stream.watsonplatform.net/text-to-speech/api/v1/synthesize", text_body)
	if err != nil {
		return
	}

	request.SetBasicAuth(speech_user, speech_pass)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "audio/wav")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}

	defer response.Body.Close()   

	if _, err := io.Copy(file_output, response.Body); err != nil {
		//
		return
	}

	fmt.Println("Salvo: "+os.Args[1]+".wav")
}
