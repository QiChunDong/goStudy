package stutest

import (
	"log"
	"net/http"
	"testing"
)

// 必须要有t *testing.T 参数
func TestDownlodGroup(t *testing.T) {
	var urls = []struct {
		url string
		statusCode int
	}{
		{
			url: "https://baijiahao.baidu.com/s?id=1731518084417586113",
			statusCode: 200,
		},
		{
			url: "https://juejin.cn/st/7129144623968026637",
			statusCode: 404,
		},
	}

	t.Log("\t开始测试下载不同的内容")
	{
		for _, u := range urls {
			t.Logf("测试下载URL：%s, 期望结果：%d", u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					log.Fatal("请求不通啊", ballotX, err)
				}
				t.Log("可以请求通！", checkMark)
	
				defer resp.Body.Close()
	
				if resp.StatusCode == u.statusCode {
					t.Logf("收到响应码%d %v", u.statusCode, checkMark)
				} else {
					t.Errorf("收到响应码%d %v %v", u.statusCode, ballotX, resp.StatusCode)
				}
			}
		}
	}
}

// go test -v
// === RUN   TestDownlodGroup
//     stutest_group_test.go:25:   开始测试下载不同的内容
//     stutest_group_test.go:28: 测试下载URL：https://baijiahao.baidu.com/s?id=1731518084417586113, 期望结果：200
//     stutest_group_test.go:34: 可以请求通！ ✓
//     stutest_group_test.go:39: 收到响应码200 ✓
//     stutest_group_test.go:28: 测试下载URL：https://juejin.cn/st/7129144623968026637, 期望结果：404
//     stutest_group_test.go:34: 可以请求通！ ✓
//     stutest_group_test.go:39: 收到响应码404 ✓
// --- PASS: TestDownlodGroup (1.35s)