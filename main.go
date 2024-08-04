package main

import (
	"button"
	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("button", func(w *unison.Window) {
		w.Content().AddChild(button.New().Layout())
	})
}
