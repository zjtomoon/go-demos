package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(800,600,webview.HintNone)
	w.Navigate("Http://www.baidu.com")
	w.Run()
}
