package main

import (
	"net/http";
	"net/url";
)

type credertial struct {
	api_id		int
	api_key		string
	api_sender 	string
}

func SendSMS(params credertial, to string, message string) {
	v := url.Values{}
	u := url.URL{}

	v.Add("id", string(params.api_id))
	v.Add("key", params.api_key)
	v.Add("ti", to)
	v.Add("from", params.api_sender)
	v.Add("text", message)

	u.Scheme = "http"
	u.Host = "bytehand.com:3800"
	u.Path = "send"
	u.RawQuery = v.Encode()

	http.Get(u.String())
}
