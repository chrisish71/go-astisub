package model

type Line struct {
	Text            string `json:"text"`
	Color           string `json:"color"`
	HorizontalAlign string `json:"horizontalAlign"`
	FontFamily      string `json:"fontFamily"`
	Bold            bool   `json:"bold"`
	Italic          bool   `json:"italic"`
	Underline       bool   `json:"underline"`
	BoxingColor     string `json:"boxingColor"`
	BoxingOpacity   int    `json:"boxingOpacity"`
}
