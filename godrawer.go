package godrawer

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
)

type ImageSpecs struct {
	Width           int
	Height          int
	text            string
	signature       string
	fontSize        float64
	BackGroundColor gg.Pattern
}

func ImageSpecsFactory(width int, height int) ImageSpecs {
	return ImageSpecs{
		Width:  width,
		Height: height,
	}
}

type ImageBuilder interface {
	SetDimensions(width, height int) ImageBuilder
	SetText(text string) ImageBuilder
	SetColor(color color.Color) ImageBuilder
	Build() image.Image
}

func NewConcreteImageBuilder(specs ImageSpecs) *ConcreteImageBuilder {
	dc := gg.NewContext(specs.Width, specs.Height)
	return &ConcreteImageBuilder{
		Specs: specs,
		DC:    dc,
	}
}

type ConcreteImageBuilder struct {
	Specs ImageSpecs
	DC    *gg.Context
}

func (i *ImageSpecs) SetFont(fontSize int) *ImageSpecs {
	i.fontSize = float64(fontSize)
	return i
}
func (i *ImageSpecs) SetText(text string) *ImageSpecs {
	i.text = text
	return i
}
func (i *ImageSpecs) SetSignature(text string) *ImageSpecs {
	i.signature = text
	return i
}

// func (b *ConcreteImageBuilder) SetDimensions(width, height int) ImageBuilder {
// 	fmt.Println("Seting Dimensions")
// 	b.Specs.Width = width
// 	b.Specs.Height = height
// 	return b
// }

// func (b *ConcreteImageBuilder) SetText(text string) ImageBuilder {
// 	fmt.Println("Setting Text")

//		b.Specs.Text = text
//		return b
//	}
// func (b *ConcreteImageBuilder) SetSignature(text string) ImageBuilder {
// 	fmt.Println("Setting Signature")

// 	b.Specs.Signature = text
// 	return b
// }

// func (b *ConcreteImageBuilder) SetColor(color color.Color) ImageBuilder {
// 	// some := gg.NewSolidPattern(color)
// 	b.Specs.BackGroundColor = gg.NewSolidPattern(color)
// 	// f := fmt.Fprintf("Setting SetColors")
// 	// b.Specs.BackGroundColor = color
// 	// b.Specs.Colors = colors
// 	return b
// }

// func (b *ConcreteImageBuilder) GetsColor() string {
// 	// some := gg.NewSolidPattern(color)
// 	// a := b.Specs.BackGroundColor

// 	// b.Specs.BackGroundColor = gg.NewSolidPattern(color)
// 	fmt.Println("Setting SetColors")
// 	// b.Specs.BackGroundColor = color
// 	// b.Specs.Colors = colors
// 	return "as"
// }

func (b ConcreteImageBuilder) DrawSolidColor(solidColor string) {
	dc := b.DC

	var c color.Color

	if sd, b := presetColors[solidColor]; b {
		c = sd
	}

	dc.SetFillStyle(gg.NewSolidPattern(c))

	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
	dc.Fill()
}

// func (i ConcreteImageBuilder) DrawGradientColors(solidColor ...string) {
// 	dc := i.DC

// 	var c color.Color
// 	for colorrrr := range solidColor {
// 		if sd, b := presetColors[colorrrr]; b {
// 			c = sd
// 		}
// 	}

// 	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
// 	dc.Fill()
// }

func countSteps(length float64) float64 {
	step := 1 / length
	// log.Fatal(step)
	return float64(step)
}

// func (i ConcreteImageBuilder) DrawGradientColors(solidColor ...string) {
// 	dc := i.DC

// 	// Create a new linear gradient
// 	gradient := gg.NewLinearGradient(0, 0, 0, float64(dc.Height()))

// 	// ste
// 	// log.Fatal(len(solidColor))
// 	stepSize := countSteps(float64(len(solidColor)))
// 	var step float64 = 0
// 	// Add color stops for each provided color
// 	for _, sc := range solidColor {
// 		if sd, b := presetColors[sc]; b {
// 			gradient.AddColorStop(step, sd)
// 			step += stepSize
// 		}
// 	}

// 	// Draw the gradient
// 	dc.SetFillStyle(gradient)
// 	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
// 	dc.Fill()
// }

func (i ConcreteImageBuilder) DrawVerticalGradientColors(solidColor ...string) {
	dc := i.DC

	// Create a new linear gradient with start and end points reversed to make it vertical
	gradient := gg.NewLinearGradient(0, float64(dc.Height()), 0, 0)

	// Calculate the step size based on the number of provided colors
	stepSize := countSteps(float64(len(solidColor)))
	var step float64 = 0

	// Add color stops for each provided color
	for _, sc := range solidColor {
		if sd, b := presetColors[sc]; b {
			gradient.AddColorStop(step, sd)
			step += stepSize
		}
	}

	// Draw the gradient
	dc.SetFillStyle(gradient)
	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
	dc.Fill()
}

func (b ConcreteImageBuilder) DrawHorizontalGradientColors(solidColor ...string) {
	dc := b.DC

	// Create a new linear gradient
	gradient := gg.NewLinearGradient(0, float64(dc.Height())/2, float64(dc.Width()), float64(dc.Height())/2)

	// Calculate the step size for each color stop
	stepSize := countSteps(float64(len(solidColor)))
	var step float64 = 0

	// Add color stops for each provided color
	for _, sc := range solidColor {
		if sd, b := presetColors[sc]; b {
			gradient.AddColorStop(step, sd)
			step += stepSize
		}
	}

	// Draw the gradient
	dc.SetFillStyle(gradient)
	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
	dc.Fill()
}

// func (i ConcreteImageBuilder) DrawHorizontalGradientColors(solidColor ...string) {
// 	dc := i.DC

// 	// Create a new linear gradient
// 	gradient := gg.NewLinearGradient(0, 0, 0, float64(dc.Height()))

// 	// ste
// 	// log.Fatal(len(solidColor))
// 	stepSize := countSteps(float64(len(solidColor)))
// 	var step float64 = 0
// 	// Add color stops for each provided color
// 	for _, sc := range solidColor {
// 		if sd, b := presetColors[sc]; b {
// 			gradient.AddColorStop(step, sd)
// 			step += stepSize
// 		}
// 	}

// 	// Draw the gradient
// 	dc.SetFillStyle(gradient)
// 	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
// 	dc.Fill()
// }

func (b *ConcreteImageBuilder) DrawText() {
	dc := b.DC
	text := b.Specs.text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)
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
	fontSize := float64(140)

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

	x := float64(b.Specs.Width) / 2.0
	y := (float64(b.Specs.Height) - totalHeight) / 2.0
	for _, line := range lines {
		dc.DrawStringAnchored(line, float64(x), float64(y), 0.5, 0.5)
		y += lineHeight
	}
	// dc.MeasureMultilineString(lines, )
	// dc.DrawStringWrapped(line)
}

func (b *ConcreteImageBuilder) Build() image.Image {
	// width := b.Specs.Width
	// height := b.Specs.Height
	dc := b.DC
	// drawGradient(dc)
	// drawText(dc, b.Specs.Text, width, height)s
	// b.DrawText()
	// b.DrawTextStringWrapped()
	b.DrawTextStringWrappedSHadow()
	drawSignature(dc, b.Specs.signature)
	// drawSignature(dc)
	dc.SavePNG("gradient1.png")

	return dc.Image()
}
