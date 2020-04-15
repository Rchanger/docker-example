package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	FileData []byte `json:"fileData"`
	CheckSum string `json:"checksum"`
}

func main() {
	host := os.Getenv("ServerHost")
	port := os.Getenv("ServerPort")
	url := fmt.Sprintf("http://" + host + ":" + port + "/getFile")
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating http request", err)
	}
	request.Header.Set("content-type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal("Error", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading resp body:", err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Err: request not proceeded", string(body))
	}

	result := Response{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal("Error while unmarshal response", err)
	}
	hasher := sha256.New()

	checksum := hex.EncodeToString(hasher.Sum(result.FileData))
	if strings.EqualFold(strings.TrimSpace(checksum), strings.TrimSpace(result.CheckSum)) {
		log.Println("File checksum mached correct file")
		err := ioutil.WriteFile("/clientdata/test.txt", result.FileData, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	http.ListenAndServe(":8080", nil)
}
