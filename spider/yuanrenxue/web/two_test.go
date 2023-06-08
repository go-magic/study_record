package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type TwoResult struct {
	Status string           `json:"status"`
	State  string           `json:"state"`
	Data   []*TwoResultData `json:"data"`
}

func (o TwoResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type TwoResultData struct {
	Value int `json:"value"`
}

func TwoRequest(url string) (int, error) {
	info, err := CallJs("./2.js", "_0x12f178")
	if err != nil {
		return 0, err
	}
	cookie := "sessionid=nq9ot014fgpkt5nzz9mg4dnexzsm4all" + ";" + info
	req, err := Get(url, cookie)
	if err != nil {
		return 0, err
	}
	req.Header.Set("User-Agent", "yuanrenxue.project")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	result := TwoResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func TwoUrl(page int) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/2?page=%d", page)
}

func TwoStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		url := TwoUrl(i)
		num, err := TwoRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestTwo(t *testing.T) {
	if err := TwoStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
