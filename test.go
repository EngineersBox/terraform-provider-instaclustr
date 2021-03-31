package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://circleci.com/api/v2/project/gh/EngineersBox/terraform-provider-instaclustr/pipeline"

	payload := strings.NewReader("{\"branch\":\"INS-12888-Enable-CircleCI-Build-Test-On-Commit\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Circle-Token", "aed5869fed56c59225e02df650d7f2c1a73baec4")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}