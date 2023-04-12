package godrawer

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strings"

	"github.com/fogleman/gg"
)

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

func (b ConcreteImageBuilder) DrawVerticalGradientColors(solidColor ...string) {
	dc := b.DC

	// Create a new linear gradient with start and end points reversed to make it vertical
	gradient := gg.NewLinearGradient(0, float64(dc.Height()), 0, 0)

	// Calculate the step size based on the number of provided colors
	stepSize := countSteps(float64(len(solidColor)))
	var step float64 = 0

	// Add color stops for each provided color
	for _, sc := range solidColor {
		if sd, ok := presetColors[sc]; ok {
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
		if sd, ok := presetColors[sc]; ok {
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

func (b *ConcreteImageBuilder) Build(name string) image.Image {
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
	fullName := fmt.Sprintf("%v.png", name)
	dc.SavePNG(fullName)

	return dc.Image()
}
