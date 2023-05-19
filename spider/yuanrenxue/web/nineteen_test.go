package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type NineteenResult struct {
	Status string                `json:"status"`
	State  string                `json:"state"`
	Data   []*NineteenResultData `json:"data"`
}

func (o NineteenResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type NineteenResultData struct {
	Value int `json:"value"`
}

func NineteenRequest(url string) (int, error) {
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
	result := NineteenResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func NineteenUrl(page int) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/19?page=%d", page)
}

func NineteenStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		url := NineteenUrl(i)
		num, err := NineteenRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestNineteen(t *testing.T) {
	if err := NineteenStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
