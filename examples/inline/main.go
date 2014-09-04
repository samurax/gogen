package main
//-------------------------const generic_templatename_template 
//inline pattern
//--------  compile with gogen && go build
const generic_dumper_template =
`
package main;
import "fmt";
import "strings";
<once>
type generic_dumper bool
func  replace_chars_dumper (str string ) string {
str =  strings.Replace(str,":","=",-1)
str =  strings.Replace(str,"{","\n....... ",-1)
str =  strings.Replace(str,",","\n.......",-1)
str =  strings.Replace(str,"}","\n",-1)
return  str
}
</once>
func (d *<data>) Dump () {
str := fmt.Sprintf("%#v",*d)
fmt.Println("_______________<data>__")
fmt.Println(replace_chars_dumper(str))
}
`
type  Data  struct  {
      label  string
      v      int
}

type  Person  struct {
      name     string
      age      int
      phone    string
}

type  Definitions struct {
      Data    generic_dumper
      Person  generic_dumper
}

type  SPerson struct {
      name  string
      age   int
      phone string
      dumper generic_dumper // template generic_template ,pattern apply
                            // to generic for the current type
			    // In this example declare a dumper generic for Sperson
}

func main () {
       d := Data    { "Tester data", 1000   }
       d.Dump()
       p := Person  { "Frank Martin",34,"8485993302020"}
       p.Dump()

       sp := SPerson { name:"Roberto",age:40,phone:"340400500555" }
       sp.Dump()
}

