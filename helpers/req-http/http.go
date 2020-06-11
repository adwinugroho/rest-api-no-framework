package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	//Request parameter
	Request struct {
		Protocol string
		Host     string
		Port     int
		Path     string //url
		Body     interface{}
	}

	//Response body
	Response struct {
		Status       bool        `json:"status"`
		Code         int         `json:"code"`
		ErrorMessage string      `json:"errMessage,omitempty"`
		Data         interface{} `json:"data,omitempty"`
		Message      interface{} `json:"message,omitempty"`
	}
)

func PostDynamic(req *Request) (interface{}, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	//in example http://string:port/path
	resp, err := http.Post(fmt.Sprintf("%s://%s:%d%s", req.Protocol, req.Host, req.Port, req.Path), "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	defer resp.Body.Close()
	//initial response
	var response interface{}
	//same as json.Marshall
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, nil
}

//use this function if service/API always return code and status
func Post(req *Request) (Response, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	resp, err := http.Post(fmt.Sprintf("%s://%s:%d%s", req.Protocol, req.Host, req.Port, req.Path), "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	defer resp.Body.Close()
	var response Response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode == 200 {
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			fmt.Println(err)
			return response, err
		}
	} else {
		response = Response{Code: resp.StatusCode, Status: false, ErrorMessage: fmt.Sprintf("%s", responseBody)}
	}

	return response, nil
}
