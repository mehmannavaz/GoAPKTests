package main

import (
	"fmt"
	"os"

	"github.com/mappu/miqt/qt"
)

func main() {
	qt.NewQApplication(os.Args)

	window := qt.NewQWidget(nil)
	window.SetWindowTitle("Hello miqt")
	window.Resize(320, 200)

	btn := qt.NewQPushButton()
	btn.SetText("Click me")
	btn.SetParent(window)
	btn.Move(110, 80)

	clicks := 0
	btn.OnClicked(func() {
		clicks++
		btn.SetText(fmt.Sprintf("Clicked %d times", clicks))
	})

	window.Show()
	qt.QApplication_Exec()
}
