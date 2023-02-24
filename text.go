package godrawer

import (
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
)

func drawText(dc *gg.Context, text string, width, height int) {
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)

	wrappedText := dc.WordWrap(text, float64(width))
	s := strings.Join(wrappedText, "\n")

	// Calculate the total height of the wrapped text
	lines := strings.Split(s, "\n")
	lineHeight := float64(fontSize) * 1.5 // adjust this value to change the line spacing
	totalHeight := float64(len(lines)) * lineHeight

	// Draw the wrapped text in the center of the image
	x := float64(width) / 2.0
	y := (float64(height) - totalHeight) / 2.0
	for _, line := range lines {
		dc.DrawStringAnchored(line, float64(x), float64(y), 0.5, 0.5)
		y += lineHeight
	}
}
