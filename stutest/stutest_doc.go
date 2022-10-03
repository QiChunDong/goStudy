package stutest

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
)

// 必须要以Example开头
func ExampleSendJson(rw http.ResponseWriter, r *http.Request) {
	const ballotX = "\u2717"
	req, _ := http.NewRequest("GET", "/sendJson", nil)
	nr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(nr, req)
	u := struct {
		Name string
		Email string
	} {}

	if err := json.NewDecoder(nr.Body).Decode(&u); err != nil {
		log.Printf("%v 报文解析失败咯：%v\n", ballotX, err)
	}
	
	log.Printf("结果%v\n", u)
	// Output:
	// {Bill 11@qq.com}
}