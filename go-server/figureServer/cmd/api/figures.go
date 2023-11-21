package api

type ShapeRequest struct {
	Figure1 string  `json:"fg1"`
	Figure2 string  `json:"fg2"`
	Data1   float64 `json:"sz1"`
	Data2   float64 `json:"sz2"`
}
