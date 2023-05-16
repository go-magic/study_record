package web

import (
	"bytes"
	"encoding/json"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
)

// DefaultCookie 从浏览器中复制
var DefaultCookie = `Hm_lvt_c99546cf032aaa5a679230de9a95c7db=1684201171; Hm_lvt_9bcbda9cbf86757998a2339a0437208e=1684201182; tk=-5931839558260948860; sessionid=a5jfokvev05jdcby4p8cuwqnplcef121; Hm_lpvt_9bcbda9cbf86757998a2339a0437208e=1684218004; Hm_lpvt_c99546cf032aaa5a679230de9a95c7db=1684219485`

func Get(url, cookie string) (*http.Request, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	defaultHeader(cookie, req)
	return req, nil
}

func Post(url, cookie string, body interface{}) (*http.Request, error) {
	var buf *bytes.Buffer
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	buf = bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, url, buf)
	defaultHeader(cookie, req)
	return req, nil
}

func defaultHeader(cookie string, req *http.Request) {
	req.Header.Add("Cookie", cookie)
	req.Header.Add("Authority", "match.yuanrenxue.cn")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cache-Control", "no-cache")

	req.Header.Add("Sec-Ch-Ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", "Windows")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")

	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
}

func CallJs(filePath, method, s string) (string, error) {
	//先读入文件内容
	bt, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	vm := otto.New()

	_, err = vm.Run(string(bt))
	if err != nil {
		return "", err
	}
	value1, err := vm.Call(method, nil, s)
	if err != nil {
		return "", err
	}
	return value1.String(), nil
}
