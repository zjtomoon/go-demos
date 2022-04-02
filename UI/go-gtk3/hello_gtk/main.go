package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

func main() {
	gtk.Init(nil)

	win,err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:",err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	l,err := gtk.LabelNew("Hello,gtk3!")
	if err != nil {
		log.Fatal("Unable to create label:",err)
	}

	win.Add(l)
	win.SetDefaultSize(800,600)
	win.ShowAll()
	gtk.Main()
}