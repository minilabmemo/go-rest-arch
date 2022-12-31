
## go-rest-arch
[![Go Report Card](https://goreportcard.com/badge/github.com/minilabmemo/go-rest-arch)](https://goreportcard.com/report/github.com/minilabmemo/go-rest-arch)


[TODO not fini]
- 這是根據[go-clean-arch](https://github.com/bxcodec/go-clean-arch)產生出來的一套restful server練習．
- Demo 簡單又常見的方法，作為一個template

### reference
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)


### Tools/libraries Used
All version info listed in go.mod
- viper
- zaplog
- swagger UI : http://localhost:8888/swagger/index.html
- Gin/httptest


### start
```
cd cmd/app-core
go run main.go
```

#### docker build 
>在`docker build -t app-core .` 時 
會發現 RUN command 或者makefile print 沒有output顯示，無法幫助debug的話，這時可以在前面加上DOCKER_BUILDKIT=0 也就是
`DOCKER_BUILDKIT=0 docker build -t app-core .`
see https://makeoptim.com/en/tool/docker-build-not-output

