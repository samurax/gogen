gogen
=====

 A very easy to use golang precompiler that generate generic types  based on templates. 
Golang with generics.


Overview:
              
	gogen  is a command line tool that can be considered a code generator for golang. 
It generates generic code based on templates and your types definition. The code generated 
don't use any external library. The program is really easy to use and give you more 
flexibility when defining types. 

How  to use it : 

	I will show you by an example how to write and commpile a generic implementation of a  
stack that can push and pop some graphics primitivies.
	  
file main.go:
          //---------------
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
                    Point        generic_stack     //Declaring a stack of Point    
		    Circle	 generic_stack     //Declaring a stack of Circle
	  }    
			    
          /*   When gogen find the pattern in the form : MemberDataName   generic_template 
	    it generates code based on a template named templatename.ge and save the 
	    resulting code in a file that will be compiled and linked with go build. Your
	    generic code become specific code and type checked by the compiler , and 
	    give your object the hability to invoke any method implemented in the templates
	    but with the apropiated type based on your declarations. */

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
                        
			//------
	  }	  

	By default templates are placed in your project directory under the directory gogen.d
and they have the extension .ge . The template can be defined inline in the code
We need to implement the generic function Push , Pop and NewStackXXX for any object
To write the template code , you only need to know two keywords  <data> and <once></once>
          
<data>  Will be replaced by your data type defined in your code.
<once>
	   //the definitions writen inside this blocks will be expanded by gogen only one time 
	   // for all types that use this template. here is where you place your types definitions
	   // and commons functions for the template 
</once>
            
file stack.ge:
//---------------
package main
	    
import "errors"   //will generate error on empty stack

<once>
	     type  generic_stack  bool       //In this example we don't need info inside the 
	                                     //generic_stack the type can be any.   
</once>  

             type   stackNode<data> struct { //Will be expanded as stackNodePoint for Point
                  link  *stackNode<data>     //A pointer to the linked object's node 
                  box   <data>	             //the data itself will be holded here by the node. 
	     }

             type Stack<data> struct {
                 link * stackNode<data>     //We start with an empty StackAnything
	     }

             func NewStack<data> () *Stack<data> { //return the stack 
                   return new ( Stack<data> )
	     }

             func (stk * Stack<data>) Push ( p <data>)  {  // pushing data 
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


	The best way to design a generic code is to write a specific implementation, 
test it well and create the template replacing the specific type with the <data> tag. 

    To precompile your code your project's files may looks as follow.


       main.go 

       gogen.d            // if you declare your generic code in a go file this isn't needed
              stack.ge    //

	  In your project directory execute:

	  gogen && go build

You can place the templates in the current directory too.

The resulting code will be writed to inkogen.go and is readable and gofmt-ed golang file

      You don't have to run gogen if you don't change the templates or the structure of 
your data definition that use generic.

      The structs that defines generics can be as many as you need and can have any name. For example if you
want to save your object using a generic template called saver , define your types as fallow:

      
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
               Point      generic_saver     //Declare generic saver  for Point
	           Circle     generic_dumper    //
          }


	The definition can be placed in any .go file in the current work directory.

For declaring generics without using a helper struct,  the following pattern can be used:
	      
	        membername    generic_membername

           type Shape  struct {
                Label    string 
		        Xo       float32
                Yo       float32
                Points    []Point    		
                Storage   generic_Storage     // Will declare a generic storage for Shape the 
		                              // type of the generic_data is implementation's dependant. 
	   }

For writing the template inline in the code, use a const keyword without a block as follow:


//---------------------------------------
package main
//-------------------------------------Template declaration for storage
const  generic_storage_template = 
`package main
import  "encoding/gob"
import  "os"
<once>
type generic_Storage  string  // gogen search generic_whatever 
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
}
`
//----------------------------------------------------End of template declaration.

type AnyThing struct {
     Description   string 
     GoodData      stuff 
}

func main () {
        o := new (AnyThing)
        o.Storage ="storage.file"
	o.Load()
}

//--------------------------------------------------------------------------------
All the code generate goes to a single file , you can use many generics for the same data
or for diferent data. The packages that every template use are imported as properly at the 
begining of the generated code.

   This is a very young project but is fully functional.   Let me know what you think 
about this implementation of generic in go , I love to hear from you. 

Thanks for reading.


