package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Base URL for httpbin
	baseURL := "https://httpbin.org"

	// HTTP GET Request
	getResp, err := http.Get(baseURL + "/get")
	if err != nil {
		fmt.Println("Error on GET request:", err)
		return
	}
	defer getResp.Body.Close()
	getBody, _ := io.ReadAll(getResp.Body)
	fmt.Println("GET Response:", string(getBody))

	// HTTP POST Request
	postBody := []byte(`{"key": "value"}`)
	postResp, err := http.Post(baseURL+"/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error on POST request:", err)
		return
	}
	defer postResp.Body.Close()
	postBodyResp, _ := io.ReadAll(postResp.Body)
	fmt.Println("POST Response:", string(postBodyResp))

	// HTTP PUT Request
	client := &http.Client{}
	putBody := []byte(`{"key": "updated value"}`)
	putReq, err := http.NewRequest(http.MethodPut, baseURL+"/put", bytes.NewBuffer(putBody))
	putReq.Header.Set("Content-Type", "application/json")
	putResp, err := client.Do(putReq)
	if err != nil {
		fmt.Println("Error on PUT request:", err)
		return
	}
	defer putResp.Body.Close()
	putBodyResp, _ := io.ReadAll(putResp.Body)
	fmt.Println("PUT Response:", string(putBodyResp))

	// HTTP DELETE Request
	deleteReq, err := http.NewRequest(http.MethodDelete, baseURL+"/delete", nil)
	deleteResp, err := client.Do(deleteReq)
	if err != nil {
		fmt.Println("Error on DELETE request:", err)
		return
	}
	defer deleteResp.Body.Close()
	deleteBodyResp, _ := io.ReadAll(deleteResp.Body)
	fmt.Println("DELETE Response:", string(deleteBodyResp))
}
