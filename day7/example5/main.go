package main

//json格式
import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int
	Birthday string
	Sex      string
	Email    string
	Address  string
}

func testStruct() {
	user1 := &User{
		UserName: "moyuanhui",
		NickName: "allen",
		Age:      23,
		Birthday: "1994-01-02",
		Sex:      "男",
		Email:    "allen@173.com",
		Address:  "大连",
	}
	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json.marshal failed,err:", err)
		return
	}
	fmt.Printf("%s\n", data)
}

func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 23
	m["sex"] = "man"

	s = append(s, m)
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 24
	m["sex"] = "female"

	s = append(s, m)
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Print("json marsha1 faild")
		return
	}
	fmt.Println(string(data))
}

func main() {
	// testStruct()
	testSlice()
}
