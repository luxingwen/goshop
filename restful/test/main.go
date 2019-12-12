package main

import (
	"goshop/libs/httpclient"

	"fmt"
	"log"

	"goshop/restful/models"

	"github.com/bxcodec/faker"
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

func testGenerate() {
	user := models.User{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", user)
}

func main() {
	fmt.Println("开始生成数据...")

	testGenerate()

	fmt.Println("开始测试登录....")
	testLogin()

}
