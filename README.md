# 批量下载图片工具

```sh
go run main.go -h

Usage 
  -count int                                                 
        goroutine count default 50 (default 50)              
  -dir string                                                
        saveDir                                              
  -img string                                                
        images file default images.txt (default "images.txt")
```


`images.txt`
```
https://images.metmuseum.org/CRDImages/ad/original/204788.jpg
https://images.metmuseum.org/CRDImages/ad/original/37808.jpg
https://images.metmuseum.org/CRDImages/ad/original/174118.jpg
https://images.metmuseum.org/CRDImages/ad/original/172134.jpg
https://images.metmuseum.org/CRDImages/ad/original/172134.jpg
......
```