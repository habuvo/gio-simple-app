// SPDX-License-Identifier: Unlicense OR MIT

package main

// A Gio program that demonstrates Gio widgets. See https://gioui.org for more information.

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go update()

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(700), unit.Dp(1200)), app.Title("easy trader"))
		if err := handle(w); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()

	app.Main()
}
