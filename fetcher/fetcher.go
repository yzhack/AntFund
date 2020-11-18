package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	cookiejar2 "net/http/cookiejar"
	"strings"
)

var Cookie []*http.Cookie
var Cookiejar *cookiejar2.Jar

// Fetch：统一请求，POST，GET
func Fetch(url, method string, body io.Reader) ([]byte, error) {
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           Cookiejar,
	}
	req, err := http.NewRequest(strings.ToUpper(method), url, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)")
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	Cookie = Cookiejar.Cookies(req.URL)
	//fmt.Println(Cookie)
	bytedata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	//fmt.Println(string(bytedata))
	return bytedata, nil
}
