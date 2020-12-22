package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//func main() {
//	client := &http.Client{}
//	method := "POST"
//	url_eachPage := "http://m.bcoderss.com/tag/%e7%be%8e%e5%a5%b3/page/1/"
//
//	resp, err := Request0(method, url_eachPage, client)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	urls:= GenUrlsFromPage(resp, err)
//
//	for _,each_url:= range urls{
//		fmt.Printf("%s\n",each_url)
//	}
//}

func GenUrlsFromPage(resp *http.Response, err error) (res_urls [][]byte) {
	re1 := regexp.MustCompile(`http://m.bcoderss.com/wp-content/.+\.jpg`)
	re_replace := regexp.MustCompile(`-[0-9x]{5,9}`)

	bodyText, err := ioutil.ReadAll(resp.Body) //这样就是一次性读入到内存,字节slice保存
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	urls := re1.FindAll(bodyText, -1)

	for _, each := range urls { //这里注意使用range函数还会返回索引
		fmt.Printf("[URL]: %s\n", each)
		url := re_replace.ReplaceAll(each, []byte(""))
		res_urls = append(res_urls, url)
	}
	return
}

func Request0(method string, url_eachPage string, client *http.Client) (*http.Response, error) {
	req, err := http.NewRequest(method, url_eachPage, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "m.bcoderss.com")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36 Edg/87.0.664.66")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Origin", "http://m.bcoderss.com")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cookie", "wordpress_logged_in_e2a38b6b918ebfd1f0be80686a411801=crazyhubox%7C1609773021%7Cdx2iWF1YU5gA3Hg5fiwsFrkbuqaL4m2niZ5kzxQKCAq%7Ccc43f7505a15d0731855134d61dd92919157094a8d156f2fee10a5e71ab9a16e")
	resp, err := client.Do(req)
	fmt.Printf("[REQ]:%s[%d]\n",url_eachPage,resp.StatusCode)
	return resp, err
}
