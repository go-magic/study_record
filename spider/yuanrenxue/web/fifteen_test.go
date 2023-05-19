package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type FifteenResult struct {
	Status string               `json:"status"`
	State  string               `json:"state"`
	Data   []*FifteenResultData `json:"data"`
}

func (o FifteenResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type FifteenResultData struct {
	Value int `json:"value"`
}

func FifteenRequest(url string) (int, error) {
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
	result := FifteenResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func FifteenUrl(page int, m string) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/15?page=%d&m=%s", page, m)
}

func GetFifteen() (string, error) {
	return CallJs("./a.js", "loadWebAssembly", []byte{0, 1})
}

func FifteenStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		m, err := GetFifteen()
		if err != nil {
			return err
		}
		url := FifteenUrl(i, m)
		num, err := FifteenRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestFifteen(t *testing.T) {

	if err := FifteenStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
