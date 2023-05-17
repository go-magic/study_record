package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

type ThirteenResult struct {
	Status string                `json:"status"`
	State  string                `json:"state"`
	Data   []*ThirteenResultData `json:"data"`
}

func (o ThirteenResult) Value() int {
	count := 0
	for _, data := range o.Data {
		count += data.Value
	}
	return count
}

type ThirteenResultData struct {
	Value int `json:"value"`
}

func ThirteenRequest(url string, cookie string) (int, error) {
	cookie = "sessionid=ypgxvbl02o5u1pggp4rcntaqy9187dey; " + cookie
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
	result := ThirteenResult{}
	if err = json.Unmarshal(b, &result); err != nil {
		return 0, err
	}
	return result.Value(), nil
}

func ThirteenCookie(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("cookie", "sessionid=ypgxvbl02o5u1pggp4rcntaqy9187dey")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return thirteenParseCookie(string(b)), nil
}

func ThirteenUrl(page int) string {
	if page == 1 {
		return "https://match.yuanrenxue.cn/api/match/13"
	}
	return fmt.Sprintf("https://match.yuanrenxue.cn/api/match/13?page=%d", page)
}

func ThirteenStart() error {
	count := 0
	for i := 1; i < 6; i++ {
		url := ThirteenUrl(i)
		cookie, err := ThirteenCookie("https://match.yuanrenxue.cn/match/13")
		if err != nil {
			return err
		}
		num, err := ThirteenRequest(url, cookie)
		if err != nil {
			return err
		}
		count += num
		time.Sleep(time.Second)
	}
	fmt.Println(count)
	return nil
}

func thirteenParseCookie(cookie string) string {
	cookie = strings.Replace(cookie, "<script>", "", -1)
	cookie = strings.Replace(cookie, "</script>", "", -1)
	cookie = strings.Replace(cookie, "document.cookie=", "", -1)
	cookie = strings.Replace(cookie, "('", "", -1)
	cookie = strings.Replace(cookie, "')", "", -1)
	cookie = strings.Replace(cookie, "+", "", -1)
	arr := strings.Split(cookie, "';")
	if len(arr) > 1 {
		cookie = arr[0]
	}
	return ReplaceAll(cookie)
}

func TestThirteen(t *testing.T) {
	if err := ThirteenStart(); err != nil {
		t.Fatal(err)
	}
	t.Log("成功")
}

const thirteenCookie = `
<script>
	document.cookie=('y')+('u')+('a')+('n')+('r')+('e')+('n')+('x')+('u')+('e')+('_')+('c')+('o')+('o')+('k')+('i')+('e')+('=')+('1')+('6')+('8')+('4')+('2')+('9')+('5')+('6')+('9')+('5')+('|')+('N')+('Q')+('p')+('U')+('C')+('J')+('2')+('X')+('X')+('7')+('n')+('D')+('w')+('h')+('H')+('0')+('U')+('e')+('P')+('8')+('q')+('Q')+('o')+('c')+('Q')+('r')+('u')+('m')+('N')+('H')+('W')+('M')+('Z')+('o')+('w')+('a')+('V')+('Q')+('6')+('1')+('T')+('N')+('r')+('I')+('r')+('E')+('Z')+('C')+('G')+('G')+('t')+('6')+('S')+('q')+('o')+('A')+('G')+('G')+('7')+('h')+('x')+('b')+('W')+('W')+('W')+('o')+('o')+('D')+('b')+('y')+('q')+('k')+('9')+('k')+('k')+('q')+('l')+('a')+('D')+('Y')+('a')+('S')+('z')+('T')+('d')+('z')+('9')+('E')+('6')+('K')+('R')+('y')+('A')+('0')+('A')+('a')+('b')+('8')+('T')+('v')+('o')+('G')+('l')+('Z')+('v')+('2')+('z')+('d')+('s')+('H')+('L')+('k')+('9')+('J')+('e')+('x')+('o')+('4')+('L')+('L')+('j')+('N')+('m')+('5')+('o')+('A')+('z')+('S')+('k')+('6')+('2')+('r')+('a')+('2')+('E')+('H')+('J')+('C')+('H')+('U')+('s')+('p')+('e')+('S')+('G')+('q')+('o')+('G')+('q')+('R')+('e')+('T')+('m')+('u')+('d')+('f')+('j')+('n')+('w')+('h')+('v')+('J')+('N')+('F')+('9')+('m')+('5')+('L')+('c')+('i')+('3')+('g')+('V')+('G')+('O')+('k')+('M')+('c')+('9')+('Y')+('x')+('f')+('B')+('a')+('X')+';path=/';location.href=location.pathname+location.search
</script>`

func TestParseCookie(t *testing.T) {
	cookie := strings.Replace(thirteenCookie, "<script>", "", -1)
	cookie = strings.Replace(cookie, "</script>", "", -1)
	cookie = strings.Replace(cookie, "document.cookie=", "", -1)
	cookie = strings.Replace(cookie, "('", "", -1)
	cookie = strings.Replace(cookie, "')", "", -1)
	cookie = strings.Replace(cookie, "+", "", -1)
	arr := strings.Split(cookie, "';")
	if len(arr) > 1 {
		cookie = arr[0]
	}
	cookie = ReplaceAll(cookie)
	t.Log(cookie)
}
