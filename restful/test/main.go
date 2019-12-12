package main

import (
	"goshop/libs/httpclient"

	"fmt"
	"log"
)

const (
	testUrl = "http://127.0.0.1:8003"
)

func testLogin() {
	m := make(map[string]string, 0)
	m["username"] = "admin"
	m["password"] = "admin"

	b, err := httpclient.Post(testUrl+"/api/user/login", m)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

}

func main() {

	fmt.Println("开始测试登录....")
	testLogin()
}
