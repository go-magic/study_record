package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type SixteenResult struct {
	Status string               `json:"status"`
	State  string               `json:"state"`
	Data   []*SixteenResultData `json:"data"`
}

func (o SixteenResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type SixteenResultData struct {
	Value int `json:"value"`
}

func SixteenRequest(url string) (int, error) {
	cookie := "sessionid=t5n1tjxzl5zjd86bb1ranxbwbjwf9snx"
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
	result := SixteenResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func SixteenUrl(page int, s, t string) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/16?page=%d&m=%s&t=%s", page, s, t)
}

func SixteenStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		tm := fmt.Sprintf("%d", time.Now().Unix()*1000)
		s, err := CallJs("./16.js", "dddd", tm)
		url := SixteenUrl(i, s, tm)
		num, err := SixteenRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestSixteen(t *testing.T) {
	if err := SixteenStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
