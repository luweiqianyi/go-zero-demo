package token

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	secretKey = "google"
)

type JsonObj struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
}

func TestGenToken(t *testing.T) {
	data := "hello world"
	token, err := GenerateToken(secretKey, data, time.Hour*24*365)
	if err != nil {
		return
	}

	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	data := "hello world"
	token, err := GenerateToken(secretKey, data, time.Hour*24*365)
	if err != nil {
		return
	}

	parseToken, err := ParseToken(token, secretKey)
	fmt.Printf("%#v %v\n", parseToken, err)

	if err != nil {
		return
	}

	ret := parseToken.(string)
	fmt.Println(ret)
	if ret == data {
		fmt.Println("success")
	}
}

func TestParseToken2(t *testing.T) {
	data := JsonObj{
		AccountName: "zhang san",
		Password:    "123456",
	}
	byteData, _ := json.Marshal(data)
	strData := string(byteData)
	token, err := GenerateToken(secretKey, strData, time.Hour*24*365)
	if err != nil {
		return
	}

	parseToken, err := ParseToken(token, secretKey)
	fmt.Printf("%#v %v\n", parseToken, err)
	if err != nil {
		return
	}

	ret := parseToken.(string)
	fmt.Println(ret)
	if ret == strData {
		fmt.Println("success")
	}
}
