package main

import (
	"goshop/libs/httpclient"

	"fmt"
	"log"
	"strconv"

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

func testRegister() {
	m := make(map[string]string, 0)
	m["username"] = "admin"
	m["password"] = "admin"

	b, err := httpclient.Post(testUrl+"/api/user/register", m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// 批量注册
func testMutilRegister() {
	m := make(map[string]string, 0)
	for i := 0; i < 100; i++ {
		m["username"] = "admin_" + strconv.Itoa(i)
		m["password"] = "admin_" + strconv.Itoa(i)
		httpclient.Post(testUrl+"/api/user/register", m)
	}
}

func testUserList() {
	b, err := httpclient.Get(testUrl + "/api/user/userlist?page=1&limit=2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	b, err = httpclient.Get(testUrl + "/api/user/userlist?page=2&limit=2")
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

	fmt.Println("开始测试注册....")
	testRegister()

	fmt.Println("开始测试登录....")
	testLogin()

	fmt.Println("开始测试批量注册....")
	testMutilRegister()

	fmt.Println("开始测试用户列表....")
	testUserList()

}
