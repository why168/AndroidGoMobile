# AndroidGoMobile
1. go交叉编译打包成arr 
	1. 目的用于学习，不光C、C++可以，Go也可以
	2. 缺点编译出来的so文件太大了（大坑）
2. 运行startWeb()
	1. 启动webService之后很卡，但是不影响访问
	2. 本机浏览器输入：http://localhost:5555/、http://localhost:5555/hello

### .bash_profile   go macOS config
```shell
# go
export GOROOT=/Users/edwin/go
export GOPATH=/Users/edwin/gopath
export GOBIN=/Users/edwin/go/bin
# export GOPROXY=https://goproxy.io
export PATH="$GOBIN:$PATH"
export GOARCH=amd64
export GOOS=android
```

<br/>

###  go bind
```shell
Usage:
    gomobile command [arguments]

Commands:
    bind        build a library for Android and iOS
    build       compile android APK and iOS app
    clean       remove object files and cached gomobile files
    init        build OpenAL for Android
    install     compile android APK and install on device
    version     print version

$ go get golang.org/x/mobile/cmd/gomobile
$ go get -d golang.org/x/mobile/example/basic
$ gomobile clean
$ gomobile init
$ gomobile build -target=android golang.org/x/mobile/example/basic
$ gomobile install golang.org/x/mobile/example/basic

$ gomobile build -target=ios golang.org/x/mobile/example/basic
```
<br/>

# Doc
* https://github.com/golang/go/wiki/Mobile

<br/>
<br/>

# MIT License

```
Copyright (c) 2019 Edwin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
