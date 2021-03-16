package main

type AvlNode struct {
	data int
	height int
	left *AvlNode
	right *AvlNode
}

type AvlTree struct {
	root *AvlNode
}

func (this *AvlTree) Insert(t **AvlNode, x int){
	if (*t) == nil {
		(*t) = new(AvlNode)
		(*t).data = x
	}else if x < (*t).data {
		this.Insert(&((*t).left), x)
		if this.GetHeight((*t).left) - this.GetHeight((*t).right) > 1{
			if x < (*t).left.data {
				*t = this.LL(*t)
			}else {
				*t = this.RR(*t)
			}
		}
	}else if x > (*t).data {
		this.Insert(&((*t).right), x)
		if this.GetHeight((*t).right) - this.GetHeight((*t).left) > 1{
			if x > (*t).right.data {
				*t = this.RR(*t)
			}else{
				*t = this.LL(*t)
			}
		}
	}else{
		;
	}

	lh := this.GetHeight((*t).left)
	rh := this.GetHeight((*t).right)
	if lh > rh {
		(*t).height = lh + 1
	}else {
		(*t).height = rh + 1
	}
}
func (this *AvlTree) Delete(t *AvlNode, x int) bool{
	return false
}

func (this* AvlTree)Contains(t *AvlNode)bool {
	return false
}

func (this* AvlTree) InorderTraversal(t *AvlNode){
}

func (this* AvlTree) PreorderTraversal(t *AvlNode){

}

func (this *AvlTree) FindMax(t *AvlNode) *AvlNode{
	if t == nil || t.right==nil{
		return t
	}
	return this.FindMax(t.right)
}
func (this *AvlTree) FindMin(t *AvlNode) *AvlNode{
	if t == nil || t.left == nil {
		return t
	}
	return this.FindMin(t.left)
}

func (this *AvlTree) GetHeight(t *AvlNode) int{
	if t == nil{
		return -1
	}else{
		return t.height
	}
}
func (this *AvlTree) LL(t *AvlNode) *AvlNode{
	q := t.left
	t.left = q.right
	q.right = t

	rh := this.GetHeight(t.right)
	lh := this.GetHeight(t.left)
	if rh > lh {
		t.height = rh + 1
	}else{
		t.height = lh + 1
	}
	return q
}

func (this *AvlTree) RR(t *AvlNode) *AvlNode{
	q := t.right
	t.right = q.left
	q.left = t

	lh := this.GetHeight(t.left)
	rh := this.GetHeight(t.right)
	if lh > rh {
		t.height = lh + 1
	}else{
		t.height = rh + 1
	}
	return q
}
/*
	t的左孩子的右子树上插入,
	先对t的左孩子节点进行rr旋转
	再对t进行ll旋转
 */
func (this *AvlTree) LR(t *AvlNode) *AvlNode{
	this.RR(t.left)
	return this.LL(t)
}

func (this *AvlTree) RL(t *AvlNode) *AvlNode{
	this.LL(t.right)
	return this.RR(t)
}

