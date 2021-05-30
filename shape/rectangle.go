package shape

type Rectangle struct {
	width  int
	height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width: width, height: height}
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
