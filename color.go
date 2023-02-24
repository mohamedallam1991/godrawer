package godrawer

import (
	"image/color"
)

// func drawGradient(dc *gg.Context) {
// 	gradient := gg.NewLinearGradient(0, 0, 0, float64(dc.Height()))

// 	gradient.AddColorStop(0, color.CMYK{20, 99, 0, 64})
// 	gradient.AddColorStop(0.35, color.CMYK{0, 72, 41, 42})
// 	gradient.AddColorStop(1, color.CMYK{0, 46, 53, 24})

// 	c := mySolidColor{}

// 	// accepts color.color and rerturn gg.Pattern
// 	d := gg.NewSolidPattern(c)

// 	dc.SetFillStyle(d)
// 	// dc.SetFillStyle(gradient)
// 	// dc.DrawRectangle(0, 0, float64(dc.Width()), float64(dc.Height()))
// 	dc.DrawRoundedRectangle(0, 0, float64(dc.Width()), float64(dc.Height()), 10)
// 	dc.Fill()
// }

type mySolidColor struct {
	Name string
	R    uint8
	G    uint8
	B    uint8
	A    uint8
}

func (c mySolidColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	a = uint32(c.A)
	r |= r << 8
	g |= g << 8
	b |= b << 8
	a |= a << 8
	return r, g, b, a
}

// var presetColors = map[string]color.Color{
// 	"red":     mySolidColor{Name: "red", R: 255, G: 0, B: 0, A: 255},
// 	"green":   mySolidColor{Name: "green", R: 0, G: 255, B: 0, A: 255},
// 	"blue":    mySolidColor{Name: "blue", R: 0, G: 0, B: 255, A: 255},
// 	"yellow":  mySolidColor{Name: "yellow", R: 255, G: 255, B: 0, A: 255},
// 	"magenta": mySolidColor{Name: "magenta", R: 255, G: 0, B: 255, A: 255},
// 	"cyan":    mySolidColor{Name: "cyan", R: 0, G: 255, B: 255, A: 255},
// }

// func NewColor(text string) {
// mySolidColor{Name: "cyan", R: 0, G: 255, B: 255, A: 255},
// }

// func (m mySolidColor) DrawSolid(dc *gg.Context) {
// func (coonc ConcreteImageBuilder) DrawSolid(mySolidColor color.Color) {

// type GradientColor struct {
// 	Gradient *gg.Gradient
// }

// type Gradient interface {
// 	gg.Pattern
// 	AddColorStop(offset float64, color color.Color)
// }

// func (g GradientColor) ColorAt(x, y int) color.Color {
// 	// return g.Gradient.ColorAt(float64(x), float64(y))
// }

// "khaki": mySolidColor{Name: "khaki", R: 240, G:
// "coral":         mySolidColor{Name: "coral", R: 255, G: 127, B: 80, A: 255},

func GetColors() []string {
	// var colors []string
	// fmt.Println("mapsize ", len(presetColors))
	listOfcolors := make([]string, len(presetColors))
	i := 0
	for color := range presetColors {
		listOfcolors[i] = color
		i++
	}
	// fmt.Println("array size ", len(listOfcolors))

	return listOfcolors
}

var presetColors = map[string]color.Color{
	"red":            mySolidColor{Name: "red", R: 255, G: 0, B: 0, A: 255},
	"green":          mySolidColor{Name: "green", R: 0, G: 255, B: 0, A: 255},
	"blue":           mySolidColor{Name: "blue", R: 0, G: 0, B: 255, A: 255},
	"yellow":         mySolidColor{Name: "yellow", R: 255, G: 255, B: 0, A: 255},
	"magenta":        mySolidColor{Name: "magenta", R: 255, G: 0, B: 255, A: 255},
	"cyan":           mySolidColor{Name: "cyan", R: 0, G: 255, B: 255, A: 255},
	"black":          mySolidColor{Name: "black", R: 0, G: 0, B: 0, A: 255},
	"white":          mySolidColor{Name: "white", R: 255, G: 255, B: 255, A: 255},
	"orange":         mySolidColor{Name: "orange", R: 255, G: 165, B: 0, A: 255},
	"purple":         mySolidColor{Name: "purple", R: 128, G: 0, B: 128, A: 255},
	"gray":           mySolidColor{Name: "gray", R: 128, G: 128, B: 128, A: 255},
	"maroon":         mySolidColor{Name: "maroon", R: 128, G: 0, B: 0, A: 255},
	"darkred":        mySolidColor{Name: "darkred", R: 139, G: 0, B: 0, A: 255},
	"brown":          mySolidColor{Name: "brown", R: 165, G: 42, B: 42, A: 255},
	"firebrick":      mySolidColor{Name: "firebrick", R: 178, G: 34, B: 34, A: 255},
	"crimson":        mySolidColor{Name: "crimson", R: 220, G: 20, B: 60, A: 255},
	"tomato":         mySolidColor{Name: "tomato", R: 255, G: 99, B: 71, A: 255},
	"coral":          mySolidColor{Name: "coral", R: 255, G: 127, B: 80, A: 255},
	"indianred":      mySolidColor{Name: "indianred", R: 205, G: 92, B: 92, A: 255},
	"lightcoral":     mySolidColor{Name: "lightcoral", R: 240, G: 128, B: 128, A: 255},
	"salmon":         mySolidColor{Name: "salmon", R: 250, G: 128, B: 114, A: 255},
	"darksalmon":     mySolidColor{Name: "darksalmon", R: 233, G: 150, B: 122, A: 255},
	"lightsalmon":    mySolidColor{Name: "lightsalmon", R: 255, G: 160, B: 122, A: 255},
	"orangina":       mySolidColor{Name: "orange", R: 255, G: 165, B: 0, A: 255},
	"darkorange":     mySolidColor{Name: "darkorange", R: 255, G: 140, B: 0, A: 255},
	"lightyellow":    mySolidColor{Name: "lightyellow", R: 255, G: 255, B: 224, A: 255},
	"lemonchiffon":   mySolidColor{Name: "lemonchiffon", R: 255, G: 250, B: 205, A: 255},
	"papayawhip":     mySolidColor{Name: "papayawhip", R: 255, G: 239, B: 213, A: 255},
	"moccasin":       mySolidColor{Name: "moccasin", R: 255, G: 228, B: 181, A: 255},
	"peachpuff":      mySolidColor{Name: "peachpuff", R: 255, G: 218, B: 185, A: 255},
	"palegoldenrod":  mySolidColor{Name: "palegoldenrod", R: 238, G: 232, B: 170, A: 255},
	"Olive":          mySolidColor{Name: "olive", R: 128, G: 128, B: 0, A: 255},
	"Teal":           mySolidColor{Name: "teal", R: 0, G: 128, B: 128, A: 255},
	"Navy":           mySolidColor{Name: "navy", R: 0, G: 0, B: 128, A: 255},
	"Maroon":         mySolidColor{Name: "maroon", R: 128, G: 0, B: 0, A: 255},
	"Purple":         mySolidColor{Name: "purple", R: 128, G: 0, B: 128, A: 255},
	"Fuchsia":        mySolidColor{Name: "fuchsia", R: 255, G: 0, B: 255, A: 255},
	"Lime":           mySolidColor{Name: "lime", R: 0, G: 255, B: 0, A: 255},
	"Aqua":           mySolidColor{Name: "aqua", R: 0, G: 255, B: 255, A: 255},
	"Silver":         mySolidColor{Name: "silver", R: 192, G: 192, B: 192, A: 255},
	"Gray":           mySolidColor{Name: "gray", R: 128, G: 128, B: 128, A: 255},
	"beige":          mySolidColor{Name: "beige", R: 245, G: 245, B: 220, A: 255},
	"bisque":         mySolidColor{Name: "bisque", R: 255, G: 228, B: 196, A: 255},
	"blanchedalmond": mySolidColor{Name: "blanchedalmond", R: 255, G: 235, B: 205, A: 255},
	"burlywood":      mySolidColor{Name: "burlywood", R: 222, G: 184, B: 135, A: 255},
	"chartreuse":     mySolidColor{Name: "chartreuse", R: 127, G: 255, B: 0, A: 255},
	"chocolate":      mySolidColor{Name: "chocolate", R: 210, G: 105, B: 30, A: 255},
	"cornflowerblue": mySolidColor{Name: "cornflowerblue", R: 100, G: 149, B: 237, A: 255},
	"cornsilk":       mySolidColor{Name: "cornsilk", R: 255, G: 248, B: 220, A: 255},
	"darkcyan":       mySolidColor{Name: "darkcyan", R: 0, G: 139, B: 139, A: 255},
	"darkgoldenrod":  mySolidColor{Name: "darkgoldenrod", R: 184, G: 134, B: 11, A: 255},
	"darkgray":       mySolidColor{Name: "darkgray", R: 169, G: 169, B: 169, A: 255},
	"darkgreen":      mySolidColor{Name: "darkgreen", R: 0, G: 100, B: 0, A: 255},
	"darkkhaki":      mySolidColor{Name: "darkkhaki", R: 189, G: 183, B: 107, A: 255},
	"darkmagenta":    mySolidColor{Name: "darkmagenta", R: 139, G: 0, B: 139, A: 255},
	"darkolivegreen": mySolidColor{Name: "darkolivegreen", R: 85, G: 107, B: 47, A: 255},
	"darkorchid":     mySolidColor{Name: "darkorchid", R: 153, G: 50, B: 204, A: 255},
	// "darksalmon":     mySolidColor{Name: "darksalmon", R: 233, G: 150, B: 122, A: 255},
	"darkseagreen":  mySolidColor{Name: "darkseagreen", R: 143, G: 188, B: 143, A: 255},
	"darkslateblue": mySolidColor{Name: "darkslateblue", R: 72, G: 61, B: 139, A: 255},
	"darkslategray": mySolidColor{Name: "darkslategray", R: 47, G: 79, B: 79, A: 255},
	"darkturquoise": mySolidColor{Name: "darkturquoise", R: 0, G: 206, B: 209, A: 255},
}
