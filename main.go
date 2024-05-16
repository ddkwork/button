package main

import (
	"button"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("button", func(w *unison.Window) {
		button.New().Layout(w.Content())
	})
}
