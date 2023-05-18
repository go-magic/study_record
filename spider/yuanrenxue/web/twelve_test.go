package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type TwelveResult struct {
	Status string              `json:"status"`
	State  string              `json:"state"`
	Data   []*TwelveResultData `json:"data"`
}

func (o TwelveResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type TwelveResultData struct {
	Value int `json:"value"`
}

func TwelveRequest(url string) (int, error) {
	cookie := "sessionid=1ydtzhf6xhyjz4bx4ozlmazt5y6icccz"
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
	result := TwelveResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func TwelveUrl(page int, s string) string {
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/12?page=%d&m=%s", page, s)
}

func TwelveStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		s := fmt.Sprintf("yuanrenxue%d", i)
		url := TwelveUrl(i, Encode(s))
		num, err := TwelveRequest(url)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func TestTwelve(t *testing.T) {
	if err := TwelveStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}
