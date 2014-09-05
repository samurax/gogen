Gogen
=====

A very easy to use golang precompiler that generates generic types  based on templates. 

####Golang with generics.


> To get started here is a very simple example:  
> ```main.go:```



        package main
    
        type X struct {
         x int 
        }

        type Y struct {
         y  int
        }

        type generics struct {
        X      generic_varview
        Y      generic_varview
        }

        func main () {

        x:= X{100}
        y:= Y{200}
 
        x.View()
        y.View()
        }

        const  generic_varview_template=
        `package main;
        import "fmt"
        <once>
        type  generic_varview bool
        </once>

        func (v *<data>) View() {
        fmt.Printf("%#v\n",*v)
        }
        `  

>To compile the above example execute:

    user-project$    gogen && go build

> The precompiler search for patterns of the type: ```data generic_TemplateName``` in all types declaration of the current directory in all .go files. When it match a variable, it replace the tag ```<data>``` with the type of the data.

>To declare the template use: 

        const generic_TemplateName_template=
        ` 
        package name
        import  ...
        func ....
        `


> Templates are normal go files but with some specials tags.To write a template you need to know two keywords:

```<data>```  - Will be replaced for the type name

```<once><once>```  - Every thing inside this tags will be expanded only one time.

>Inside a block <once> is where you declare common types and common functions 
>for the template. every template have its owns blocks once

> The type of the keyword ```generic_template``` can be any valid golang type
> and should be declared inside the template. The declaration of generic can
> be done without using a helper type as in the above example. the follow declaration 
> is valid too for generating generics :

        type Circle struct {
        x        int
        y        int
        radius   int
        stack    generic_stack  
        }
                        
> In this case gogen generates using Circle as <data> :

        func (any *<data>) Draw () {
                 .......
        }
> Will be expanded for the type Circle as:

        func (any *Circle) Draw () {
                  ....
        }
> The declaration of this type declares stack inside the struct , this can 
be usefull if generic_stack is declared with usefull information. gogen never
modifies the user code and all generated code goes to a file named inkogen.go.
There isn't any library or dependency other than the code of the templates.
the code placed outside the block <once></once> will be expanded one time
for every data type that use generic. The templates have extension .ge and are 
placed in a directory named gogen.d in the project directory an can be reused 
easily. If you declare the template inline in the code as in the example, you
don't need this directory.  Don't forget to use the command gogen every time
you modify a template file or the code inside the const declaration of the template.

      user-project$:  gogen && go build
    

> The project is in very early stage  but is fully functional. I started as 
> experiment to see how generic fit into golang. IMHO fit good.  
> Let me know what you think about this implementation of generic in go. 

> I love to hear from you. 

> Thanks for reading.

