package main

import "parse/quan"
//http://m.bcoderss.com/tag/%e6%b8%b8%e6%88%8f/
func main() {
	tag := "%e6%b8%b8%e6%88%8f"
	pageNum := 10
	file_dir := "img_dm"

	quan.DownloadImgs(&file_dir,&tag,&pageNum)
}

