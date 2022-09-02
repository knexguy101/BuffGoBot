package main

import (
	"bufio"
	"github.com/knexguy101/BuffGo/buff/login"
	structs "github.com/knexguy101/BuffGo/models/client"
	"os"
	"strconv"
	"strings"
)

func main() {

	cookies, err := loadCookies()
	if err != nil {

		data, err := login.Login()
		if err != nil {
			panic(err)
		}

		err = saveCookies(data.Cookies)
		if err != nil {
			panic(err)
		}

		cookies = data.Cookies
	}

	client := structs.NewHttpClient()

	login.AddLoginDataToClient(&login.LoginData{
		Cookies: cookies,
	}, client)

	filters, err := readLines("./filters.csv")
	if err != nil {
		panic(err)
	}

	for k, v := range filters {
		if k == 0 {
			continue
		}
		s := strings.Split(v, ",")
		price, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			panic(err)
		}
		go flow(s[0], price, client)
	}


	c := make(chan bool)
	<- c
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}