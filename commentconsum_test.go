package main
//import  "strings"
import  "testing"
//import  "fmt"

func TAssert (  cond  bool ,  msg string  , t* testing.T) {
	if cond == false { t.Error("ERR:",msg) 
        }
	t.Log("TEST:",msg)
}


func TestEat ( t *testing.T )  {


code := `texto/*comment*/`

         cet := CommentEater {}
         codeline := cet.Eat(code)
	 TAssert ( codeline == "texto" , "Test1" , t )
         codeline =  cet.Eat("come/*------")
	 TAssert ( codeline == "come" , "Test2" , t )
         codeline =  cet.Eat("text invalid for output")
	 TAssert ( codeline =="" , "Test3",t)
         codeline = cet.Eat ("texto//test")
	 TAssert ( codeline =="" , "Test4",t)
         codeline = cet.Eat ("texto*/end")
	 TAssert ( codeline =="end" , "Test5",t)
         codeline = cet.Eat ("test//comment")
	 TAssert ( codeline =="test" , "Test6",t)
         codeline = cet.Eat ("test/*comment*/ pass")
	 TAssert ( codeline =="test pass" , "Test7",t)
	 codeline = cet.Eat ("text/*coments")
	 TAssert ( codeline =="text" , "Test8",t)
	 codeline = cet.Eat ("text/*coments")
	 TAssert ( codeline =="" , "Test9",t)
	 codeline = cet.Eat ("codeline")
	 TAssert ( codeline =="" , "Test10",t)
	 codeline = cet.Eat ("codeline*/pass")
	 TAssert ( codeline =="pass","Test11",t)
	  codeline = cet.Eat ("/**/test")
	 TAssert ( codeline =="test","Test12",t)


}



