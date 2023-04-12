package godrawer

type ImageSpecs struct {
	Width           int
	Height          int
	text            string
	signature       string
	fontSize        float64
	BackGroundColor gg.Pattern
}

func (i *ImageSpecs) SetFontSize(fontSize int) *ImageSpecs {
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
