package main

import (
	"fmt"
	"os"

	"github.com/mappu/miqt/qt"
)

func main() {
	app := qt.NewQApplication(len(os.Args), os.Args)
	defer app.Delete()

	window := qt.NewQWidget(nil, 0)
	window.SetWindowTitle("Hello miqt")
	window.Resize2(320, 200)

	btn := qt.NewQPushButton2("Click me", window)
	btn.Move2(110, 80)

	clicks := 0
	btn.OnClicked(func(checked bool) {
		clicks++
		btn.SetText(fmt.Sprintf("Clicked %d times", clicks))
	})

	window.Show()
	app.Exec()
}
