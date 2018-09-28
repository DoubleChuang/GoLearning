# Golang 品質驗證
* [golangci](https://golangci.com/)
* [go-critic](https://go-critic.github.io/overview)

* [![Go驗證教學](http://img.youtube.com/vi/lXzQ8ZHUpPY/0.jpg)](http://www.youtube.com/watch?v=lXzQ8ZHUpPY "Go驗證教學")
# Go Web
* [iris](https://iris-go.com/)
* [beego](https://beego.me/)

# Go Compile
* Cross-compile linux program with powershell in VS Code on Windows system
`
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -ldflags "-s -w" -o BIN SRC.go
`
* Cross-compile windows program on linux system
`
GOOS=windows GOARCH=amd64 
go build -ldflags "-s -w" -o BIN SRC.go
`
* Cross-compile Arm program on linux system
`
GOOS=linux GOARCH=arm 
go build -ldflags "-s -w" -o BIN SRC.go
`
