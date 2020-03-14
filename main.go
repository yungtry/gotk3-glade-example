package main

import (
	"github.com/gotk3/gotk3/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)
	filename := "main.glade"

	bld, _ := gtk.BuilderNew()
	bld.AddFromFile(filename)

	obj, _ := bld.GetObject("window")
	window, _ := obj.(*gtk.Window)

	window.SetTitle("GO GTK3 Glade Example")
	window.SetDefaultSize(300, 300)
	window.ShowAll()
	window.Connect("destroy", func() {
		gtk.MainQuit()
	})


	gtk.Main()
}