package api

import (
	"encoding/json"
	"fmt"
	"kpdl/api/struct"
	"log"
	"strconv"
	"time"
)
import "net/http"

type Api struct {
	ApiKey string
}

func (api Api) GetDate(method string) _struct.Response {
	var myClient = &http.Client{Timeout: 10 * time.Second}
	url := "https://api.service-kp.com/v1/" + method + "?access_token=" + api.ApiKey
	//fmt.Println("url: ", url)
	resp, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var data _struct.Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}
	return data
}

func (api Api) GetInfoItem(id int) _struct.Item {
	r := api.GetDate("items/" + strconv.Itoa(id))
	return r.Item
}
