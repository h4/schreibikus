package main

import (
	"net/http";
	"net/url";
	"fmt";
	"io/ioutil";
	"log";
	"strconv";
)

type credertial struct {
	api_id		int
	api_key		string
	api_sender 	string
}

func SendSMS(params credertial, to string, message string) {
	v := url.Values{}
	u := url.URL{}

	v.Add("id", strconv.Itoa(params.api_id))
	v.Add("key", params.api_key)
	v.Add("to", to)
	v.Add("from", params.api_sender)
	v.Add("text", message)

	u.Scheme = "http"
	u.Host = "bytehand.com:3800"
	u.Path = "send"
	u.RawQuery = v.Encode()

	fmt.Printf("%s\n", u.String())

	_, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}
}
