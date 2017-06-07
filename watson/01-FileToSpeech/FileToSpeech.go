package main

import (
"fmt"
"io"
"net/http"
"strings"
"os"
)

const (
	speech_user = "your_user"
	speech_pass = "your_pass"
)

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

	file_text = strings.Replace(file_text, " ", "+", -1)
	file_text = strings.Replace(file_text, "\n", "+", -1)

	file_output, err := os.Create(os.Args[1]+".wav")
	if err != nil {
		//
		return
	}
	defer file_output.Close()

	file_url := fmt.Sprintf("https://stream.watsonplatform.net/text-to-speech/api/v1/synthesize?accept=audio/wav&text="+file_text)
	request, err := http.NewRequest("GET", file_url, nil)
	if err != nil {
		//x
		return
	}        

	request.SetBasicAuth(speech_user, speech_pass)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		//
		return
	}

	defer response.Body.Close()       

	if _, err := io.Copy(file_output, response.Body); err != nil {
		//
		return
	}

	fmt.Println("Arquivo gerado!")
}
