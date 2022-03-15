package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func LoadAPI(url string) ViewData {
	fmt.Println("Loading API from :")
	fmt.Println("\t", url)

	vd := ViewData{}

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		return vd
		log.Fatal(jsonErr)
	}

	fmt.Println("API loaded successfully.")
	return vd
}
