package godrawer

import (
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

func drawSignature(dc *gg.Context, text string) {
	signatureFnot := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	if err := dc.LoadFontFace(signatureFnot, 50); err != nil {
		log.Fatal(err)
	}
	dc.SetColor(color.Black)

	signature := text
	marginX := 50.0
	marginY := -5.0

	textWidth, textHeight := dc.MeasureString(signature)
	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(signature, x, y)
}

func drawSignatureGradient(dc *gg.Context) {
	signatureFont := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	if err := dc.LoadFontFace(signatureFont, 50); err != nil {
		log.Fatal(err)
	}

	// Define the gradient colors for the text
	gradient := gg.NewLinearGradient(0, 0, float64(dc.Width()), 0)
	gradient.AddColorStop(0, color.RGBA{R: 255, G: 100, B: 0, A: 255})
	gradient.AddColorStop(1, color.RGBA{R: 255, G: 200, B: 0, A: 255})

	// Set the fill color to the gradient
	dc.SetFillStyle(gradient)

	// Draw the text with a shadow
	signature := "MohamedAllam.Tech"
	marginX := 50.0
	marginY := -5.0
	shadowX := marginX / 2
	shadowY := marginY / 2

	textWidth, textHeight := dc.MeasureString(signature)
	x := float64(dc.Width()) - textWidth - marginX
	y := float64(dc.Height()) - textHeight - marginY

	// Draw the text shadow
	dc.SetColor(color.Black)
	dc.DrawStringAnchored(signature, x+shadowX, y+shadowY, 1, 1)

	// Draw the gradient text on top of the shadow
	dc.SetColor(color.White)
	dc.DrawStringAnchored(signature, x, y, 1, 1)
}

func drawSignatureGradient2(dc *gg.Context) {
	signatureFont := "/Users/mohamedallam/Library/Fonts/Ubuntu Mono derivative Powerline.ttf"
	if err := dc.LoadFontFace(signatureFont, 50); err != nil {
		log.Fatal(err)
	}

	// Calculate the size of the signature text and the background rectangle
	signatureText := "MohamedAllam.Tech"
	signatureMarginX := 50.0
	signatureMarginY := 50.0
	signatureTextWidth, signatureTextHeight := dc.MeasureString(signatureText)
	backgroundRectWidth := signatureTextWidth + signatureMarginX*2
	backgroundRectHeight := signatureTextHeight + signatureMarginY*2

	// Calculate the position of the background rectangle and the signature text
	x := float64(dc.Width()) - backgroundRectWidth
	y := float64(dc.Height()) - backgroundRectHeight
	signatureTextX := x + signatureMarginX
	signatureTextY := y + signatureMarginY + signatureTextHeight

	// Draw a semi-transparent black background rectangle and a shadow effect
	dc.SetColor(color.RGBA{0, 0, 0, 128})
	// dc.DrawRoundedRectangle(x, y, backgroundRectWidth, backgroundRectHeight, 10)
	dc.Fill()
	dc.SetColor(color.White)
	// dc.SetShadow(10, 10, 10, color.Black)
	dc.DrawString(signatureText, signatureTextX, signatureTextY)
	// dc.SetShadow(0, 0, 0, color.Transparent)
}
