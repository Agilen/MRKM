package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Test() error {
	cert, _ := ioutil.ReadFile("sign.crt")
	req := SignRequest{
		Data: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 45, 65, 7},
		Storage: Storage{
			KeyStore: "./sign.pfx",
			Slot:     "@sign",
			Pin:      "1234",
		},
		Certificate: cert,
	}
	fmt.Println(base64.StdEncoding.EncodeToString([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 45, 65, 7}))
	jsonreq, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r := bytes.NewReader(jsonreq)

	response, err := http.Post("http://127.0.0.1:9999/sign", "application/json; charset=UTF-8", r)
	if err != nil {
		return err
	}

	reader := response.Body

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	var res SignResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	fmt.Println(res)

	jsonreq, err = json.Marshal(res)
	if err != nil {
		return err
	}

	r = bytes.NewReader(jsonreq)
	response, err = http.Post("http://127.0.0.1:9999/verify", "application/json; charset=UTF-8", r)
	if err != nil {
		return err
	}

	reader = response.Body

	body, err = ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	fmt.Println(base64.StdEncoding.DecodeString(string(body)))

	return nil
}
