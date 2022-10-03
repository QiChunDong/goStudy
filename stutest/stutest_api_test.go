package stutest

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Route() {
	http.HandleFunc("/sendJson", SendJson)
}

func SendJson(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name string
		Email string
	} {
		Name: "Bill",
		Email: "111@qq.com",
	}

	rw.WriteHeader(200)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(&u)
}

// 初始化路由
func init() {
	Route()
}

func TestApi(t *testing.T) {
	t.Log("\t开始测试API")
	{
		req, err := http.NewRequest("GET", "/sendJson", nil)
		if err != nil {
			log.Fatal("sendJson请求不通啊", ballotX, err)
		}
		t.Log("sendJson可以请求通！", checkMark)

		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Errorf("%v 收到响应码%d 应该是%v", ballotX, rw.Code, 200)
		} else {
			t.Logf("%v 收到响应码%d", checkMark, rw.Code)
		}

		u := struct {
			Name string
			Email string
		} {}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatalf("%v 报文解析失败咯：%v", ballotX, err)
		}
		t.Logf("%v 报文解析成功！", checkMark)

		if u.Name == "Bill" {
			t.Logf("%v 姓名是对的：%v", checkMark, u.Name)
		} else {
			t.Fatalf("%v 姓名不是Bill：%v", ballotX, u.Name)
		}

		if u.Email == "111@qq.com" {
			t.Logf("%v 邮箱是对的：%v", checkMark, u.Email)
		} else {
			t.Fatalf("%v 邮箱不是111@qq.com：%s", ballotX, u.Email)
		}
	}
}