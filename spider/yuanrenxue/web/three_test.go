package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type ThreeResult struct {
	Status string             `json:"status"`
	State  string             `json:"state"`
	Data   []*ThreeResultData `json:"data"`
}

func (o ThreeResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type ThreeResultData struct {
	Value int `json:"value"`
}

func ThreeRequest(url string) (int, error) {
	cookie := "sessionid=sqf1foifk9w5dtwdm5x2kxyjmafh1nur"
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
	result := ThreeResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func ThreePost(url string) error {
	var buf *bytes.Buffer
	b, err := json.Marshal(nil)
	if err != nil {
		return err
	}
	buf = bytes.NewBuffer(b)
	req, _ := http.NewRequest(http.MethodPost, url, buf)
	threeHeader(req)

	client := http.Client{
		Transport: http.DefaultTransport,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	return nil
}

func ThreeUrl(page int) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/3?page=%d", page)
}

func ThreeStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		url := ThreeUrl(i)
		if err := ThreePost("https://match.yuanrenxue.cn/jssm"); err != nil {
			return err
		}
		num, err := ThreeRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestThree(t *testing.T) {
	if err := ThreeStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
