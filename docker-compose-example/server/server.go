package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

type Response struct {
	FileData []byte `json:"fileData"`
	CheckSum string `json:"checksum"`
}

func main() {
	port := flag.String("Port", "8090", "a string")
	flag.Parse()
	log.Println("port ", *port)
	e := echo.New()
	e.GET("/getFile", func(c echo.Context) error {
		err := ioutil.WriteFile("/serverdata/test.txt", []byte("Testing"), os.ModePerm)
		if err != nil {
			log.Fatal("File write error", err)
		}
		fileData, err := ioutil.ReadFile("/serverdata/test.txt")
		if err != nil {
			log.Fatal("File reading error", err)
		}
		hasher := sha256.New()
		checksum := hex.EncodeToString(hasher.Sum(fileData))
		log.Println("File contain", string(fileData), err, checksum)
		return c.JSON(http.StatusOK, Response{FileData: fileData, CheckSum: checksum})
	})
	e.Logger.Fatal(e.Start(":" + *port))
}
