package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func figureMax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var shapeRequest ShapeRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&shapeRequest)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		s1 := sizeFigures(w, shapeRequest.Figure1, shapeRequest.Data1)
		s2 := sizeFigures(w, shapeRequest.Figure2, shapeRequest.Data2)
		if s1 == -1 || s2 == -1 {
			return
		}
		var res float64
		if s1 > s2 {
			res = 1
		} else if s1 < s2 {
			res = 2
		} else {
			res = 0
		}
		ans := FiguresStruct{fg1: shapeRequest.Figure1, fg2: shapeRequest.Figure2, val: res}
		AddFiguresTo1(ans)

	} else if r.Method == "GET" {
		fmt.Fprintln(w, GetDatabase1())
	}
}

func figureSum(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var shapeRequest ShapeRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&shapeRequest)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		s1 := sizeFigures(w, shapeRequest.Figure1, shapeRequest.Data1)
		s2 := sizeFigures(w, shapeRequest.Figure2, shapeRequest.Data2)
		if s1 == -1 || s2 == -1 {
			return
		}
		ans := FiguresStruct{fg1: shapeRequest.Figure1, fg2: shapeRequest.Figure2, val: s1 + s2}
		AddFiguresTo2(ans)

	} else if r.Method == "GET" {
		fmt.Fprintln(w, GetDatabase2())
	}
}
func sizeFigures(w http.ResponseWriter, str string, l float64) float64 {
	switch str {
	case "circle":
		return 3.14 * l * l
	case "square":
		return l * l
	default:
		http.Error(w, "Invalid figure", http.StatusBadRequest)
		return -1
	}
}
