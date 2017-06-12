package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/atotto/clipboard"
)

func main() {	
	var clipboard_text string

	if(len(os.Args) < 2) {
		fmt.Println("Você deve selecionar o o padrão!")
		return
	}

	clipboard_file, err := os.Open(string(os.Args[1]))
	if err != nil {
		return
	}
	defer clipboard_file.Close()

	stat, err := clipboard_file.Stat()
	if err != nil {
  	return
  }

  clipboard_bytes := make([]byte, stat.Size())

	_, err = clipboard_file.Read(clipboard_bytes)
	if err != nil {
		return
  }

  clipboard_text = string(clipboard_bytes)

  for i := 0; i < len(os.Args)-1; i++ {
  	clipboard_text = strings.Replace(clipboard_text, "{{ $"+strconv.Itoa(i)+" }}", string(os.Args[i+1]), -1)
  }

  err = clipboard.WriteAll(clipboard_text)
  if err != nil {
  	fmt.Println(err)
  } else {
  	fmt.Println("Copiado para o clipboard!")
  }

}
