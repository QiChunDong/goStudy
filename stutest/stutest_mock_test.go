package stutest

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var feed = `
<?xml verison="1.0" encoding="UTF-8"?>
<rss>
	<channel>
		<title>csjdhvhjfdv</title>
	</channel>
</rss>
`
// 返回处理请求的服务器指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}
	// HandlerFunc是适配器  适配处理函数
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestDownlodFeed(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("\t开始测试MOCK下载内容")
	{
		t.Logf("测试下载URL：%s, 期望结果：%d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				log.Fatal("请求不通啊", ballotX, err)
			}
			t.Log("可以请求通！", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("收到响应码%d %v", statusCode, checkMark)
			} else {
				t.Errorf("收到响应码%d %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}