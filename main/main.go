package main

import (
	u "evysave/utils"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	u.Prod_check()
	widget := widgets.NewQWidget(nil, 0)
	widget.SetWindowTitle("My Window")

	button := widgets.NewQPushButton2("My Button", nil)
	button.ConnectClicked(func(checked bool) {
		widget.SetWindowTitle("Button clicked!")
	})

	layout := widgets.NewQVBoxLayout2(widget)
	layout.AddWidget(button, 0, 0)

	widget.SetLayout(layout)

	widget.Show()

	app.Exec()
}
