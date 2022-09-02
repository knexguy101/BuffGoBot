package main

import (
	"encoding/json"
	"github.com/go-rod/rod/lib/proto"
	"os"
)

func loadCookies() ([]*proto.NetworkCookie, error) {
	s, err := os.ReadFile("./cookies.json")
	if err != nil {
		return nil, err
	}

	var cookies []*proto.NetworkCookie
	err = json.Unmarshal(s, &cookies)
	return cookies, err
}

func saveCookies(c []*proto.NetworkCookie) error {
	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile("./cookies.json", b, 0777)
}