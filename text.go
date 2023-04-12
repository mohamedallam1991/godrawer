package godrawer

import (
	"fmt"
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

func (b *ConcreteImageBuilder) DrawTextStringWrappedSHadow() {
	dc := b.DC
	text := b.Specs.text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := b.Specs.fontSize

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)

	// Draw the wrapped text in the center of the image
	x := float64(b.Specs.Width) / 2.0
	y := float64(b.Specs.Height) / 2.0
	width := float64(b.Specs.Width) - 20
	lineSpacing := 1.5 // adjust this value to change the line spacing
	align := gg.AlignCenter

	// Add shadow to the text
	dc.SetColor(color.Black)

	// Draw a shadow with an offset before drawing the actual text
	dc.SetHexColor("#000000")
	// dc.DrawStringWrapped(text, x+3, y+3, 0.5, 0.5, width, lineSpacing, align)
	dc.DrawStringWrapped(text, x+10, y+10, 0.5, 0.5, width, lineSpacing, align)

	// Draw the actual text
	dc.SetHexColor("#FFFFFF")
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, width, lineSpacing, align)
}

func (b *ConcreteImageBuilder) DrawTextStringWrapped() {
	dc := b.DC
	text := b.Specs.text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := b.Specs.fontSize

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)

	// Draw the wrapped text in the center of the image
	x := float64(b.Specs.Width) / 2.0
	y := float64(b.Specs.Height) / 2.0
	width := float64(b.Specs.Width)
	lineSpacing := 1.5 // adjust this value to change the line spacing
	align := gg.AlignCenter
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, width, lineSpacing, align)
}

func (b *ConcreteImageBuilder) DrawTextAnchored() {
	dc := b.DC
	text := b.Specs.text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := b.Specs.fontSize
	// height := b.Specs.Height
	// width := b.Specs.Width
	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	a := GetColors()[9]
	dc.SetColor(presetColors[a])

	wrappedText := dc.WordWrap(text, float64(b.Specs.Width))
	s := strings.Join(wrappedText, "\n")

	// Calculate the total height of the wrapped text
	lines := strings.Split(s, "\n")
	lineHeight := float64(fontSize) * 1.5 // adjust this value to change the line spacing
	totalHeight := float64(len(lines)) * lineHeight

	// Draw the wrapped text in the center of the image
	x := float64(b.Specs.Width) / 2.0
	y := (float64(b.Specs.Height) - totalHeight) / 2.0
	for _, line := range lines {
		dc.DrawStringAnchored(line, float64(x), float64(y), 0.5, 0.5)
		y += lineHeight
	}
}
func (b *ConcreteImageBuilder) DrawTextMultiline() {
	dc := b.DC
	text := b.Specs.text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := b.Specs.fontSize

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)

	wrappedText := dc.WordWrap(text, float64(b.Specs.Width))
	s := strings.Join(wrappedText, "\n")

	lines := strings.Split(s, "\n")
	lineHeight := float64(fontSize) * 1.5 // adjust this value to change the line spacing
	totalHeight := float64(len(lines)) * lineHeight
	// dc.MeasureMultilineString(lines)

	x := float64(b.Specs.Width) / 2.0
	y := (float64(b.Specs.Height) - totalHeight) / 2.0
	for _, line := range lines {
		dc.DrawStringAnchored(line, float64(x), float64(y), 0.5, 0.5)
		y += lineHeight
	}
	// dc.DrawStringWrapped(line)
}
