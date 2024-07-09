package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/wanghaha-dev/downloadImage/common"
	"github.com/wanghaha-dev/downloadImage/utils"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	flag.IntVar(&common.GoCount, "count", 50, "goroutine count default 50")
	flag.StringVar(&common.ImgFile, "img", "images.txt", "images file default images.txt")
	flag.StringVar(&common.SaveDir, "dir", "", "saveDir")

	flag.Parse()

	if !fileutil.IsExist(common.SaveDir) {
		fileutil.CreateDir(common.SaveDir)
	}

	wg := new(sync.WaitGroup)
	dataCh := make(chan string, 10000000)

	wg.Add(common.GoCount)
	for i := 0; i < common.GoCount; i++ {
		go Consumer(dataCh, wg)
	}

	Producer(dataCh)

	wg.Wait()
	log.Println("=========================== app finish ===========================")

	var tmp string
	fmt.Scan(&tmp)
}

func Producer(dataCh chan string) {
	f, err := os.Open(common.ImgFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		imgUrl := string(line)
		imgUrl = strings.TrimSpace(imgUrl)

		// 生产数据，数据加入队列
		dataCh <- imgUrl
	}

	close(dataCh)
	log.Println("Data push finish, dataCh close.")
}

func Consumer(dataCh chan string, wg *sync.WaitGroup) {
	for data := range dataCh {
		log.Println("管道剩余:", len(dataCh))
		utils.Download(data)
	}

	wg.Done()
}
