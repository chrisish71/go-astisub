package model

type Subtitle struct {
	VerticalAlign   int    `json:"verticalAlign"`
	HorizontalAlign string `json:"horizontalAlign"`
	StartTime       int64  `json:"startTime"`
	FontSize        string `json:"fontSize"`
	EndTime         int64  `json:"endTime"`
	Lines           []Line `json:"lines"`
}
