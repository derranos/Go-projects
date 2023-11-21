package api

type FiguresStruct struct {
	fg1 string
	fg2 string
	val float64
}

var figuresDatabase1 []FiguresStruct
var figuresDatabase2 []FiguresStruct

func AddFiguresTo1(figures FiguresStruct) {
	figuresDatabase1 = append(figuresDatabase1, figures)
}
func AddFiguresTo2(figures FiguresStruct) {
	figuresDatabase2 = append(figuresDatabase2, figures)
}
func GetDatabase1() []FiguresStruct {
	return figuresDatabase1
}
func GetDatabase2() []FiguresStruct {
	return figuresDatabase2
}
