
Gogen
=====

A very easy to use golang precompiler that generate generic types  based on templates. Golang with generics.

###How it works.
              
  gogen  is a command line tool.
don't use any external library. The program is really easy to use and give you more flexibility when defining your types. 

How  to use : 

I will show you by an example how to write and commpile a generic implementation of a  generic stack that can push and pop some graphics primitivies.
	  
file main.go:


        package main

        type Point   struct {
        x       int   
        y       in 
        label   string
        } 

        type Circle struct {
               p       Point 
               radius  int   
        }        

        type AnyNameIsValid struct  {
                    Point  generic_stack     //Declaring a stack of Point    
                    Circle generic_stack     //Declaring a stack of Circle
	  }    
	  

>When gogen find the pattern in the form : 
        
            DataName  generic_template
            
>It generates code based on the template named templatename.ge, do all sustitution as needed and save the resulting code in a file that will be compiled and linked with go build. Your generic code become specific code and type checked by the compiler whe you execute go build. 

          func main () {
                spoints := NewStackPoint() 
					 
                p1:=Point{10,20,"test1"}
                p2:=Point{20,10,"test2"}
    
                spoints.Push(p1)
                spoints.Push(p2)

                scircles := NewStackCircle()
                c1:= Circle{p1,30}
                c2:= Circle{p2,40}
			
                scircles.Push(c1)
                scircles.Push(c2)	
                         
                cx , err := scircles.Pop()	
                //------  cx = c2

                cx.p , err := spoints.Pop()
                        
	  }	  

>By default templates are placed in your project directory under the directory gogen.d and they have the extension .ge  The template can be defined inline in the code. To continue with the example We need to implement the generic function Push, Pop and NewStack<data>. 
To write the template code, you only need to know two keywords:
        
         <data>   and       <once>      </once>
          
***<data>***      Will be replaced by your data type defined in your code.

***<once>***
The definitions written inside this blocks will be expanded by gogen only one time for all types that use this template. here is where you place your types definitions and commons functions for the template. 
***</once>***
            
file stack.ge:

    package main	    
    import "errors"   //will generate error on empty stack

    <once>
    type  generic_stack  bool   
    </once>  

    type   stackNode<data> struct { 
    stackNodePoint for Point
    link  *stackNode<data>     
    box   <data>	       
    }

    type Stack<data> struct {
    link * stackNode<data>    
    }

    func NewStack<data> () *Stack<data> { 
    return new ( Stack<data> )
    }

    func (stk * Stack<data>) Push ( p <data>)  { 
        nn := new (stackNode<data>)
        nn.box = p
        nn.link = stk.link
        stk.link =  nn
    }

    func (stk * Stack<data>) Push (p <data>)  { 
        nn := new (stackNode<data>) 
        nn.box = p
        nn.link = stk.link
        stk.link =  nn
    }

    func (stk * Stack<data>) Pop () (* <data> , error ) {
    if  stk.link == nil {  return  nil , errors.New("No more") }
    ret :=  &stk.link.box
    stk.link = stk.link.link
    return ret , nil
    }


	The best way to design a generic code is to write a specific implementation, test it well and create the template replacing the example type with the <data> tag. 

To precompile your code your project's files may looks as follow.

       main.go 

       gogen.d    
              stack.ge

	
>In your project directory execute:

    gogen && go build

You can place the templates in the current directory too. The resulting code 
will be writed to inkogen.go by default and is readable and gofmt-ed golang file

You don't have to run gogen if you don't change the templates or the structure of your data definition that use generic.

 The structs that defines generics can be as many as you need and can have any name. For example if you want to save your object using a generic template called saver , define your types as fallow:

      
            type Point   struct {
                x       int   
                y       in 
                label   string
                } 

            type Circle struct {
               p       Point 
               radius  int   
            }        

            type generic_stacks struct {
               Point      generic_stack  
               Circle     generic_stack
            }        

            type generic_others struct {
            Point      generic_saver    
            Circle     generic_dumper 
            }


	The definition can be placed in any .go file in the current work directory.
>You can declare generic data without using a helper struct,the following pattern can be used:
	      
***membername    generic_membername***

           type Shape  struct {
                Label     string 
                Xo        float32
                Yo        float32
                Points    []Point    		
                Storage   generic_Storage   
	   }

   Storage generic_Storage declare an Storage for Shape 

For writing the template inline in the code, use a const keyword without a block as follow:


        package main

        const  generic_storage_template = 
        `package main
        import  "encoding/gob"
        import  "os"
        <once>
        type generic_Storage  string   
        </once>

        func  (pg *<data>) Store() error {
        f , err := os.OpenFile (pg.Storage,os.O_WRONLY | os.O_TRUNC | os.O_CREATE,644)
        defer f.Close()
        if err == nil  {  encoder := gob.NewEncoder(f)
        err = encoder.Encode(*pg) 
        }
        return err 
        }

        func  (pg *<data>) Load() error  {
        f , err := os.Open(pg.Storage) 
        if err == nil {  defer f.Close()
        decoder := gob.NewDecoder(f) 
        err = decoder.Decode(pg)
        }
        return err
        }`

Declarations outside the template:

        type AnyThing struct {
         Data   string
         Other [] strings
         storage   generic_storage
        }


        func main () {
        o := new (AnyThing)
        o.Storage ="storage.file"
    	  o.Load()
        }


>All the code generate goes to a single file , you can use more than one generics for the same data. The packages that every template use are imported as properly at the begining of the generated code.

>The project is in very early stage  but is fully functional. 
>Typical generic implementations can be designed , I only create a stack but only >as example. 

   Cool complexity only emerge from cool simplicity. And golang is cool. 
   
 Let me know what you think about this implementation of generic in go. 

 I love to hear from you. 

Thanks for reading.

