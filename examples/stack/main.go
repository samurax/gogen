package main 

type Point struct {
x      int
y      int
label  string
}

type  Circle  struct {
p      Point
r      int
label  string
}

type      GenericsStack struct {
 Circle   generic_stack
 Point    generic_stack
}
type      Dumpers  struct {
 Circle   generic_dumper //dumper for Circle 
 Point    generic_dumper //dumper for Point
}


func main () {
	st := NewStackPoint()
	p:=Point{59,56,"test"}
	st.Push(p)

	stkc := NewStackCircle()
	c := Circle{p,20,"circle"}
        stkc.Push(c)

        r , _ := st.Pop()
	r.Dump()
	cir,_ := stkc.Pop()
	cir.Dump()
}
