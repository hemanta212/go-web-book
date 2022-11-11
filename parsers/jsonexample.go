package parsers

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type ServerInfo struct {
	ServerName string `json:"server-Name"`
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []ServerInfo `json:"servers"`
}

// Simple parsing. one to one mapping if we already know the json structure in advance
func ParseJSONToStruct() {
	var s ServerSlice
	data, err := ReadFile("parsers/servers.json")
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	json.Unmarshal(data, &s)
	fmt.Println(s)
}

// More general parsing of data we dont know anything about using go reflection
func ParseJSONStd() {
	data := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		fmt.Printf("error : %v", err)
		return
	}
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is a stirng", vv)
		case int:
			fmt.Println(k, "is a int", vv)
		case float64:
			fmt.Println(k, "is a float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is a type not handled", vv)

		}
	}
}

// Best way to parse unkown json using simplejson package from bitly
func ParseJSON() {
	js, err := simplejson.NewJson([]byte(`{
            "test": {
                "array": [1, "2", 3],
                "int": 10,
                "float": 5.150,
                "bignum": 9223372036854775807,
                "string": "simplejson",
                "bool": true
            }
        }`))
	if err != nil {
		fmt.Printf("Error parsing json %v", err)
		return
	}
	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()
	fmt.Printf("Array:%v, Int: %v, string: %s\n", arr, i, ms)
	return
}

func GenJSON() {
	var s ServerSlice
	s.Servers = append(s.Servers, ServerInfo{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, ServerInfo{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
