package main

import "fmt"
import "strings"

type generic_dumper bool

func replace_chars_dumper(str string) string {
	str = strings.Replace(str, ":", "=", -1)
	str = strings.Replace(str, "{", "\n....... ", -1)
	str = strings.Replace(str, ",", "\n.......", -1)
	str = strings.Replace(str, "}", "\n", -1)
	return str
}
func (d *Data) Dump() {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("_______________Data__")
	fmt.Println(replace_chars_dumper(str))
}
func (d *Person) Dump() {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("_______________Person__")
	fmt.Println(replace_chars_dumper(str))
}
func (d *SPerson) Dump() {
	str := fmt.Sprintf("%#v", *d)
	fmt.Println("_______________SPerson__")
	fmt.Println(replace_chars_dumper(str))
}
