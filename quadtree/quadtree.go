/*
//四叉树节点类，用头节点代表四叉树
//坐标系坐上角为原点，左往右为x轴递增，上往下y轴递增
//本四叉树的策略是：1，插入时动态分配节点和删除节点，不是满树；2，当矩形区域完全包含某个节点时才获取或剔除；3，对象放在完全包含它的区域节点内，非根节点也存储对象
*/
package quadtree

//四叉树类型枚举
const (
	ROOT        int8 = iota //根
	UPRIGHT                 //象限Ⅰ
	UPLEFT                  //象限Ⅱ
	BOTTOMLEFT              //象限Ⅲ
	BOTTOMRIGHT             //象限Ⅳ
)

//QuadTreeNode ...
type QuadTreeNode struct {
	//父、子节点，分四个象限
	parent          *QuadTreeNode
	upRightNode     *QuadTreeNode
	upLeftNode      *QuadTreeNode
	bottomLeftNode  *QuadTreeNode
	bottomRightNode *QuadTreeNode
	objects         []*Object // 管理对象
	//节点类型
	quadType int8
	//坐标和长宽属性，左上角为锚点
	x      float32
	y      float32
	width  float32
	height float32

	level    int //当前深度
	maxLevel int //最大深度
}

//NewQuadTreeNode ...
func NewQuadTreeNode(_x, _y, _width, _height float32, _level, _maxLevel int, _quadType int8, _parent *QuadTreeNode) *QuadTreeNode {
	a := &QuadTreeNode{}
	a.x = _x
	a.y = _y
	a.width = _width
	a.height = _height
	a.level = _level
	a.maxLevel = _maxLevel
	a.quadType = _quadType
	a.parent = _parent
	a.objects = nil
	return a
}

//IsContain 判断某个区域是否包含某对象
func (q *QuadTreeNode) IsContain(px, py, w, h float32, obj *Object) bool {
	if obj.x == px && obj.x+obj.width <= px+w && obj.y >= py && obj.y+obj.height <= py+h {
		return true
	}
	return false
}

//IsContain 判断某个区域是否包含某对象
func (q *QuadTreeNode) IsContainX(px, py, w, h float32, qnode *QuadTreeNode) bool {
	if qnode.x == px && qnode.x+qnode.width <= px+w && qnode.y >= py && qnode.y+qnode.height <= py+h {
		return true
	}
	return false
}

//InsertObject 插入对象
func (q *QuadTreeNode) InsertObject(obj *Object) {
	if q.level == q.maxLevel {
		if q.objects == nil {
			q.objects = make([]*Object, 0)
		}
		q.objects = append(q.objects, obj)
		return
	}

	if q.IsContain(q.x+q.width/2, q.y, q.width/2, q.height/2, obj) {
		if q.upRightNode == nil {
			q.upRightNode = NewQuadTreeNode(q.x+q.width/2, q.y, q.width/2, q.height/2, q.level+1, q.maxLevel, UPRIGHT, q)
		}
		q.upRightNode.InsertObject(obj)
	} else if q.IsContain(q.x, q.y, q.width/2, q.height/2, obj) {
		if q.upLeftNode == nil {
			q.upLeftNode = NewQuadTreeNode(q.x, q.y, q.width, q.height, q.level+1, q.maxLevel, UPLEFT, q)
		}
		q.upLeftNode.InsertObject(obj)
	} else if q.IsContain(q.x, q.y+q.width/2, q.width/2, q.height/2, obj) {
		if q.bottomLeftNode == nil {
			q.bottomLeftNode = NewQuadTreeNode(q.x, q.y+q.width/2, q.width/2, q.height/2, q.level+1, q.maxLevel, BOTTOMLEFT, q)
		}
		q.bottomLeftNode.InsertObject(obj)
	} else if q.IsContain(q.x+q.width/2, q.y+q.height/2, q.width/2, q.height/2, obj) {
		if q.bottomRightNode == nil {
			q.bottomRightNode = NewQuadTreeNode(q.x+q.width/2, q.y+q.height/2, q.width/2, q.height/2, q.level+1, q.maxLevel, BOTTOMRIGHT, q)
		}
		q.bottomRightNode.InsertObject(obj)
	}

}

//GetObjectsAt 查询对象,获得一片区域里的对象链表，此处只考虑完全包含的
func (q *QuadTreeNode) GetObjectsAt(px, py, w, h float32) []*Object {
	resObjects := make([]*Object, 0)
	if q.IsContainX(px, py, w, h, q) {
		resObjects = append(resObjects, q.objects...)
		if q.level == q.maxLevel {
			return resObjects
		}
	}
	if q.upRightNode != nil {
		upRightChild := q.upRightNode.GetObjectsAt(px, py, w, h)
		resObjects = append(resObjects, upRightChild...)
	}

	if q.upLeftNode != nil {
		upLeftChild := q.upLeftNode.GetObjectsAt(px, py, w, h)
		resObjects = append(resObjects, upLeftChild...)
	}

	if q.bottomLeftNode != nil {
		bottomLeftChild := q.bottomLeftNode.GetObjectsAt(px, py, w, h)
		resObjects = append(resObjects, bottomLeftChild...)
	}

	if q.bottomRightNode != nil {
		bootomRightChild := q.bottomRightNode.GetObjectsAt(px, py, w, h)
		resObjects = append(resObjects, bootomRightChild...)
	}

	return resObjects
}

//RemoveObjectsAt 删除对象，删除一片区域里的对象和节点，此处只考虑完全包含的
func (q *QuadTreeNode) RemoveObjectsAt(px, py, w, h float32) {
	//如果本层节点被包含则删除本层节点的对象
	//这个判断主要是对根节点起作用，其他子节点实际在上层都做了判断
	if q.IsContainX(px, py, w, h, q) {
		q.objects = make([]*Object, 0)
		if q.level == q.maxLevel {
			return
		}
	}
	if q.upRightNode != nil && q.IsContainX(px, py, w, h, q.upRightNode) {
		q.upRightNode.RemoveObjectsAt(px, py, w, h)
		q.upRightNode = nil
	}

	if q.upLeftNode != nil && q.IsContainX(px, py, w, h, q.upLeftNode) {
		q.upLeftNode.RemoveObjectsAt(px, py, w, h)
		q.upLeftNode = nil
	}

	if q.bottomLeftNode != nil && q.IsContainX(px, py, w, h, q.bottomLeftNode) {
		q.bottomLeftNode.RemoveObjectsAt(px, py, w, h)
		q.bottomLeftNode = nil
	}

	if q.bottomRightNode != nil && q.IsContainX(px, py, w, h, q.bottomRightNode) {
		q.bottomRightNode.RemoveObjectsAt(px, py, w, h)
		q.bottomRightNode = nil
	}
}
