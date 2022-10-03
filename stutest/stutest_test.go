package stutest

import (
	"log"
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// 必须要以Test开头
// 文件也要以_test.go结尾
// 必须要有t *testing.T 参数
func TestDownlod(t *testing.T) {
	url := "https://flowus.cn/product"
	statusCode := 200

	t.Log("\t开始测试下载内容")
	{
		t.Logf("测试下载URL：%s, 期望结果：%d", url, statusCode)
		{
			resp, err := http.Get(url)
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



func Test123(t *testing.T) {
	a := 2
	if a != 3 {
		t.Logf("测试通过！")
		return
	}
	t.Logf("测试不通过")
}