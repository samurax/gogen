package main

import "errors"

type generic_stack bool
type stackNodePoint struct {
	link *stackNodePoint
	box  Point
}
type StackPoint struct {
	link *stackNodePoint
}

func NewStackPoint() *StackPoint {
	return new(StackPoint)
}
func (stk *StackPoint) Push(p Point) {
	nn := new(stackNodePoint)
	nn.box = p
	nn.link = stk.link
	stk.link = nn
}
func (stk *StackPoint) Pop() (*Point, error) {
	if stk.link == nil {
		return nil, errors.New("No more")
	}
	ret := &stk.link.box
	stk.link = stk.link.link
	return ret, nil
}

type stackNodeCircle struct {
	link *stackNodeCircle
	box  Circle
}
type StackCircle struct {
	link *stackNodeCircle
}

func NewStackCircle() *StackCircle {
	return new(StackCircle)
}
func (stk *StackCircle) Push(p Circle) {
	nn := new(stackNodeCircle)
	nn.box = p
	nn.link = stk.link
	stk.link = nn
}
func (stk *StackCircle) Pop() (*Circle, error) {
	if stk.link == nil {
		return nil, errors.New("No more")
	}
	ret := &stk.link.box
	stk.link = stk.link.link
	return ret, nil
}
