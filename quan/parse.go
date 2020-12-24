package quan

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"parse/tools"
	"path"
)

func DownloadImgs(file_dir *string, tag *string,start *int ,pageNum *int) {
	client := &http.Client{}
	resp_chan := make(chan *http.Response)
	wait := make(chan interface{})

	resp_chan = Rer(start,pageNum, *tag, client, resp_chan)
	Der(resp_chan,file_dir,wait)
	//<-wait
}


func Rer(start *int,pageNum *int, tag string, client *http.Client, resp_chan chan *http.Response) chan *http.Response {
	go func() {
		for i := *start; i < *pageNum+1; i++ {
			pageUrl := fmt.Sprintf(`http://m.bcoderss.com/tag/%s/page/%d/`, tag, i)
			resp, err := tools.Request0("POST", pageUrl, client)
			if err != nil {
				panic("erro in the page request.")
			}
	
			urls := tools.GenUrlsFromPage(resp, err) //生成每一页爬取的urls

			go func() {
				for _, each_url := range urls {
					func(url *[]byte) {
						resp, err := tools.Request0("GET", string(*url), client)
						if err != nil {
							log.Fatal(err)
						}
						resp_chan <- resp
					}(&each_url)
				}
			}()
		}
	}()
	return resp_chan
}

func Der(resp_chan chan *http.Response,dir *string,wait chan interface{}) {
	for each := range resp_chan {
		go func(resp *http.Response) {
			fmt.Println("start download.")
			url := resp.Request.URL.String()
			name := path.Base(url)
			directory := getDirectory(dir)
			img_path := directory +"/" + name
			downloadImg(resp, nil, img_path)
		}(each)
	}

}

func getDirectory(dir *string) string {
	directory := fmt.Sprintf("/Users/tomjack/Pictures/%s", *dir)
	_, err := os.Stat(directory)
	if err != nil {
		err = os.Mkdir(directory, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Make the directory successfully!\n")
		}
	}
	return directory
}


func downloadImg(resp *http.Response, err error, imgPath string) {
	if err != nil {
		//println(err)
		panic(err)
	}

	reader := bufio.NewReaderSize(resp.Body, 32*1024)
	file, err := os.Create(imgPath)
	fmt.Println(imgPath)
	writer := bufio.NewWriter(file)
	written, _ := io.Copy(writer, reader)
	// 输出文件字节大小
	fmt.Printf("Total length: %d[%s]\n", written,resp.Request.URL)
	defer resp.Body.Close()
}
