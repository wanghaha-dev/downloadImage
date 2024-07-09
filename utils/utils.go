package utils

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/wanghaha-dev/downloadImage/common"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Download(imgUrl string) {
	filename := filepath.Join(common.SaveDir, filepath.Base(imgUrl))

	// 检测文件是否存在并且是否损坏
	if fileutil.IsExist(filename) {
		if size, _ := fileutil.FileSize(filename); size == 0 {
			// 文件损坏，重新下载
			log.Println(filename, "文件损坏.重新下载，文件大小为:", size)
		} else {
			// 文件已经存在并且未损坏，则直接忽略
			log.Println(filename, "已经存在.")
			return
		}
	}

	resp, err := http.Get(imgUrl)
	if err != nil {
		WriteErrImgLog(imgUrl)
		return
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		WriteErrImgLog(imgUrl)
		return
	}

	_ = os.WriteFile(filename, buf, os.ModePerm)
	log.Println(filename, "download ok.")
}

func WriteErrImgLog(imgUrl string) {
	file, _ := os.OpenFile("errImgLog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	file.WriteString(imgUrl + "\n")
	file.Close()
}
