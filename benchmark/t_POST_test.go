package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestPostUrlencoded(t *testing.T) {
	resp, err := http.Post("https://www.abcd123.top/api/v1/login",
	"application/x-www-form-urlencoded",
	strings.NewReader("username=test&password=ab123123"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}