package web

import (
	"encoding/json"
	"fmt"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type OneResult struct {
	Status string           `json:"status"`
	State  string           `json:"state"`
	Data   []*OneResultData `json:"data"`
}

func (o OneResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type OneResultData struct {
	Value int `json:"value"`
}

func OneRequest(url string) (int, error) {
	req, err := Get(url, DefaultCookie)
	if err != nil {
		return 0, err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	result := OneResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func OneCallJs(filePath, method, s string) string {
	//先读入文件内容
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(string(bytes))
	if err != nil {
		panic(err)
	}
	value1, err := vm.Call(method, nil, s)
	if err != nil {
		panic(err)
	}
	return value1.String()
}

func OneUrl(page int, md5, tm string) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/1?page=%d&m=%s丨%s", page, md5, tm)
}

func OneStart() error {
	tm := time.Now().Unix()*1000 + 100000000
	s := OneCallJs("./1.js", "hex_md5", fmt.Sprintf("%d", time.Now().Unix()*1000+100000000))
	count := 0
	for i := 0; i < 5; i++ {
		url := OneUrl(i+1, s, fmt.Sprintf("%d", tm/1000))
		num, err := OneRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count / 50)
	return nil
}

func TestOne(t *testing.T) {
	if err := OneStart(); err != nil {
		t.Fatal("失败")
	}
	t.Log("成功")
}
