package main

import "fmt"
import "github.com/valyala/fasthttp"

func httpGet(url string) {
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
	}

	fmt.Println(string(resp))

}

func main() {
	httpGet("http://www.baidu.com")
}
