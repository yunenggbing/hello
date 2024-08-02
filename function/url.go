package function

import (
	"fmt"
	"image/color"
	"net/url"
)

type Point struct {
	X, Y float64
}
type ColorPoint struct {
	Point
	Color color.RGBA
}
type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
func main() {
	m := url.Values{"lang": {"zh-CN"}}
	m.Add("item", "1")
	m.Add("item", "2")
	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])
	m = nil
	fmt.Println(m.Get("item"))
	m.Add("item", "3")
	//cp := ColorPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
}
