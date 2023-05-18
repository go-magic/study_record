package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type SeventeenResult struct {
	Status string                 `json:"status"`
	State  string                 `json:"state"`
	Data   []*SeventeenResultData `json:"data"`
}

func (o SeventeenResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type SeventeenResultData struct {
	Value int `json:"value"`
}

func SeventeenRequest(url string) (int, error) {
	cookie := "sessionid=6b1uo8ij7y4vic05x5vc4wsazqm01pq6"
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
	result := SeventeenResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func SeventeenUrl(page int) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/17?page=%d", page)
}

func SeventeenStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		url := SeventeenUrl(i)
		num, err := SeventeenRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestSeventeen(t *testing.T) {
	if err := SeventeenStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
