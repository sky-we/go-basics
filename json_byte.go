package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//json string ---> []byte
	//str := `{"name":"John Doe"}`
	str := "{\"name\":\"John Doe\"}"
	ToByteArr, _ := json.Marshal(str)
	fmt.Println(ToByteArr) // [34 123 92 34 110 97 109 101 92 34 58 92 34 74 111 104 110 32 68 111 101 92 34 125 34] 双引号34 反斜杠92

	// []byte---> json string
	byteArr2 := []byte{34, 123, 92, 34, 110, 97, 109, 101, 92, 34, 58, 92, 34, 74, 111, 104, 110, 32, 68, 111, 101, 92, 34, 125, 34}
	var receive2r string
	err2 := json.Unmarshal(byteArr2, &receive2r)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(receive2r) // {"name":"John Doe"} 反引号34被省略

	//  byte[] ----> map struct
	byteArr := []byte{123, 34, 110, 97, 109, 101, 34, 58, 32, 34, 74, 111, 104, 110, 32, 68, 111, 101, 34, 125}
	var receiver map[string]string
	err := json.Unmarshal(byteArr, &receiver)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(receiver)

	// map struct ----> byte[]
	maps := map[string]string{
		"name": "John Doe",
	}
	result, err := json.Marshal(maps)
	fmt.Println(result) // [123 34 110 97 109 101 34 58 34 74 111 104 110 32 68 111 101 34 125]

	// json string -----类型转换---->  byte[]  ----反序列化----> go struct

	srcStr := `{"name":"John Doe"}`

	var receiver3 map[string]string

	err4 := json.Unmarshal([]byte(srcStr), &receiver3)
	if err4 != nil {
		return
	}
	fmt.Println(receiver3["name"])

	// go struct -----序列化---->  byte[]  ----类型转换----> json string

	receiver5 := map[string]string{
		"name": "John Doe",
	}
	result, err5 := json.Marshal(receiver5)
	if err5 != nil {
		return
	}
	fmt.Println(result)
	fmt.Println(string(result))

}
