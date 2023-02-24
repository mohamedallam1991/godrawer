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
	Text            string
	Signature       string
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

type ConcreteImageBuilder struct {
	Specs ImageSpecs
	DC    *gg.Context
}

func NewConcreteImageBuilder(specs ImageSpecs) *ConcreteImageBuilder {
	dc := gg.NewContext(specs.Width, specs.Height)
	return &ConcreteImageBuilder{
		Specs: specs,
		DC:    dc,
	}
}

func (b *ConcreteImageBuilder) SetDimensions(width, height int) ImageBuilder {
	fmt.Println("Seting Dimensions")
	b.Specs.Width = width
	b.Specs.Height = height
	return b
}

func (b *ConcreteImageBuilder) SetText(text string) ImageBuilder {
	fmt.Println("Setting text")

	b.Specs.Text = text
	return b
}
func (b *ConcreteImageBuilder) SetSignature(text string) ImageBuilder {
	fmt.Println("Setting text")

	b.Specs.Signature = text
	return b
}

func (b *ConcreteImageBuilder) SetColor(color color.Color) ImageBuilde	r {
	// some := gg.NewSolidPattern(color)
	b.Specs.BackGroundColor = gg.NewSolidPattern(color)
	f := fmt.Fprintf("Setting SetColors")
	// b.Specs.BackGroundColor = color
	// b.Specs.Colors = colors
	return b
}

func (b *ConcreteImageBuilder) GetsColor() string {
	// some := gg.NewSolidPattern(color)
	a := b.Specs.BackGroundColor

	// b.Specs.BackGroundColor = gg.NewSolidPattern(color)
	fmt.Println("Setting SetColors")
	// b.Specs.BackGroundColor = color
	// b.Specs.Colors = colors
	return "as"
}

func (i ConcreteImageBuilder) DrawSolidColor(solidColor string) {
	dc := i.DC

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

func (i ConcreteImageBuilder) DrawGradientColors(solidColor ...string) {
	dc := i.DC

	// Create a new linear gradient
	gradient := gg.NewLinearGradient(0, 0, 0, float64(dc.Height()))

	// ste
	// log.Fatal(len(solidColor))
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
	text := b.Specs.Text
	fmt.Println("heyyydrawText ")
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)
	height := b.Specs.Height
	width := b.Specs.Width
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

func (b *ConcreteImageBuilder) Build() image.Image {
	// width := b.Specs.Width
	// height := b.Specs.Height
	dc := b.DC
	// drawGradient(dc)
	// drawText(dc, b.Specs.Text, width, height)s
	b.DrawText()
	drawSignature(dc, b.Specs.Signature)
	// drawSignature(dc)
	dc.SavePNG("gradient1.png")

	return dc.Image()
}

func (b *ConcreteImageBuilder) Builda() image.Image {

	fmt.Println("image is going to be built")
	dc := gg.NewContext(b.Specs.Width, b.Specs.Height)

	return dc.Image()

	dc = gg.NewContext(b.Specs.Width, b.Specs.Height)

	gradient := gg.NewLinearGradient(0, 0, 0, float64(b.Specs.Height))
	// for i, c := range b.Specs.Colors {
	// 	gradient.AddColorStop(float64(i)/float64(len(b.Specs.Colors)-1), c)
	// }

	dc.SetFillStyle(gradient)
	dc.DrawRectangle(0, 0, float64(b.Specs.Width), float64(b.Specs.Height))
	dc.Fill()

	fontPath := "/path/to/font.ttf"
	fontSize := float64(72)
	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)

	wrappedText := dc.WordWrap(b.Specs.Text, float64(b.Specs.Width))
	s := strings.Join(wrappedText, "\n")
	lineHeight := float64(fontSize) * 1.5
	totalHeight := float64(len(strings.Split(s, "\n"))) * lineHeight
	x := float64(b.Specs.Width) / 2.0
	y := (float64(b.Specs.Height) - totalHeight) / 2.0
	for _, line := range strings.Split(s, "\n") {
		dc.DrawStringAnchored(line, x, y, 0.5, 0.5)
		y += lineHeight
	}

	return dc.Image()
}

func draaaawing(text string) image.Image {
	width := 1080
	height := 1920
	// height := 1080
	dc := gg.NewContext(width, height)
	gradient := gg.NewLinearGradient(0, 0, 0, float64(height))

	// gradient := gg.NewRadialGradient(0, 0, 0, float64(height), float64(height), float64(height))
	// var radius float64 = float64(width) / 2
	// gradient := gg.NewRadialGradient(float64(width)/2, float64(height)/2, 60, float64(width)/2, float64(height)/2, radius)
	// gg.NewRadialGradient()
	// rgba()
	// color.RGBA64Model
	// a :=
	// color.CMYKModel
	// this are what I like
	// gradient.AddColorStop(0, color.CMYK{91, 70, 0, 83})
	// gradient.AddColorStop(0.35, color.CMYK{50, 70, 0, 45})
	// gradient.AddColorStop(0.7, color.CMYK{50, 70, 0, 45})
	// gradient.AddColorStop(1, color.CMYK{34, 98, 0, 16})
	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})

	// gradient.AddColorStop(0, color.RGBA{4, 13, 44, 1})
	// gradient.AddColorStop(0.5, color.RGBA64{70, 42, 139, 1})
	// gradient.AddColorStop(1, color.RGBA64{141, 5, 214, 1})
	// color.RGBA64
	// gradient.AddColorStop(0, color.RGBA{0, 0, 255, 255})
	// gradient.AddColorStop(0.5, color.RGBA{0, 0, 0, 0})
	// gradient.AddColorStop(1, color.RGBA{0, 255, 0, 255})
	dc.SetFillStyle(gradient)
	// dc.GetCurrentPoint() /2
	// dc.DrawRoundedRectangle(20, 20, float64(width)-20, float64(height)-20, 20)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)
	fmt.Println("text length", len(text))

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
		// dc.DrawString(line, x, y)
		y += lineHeight
	}

	// text := "hey how is it going my man"
	// textWidth, textHeight := dc.MeasureString(text)
	// textX := (float64(width) - textWidth) / 2
	// textY := (float64(height) - textHeight) / 2
	// dc.DrawString(text, textX, textY)

	// fontPath := filepath.Join("fonts", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.Black)
	signature := "MohamedAllam.Tech"
	// marginX := 50.0
	marginX := 50.0
	marginY := -5.0
	textWidth, textHeight := dc.MeasureString(signature)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.DrawString(signature, x, y)

	dc.SavePNG("gradient.png")

	return dc.Image()
}

// func (q quote) MessagePrint() {
// 	input := string(q)
// 	switch {
// 	case strings.HasPrefix(input, "Story:"):
// 		s, _ := strings.CutPrefix(input, "Story:")
// 		q.drawFace(s)
// 	case strings.HasPrefix(input, "Post:"):
// 		s, _ := strings.CutPrefix(input, "Post:")
// 		q.drawFace(s)
// 	default:
// 		log.Fatal("not input matched")
// 	}
// 	// fmt.Println(q)
// }

func drawFace(text string) image.Image {

	width := 1080
	height := 1920
	switch text {
	case "Story":
		height = 1920
		// s, _ := strings.CutPrefix(input, "Story:")
		// drawFacebookStory(s)
	case "Post":
		height = 1080
		// s, _ := strings.CutPrefix(input, "Post:")
		// drawFacebookPost(s)
	default:
		// log.Fatal("not input matched")
	}
	dc := gg.NewContext(width, height)
	gradient := gg.NewLinearGradient(0, 0, 0, float64(height))

	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})

	dc.SetFillStyle(gradient)

	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)
	fmt.Println("text length", len(text))

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
	dc = makeSignature(dc)

	dc.SavePNG("gradient.png")

	return dc.Image()
}

func drawFacebookStory(text string) image.Image {
	width := 1080
	height := 1920
	dc := gg.NewContext(width, height)
	gradient := gg.NewLinearGradient(0, 0, 0, float64(height))

	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})

	dc.SetFillStyle(gradient)

	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)
	fmt.Println("text length", len(text))

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
		// dc.DrawString(line, x, y)
		y += lineHeight
	}

	// if err := dc.LoadFontFace(fontPath, 50); err != nil {
	// 	log.Fatal(err)
	// }
	// dc.SetColor(color.Black)
	// signature := "MohamedAllam.Tech"
	// // marginX := 50.0
	// marginX := 50.0
	// marginY := -5.0
	// textWidth, textHeight := dc.MeasureString(signature)
	// x = float64(dc.Width()) - textWidth - marginX
	// y = float64(dc.Height()) - textHeight - marginY
	// dc.DrawString(signature, x, y)
	dc = makeSignature(dc)

	dc.SavePNG("gradient.png")

	return dc.Image()
}

func drawFacebookPost(text string) image.Image {
	width := 1080
	height := 1080
	dc := gg.NewContext(width, height)
	gradient := gg.NewLinearGradient(0, 0, 0, float64(height))

	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})

	dc.SetFillStyle(gradient)

	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := float64(72)

	err := dc.LoadFontFace(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.White)
	fmt.Println("text length", len(text))

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
		// dc.DrawString(line, x, y)
		y += lineHeight
	}
	// signatureFnot := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	// if err := dc.LoadFontFace(signatureFnot, 50); err != nil {
	// 	log.Fatal(err)
	// }
	// dc.SetColor(color.Black)
	// signature := "MohamedAllam.Tech"
	// marginX := 50.0
	// marginY := -5.0
	// textWidth, textHeight := dc.MeasureString(signature)
	// x = float64(dc.Width()) - textWidth - marginX
	// y = float64(dc.Height()) - textHeight - marginY
	// dc.DrawString(signature, x, y)
	dc = makeSignature(dc)
	dc.SavePNG("gradient.png")

	return dc.Image()
}

func makeSignature(dc *gg.Context) *gg.Context {
	signatureFnot := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	if err := dc.LoadFontFace(signatureFnot, 30); err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.Black)
	signature := "MohamedAllam.Tech"
	marginX := 50.0
	marginY := -5.0
	textWidth, textHeight := dc.MeasureString(signature)
	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(signature, x, y)
	return dc
}

func drawFacebookPost2(text string) image.Image {
	// Define the size of the image
	width := 1080
	height := 1080

	// Create a new drawing context
	dc := gg.NewContext(width, height)

	// Define the gradient colors and add them to the context
	gradient := gg.NewLinearGradient(0, 0, 0, float64(height))
	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})
	dc.SetFillStyle(gradient)

	// Draw the gradient background
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.Fill()

	// Load the font and set the font size
	fontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	fontSize := 72
	if err := dc.LoadFontFace(fontPath, float64(fontSize)); err != nil {
		log.Fatal(err)
	}

	// Set the text color and wrap the text to fit within the image
	dc.SetColor(color.White)
	wrappedText := dc.WordWrap(text, float64(width))
	s := strings.Join(wrappedText, "\n")

	// Calculate the total height of the wrapped text and center it vertically
	lines := strings.Split(s, "\n")
	lineHeight := float64(fontSize) * 1.5 // adjust this value to change the line spacing
	totalHeight := float64(len(lines)) * lineHeight
	x := float64(width) / 2.0
	y := (float64(height) - totalHeight) / 2.0

	// Draw the wrapped text in the center of the image
	for _, line := range lines {
		dc.DrawStringAnchored(line, x, y, 0.5, 0.5)
		y += lineHeight
	}

	// Load the signature font and draw the signature text in the bottom-right corner
	signatureFontPath := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	signatureFontSize := 50
	if err := dc.LoadFontFace(signatureFontPath, float64(signatureFontSize)); err != nil {
		log.Fatal(err)
	}
	signatureText := "MohamedAllam.Tech"
	marginX := 50.0
	marginY := -5.0
	textWidth, textHeight := dc.MeasureString(signatureText)
	x = float64(dc.Width()) - textWidth - marginX
	y = float64(dc.Height()) - textHeight - marginY
	dc.SetColor(color.Black)
	dc.DrawString(signatureText, x, y)

	// Save the image as a PNG file
	if err := dc.SavePNG("gradient.png"); err != nil {
		log.Fatal(err)
	}

	return dc.Image()
}
