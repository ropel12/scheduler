package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ropel12/scheduler/entities"
)

func ApiCall(token string) entities.HttpResponse {
	req, err := http.NewRequest("GET", "http://localhost:8000/quiz/cron/"+token, nil)
	if err != nil {
		log.Printf("[ERROR]WHEN GETTING DATA FROM API, err:%v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR]WHEN GETTING DATA FROM API, err:%v", err)

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR]WHEN GETTING DATA FROM API, err:%v", err)
	}

	var responseData entities.HttpResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("[ERROR]WHEN UNMARSHAL RESPONSE DATA FROM API, err:%v", err)
	}
	return responseData
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
