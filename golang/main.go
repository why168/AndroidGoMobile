package github

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"simple/model"
)

// GOOS=linux GOARCH=arm64 go build -o server.so
// go tool dist list
// android/386
// android/amd64
// android/arm
// android/arm64

// gomobile init -ndk /Users/edwin/Library/Android/sdk/ndk-bundle/
// echo $ANDROID_HOME
// export ANDROID_HOME=/Users/edwin/Library/Android/sdk/ndk-bundle
// gomobile bind -target=android
// gomobile bind [-target android|ios] [-bootclasspath <path>] [-classpath <path>] [-o output] [build flags] [package]
//

// gomobile build -target=android golang.org/x/mobile/example/basic
// gomobile build -target=android /Users/edwin/gopath/src/golang.org/x/mobile/example/basic

// golang.org/x/mobile/cmd/gobind
// golang.org/x/mobile/cmd/gomobile

//func main() {
//	StartWeb()
//}

func StartWeb() {
	http.HandleFunc("/", getIndex)
	http.HandleFunc("/hello", getHello)
	http.ListenAndServe(":5555", nil)
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	res := model.BaseResponse{
		Code:    1,
		Message: "请求成功",
		Data:    "This is a golang web server!",
	}
	w.WriteHeader(200)
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hello Golang!")
	res := model.BaseResponse{
		Code:    1,
		Message: "请求成功",
		Data:    "Hello，AndroidGoMobile！",
	}
	w.WriteHeader(200)
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}

func Add(x int32, y int32) int32 {
	return x + y
}

func Md5(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	//将[]byte转成16进制
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
