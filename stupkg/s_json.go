// json比xml
// 处理的标签要少
// 网络传输时消息数据更少
// 可转换为BSON(Binary JavaScript Object Notation 二进制js对象标记),进一步缩小消息长度
package stupkg

import (
	"encoding/json"
	"fmt"
)

type Contact struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "323.444.444.4",
		"cell": "32.45.4"
	}
}`

func TestUnmashal() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}
	fmt.Println("json unmarshal success:", c)
}

func TestUnmashalMap() {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &m)
	if err != nil {
		fmt.Println("json unmarshal err:", err)
		return
	}
	fmt.Println("name:", m["name"])
	fmt.Println("title:", m["title"])
	fmt.Println("contact.home:", m["contact"].(map[string]interface{})["home"])
	fmt.Println("contact.cell:", m["contact"].(map[string]interface{})["cell"])
}


func TestJsonEncode() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "323.444.444.4",
		"cell": "32.45.4",
	}

	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		fmt.Println("marshal err:", err)
		return
	}
	fmt.Println("marshal success:", string(data))
}