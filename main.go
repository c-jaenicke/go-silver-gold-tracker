package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Source
// https://www.gold.de/chartgenerator/
var url = "https://api.edelmetalle.de/public.json"
var respObj ResponseObject

func makeRequest() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to make request", err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("response code was not 200! got: ", resp.StatusCode)
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body:", err)
	}

	json.Unmarshal(respData, &respObj)
}

func convertTime() string {
	timeInteger, err := strconv.ParseInt(strconv.Itoa(respObj.Timestamp), 10, 64)
	if err != nil {
		log.Fatal("failed to convert time string to integer")
	}

	timestamp := time.Unix(timeInteger, 0)

	return timestamp.Format("15:04")
}

func printWidget() {
	fmt.Println(fmt.Sprintf("%s Ag: %.2f$; Au: %.2f$", convertTime(), respObj.SilberUsd, respObj.GoldUsd))
}

func main() {
	makeRequest()
	printWidget()
}
