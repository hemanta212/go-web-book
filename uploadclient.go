package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename, url string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very imp
	// Multipart deals with anything we need to create multipart form
	// form file uploads we have to give a buffer first
	// Then we tell it to convert it to a FormFile format using the buffer to take from storage to memory
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)

	if err != nil {
		fmt.Println("Error writing to buffer")
		return err
	}

	// open file handle
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func runClient() {
	targetUrl := "http://localhost:5000/upload"
	filename := "./test.xlsx"
	postFile(filename, targetUrl)

}
