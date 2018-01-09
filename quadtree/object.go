package quadtree

type Object struct {
	x      float32
	y      float32
	width  float32
	height float32
}

func NewObject(_x, _y, _width, _height float32) *Object {
	o := &Object{}
	o.x = _x
	o.y = _y
	o.width = _width
	o.height = _height
	return o
}
