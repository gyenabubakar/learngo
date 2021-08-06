package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, httpErr := http.Get("http://google.com")
	if httpErr != nil {
		log.Fatal(httpErr)
	}

	writer := logWriter{}

	if w, e := io.Copy(writer, resp.Body); e != nil {
		log.Fatal(e)
	} else {
		fmt.Println()
		fmt.Println()
		fmt.Println("Bytes written:", w)
	}
}

type logWriter struct {}

func (l logWriter) Write(data []byte) (w int, e error)  {
	w = len(data)
	if w == 0 {
		e = errors.New("no data to be written")
		return
	}
	fmt.Println(string(data))
	return
}