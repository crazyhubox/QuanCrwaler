package main

import "parse/quan"

//http://m.bcoderss.com/tag/%e6%b8%b8%e6%88%8f/
func main() {
	tag := "城市"
	start := 1
	pageNum := 15
	file_dir := "img_city"
	quan.DownloadImgs(&file_dir, &tag, &start , &pageNum)
}
