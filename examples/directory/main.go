package  main

import  "fmt"
import  "sync"
import  "flag"

type  Godata   struct {
      Count    int
}


type Search_go struct {
     count       int
     Search_go   generic_dirscan
     Data        Godata 
}

type  Htmldata struct {
      Html      int 
      Css       int
      Js        int
 }

 type  Dumpers      struct {
        Godata      generic_dumper
	Htmldata    generic_dumper
 }

type Search_html struct {
        Data         Htmldata
	Search_html  generic_dirscan
}

func (sg * Search_go ) DirInfo (prefix string)   {
	msg:=""
	  switch  sg.Search_go.Ext {
	             case ".go":
		          sg.Data.Count++
	                  msg="_________GO"
		     default:
		    return
	 }
	 fmt.Println( "go",prefix , sg.Search_go.Base,msg )
}

func (sg * Search_html ) DirInfo (prefix string)   {
	msg :="________"
	 switch sg.Search_html.Ext {
                  case ".html":
			sg.Data.Html ++
			msg += "HTML"
		  case ".js":
			sg.Data.Js  ++
			msg += "JS"
		  case  ".css":	
		        sg.Data.Css ++
			msg +="CSS"
		  default:
			 return 
	   }

	 fmt.Println("html",prefix,sg.Search_html.Base,msg) 
}

func main () {
              flag.Parse()  
	      args := flag.Args()
	      startdir := "../../"
	      if (len(args)>0)  {  startdir = args[0] } 

              gofind := new (Search_go)
	      htfind := new (Search_html)
              var  wg  sync.WaitGroup
              wg.Add(2)

	      go func () {  htfind.Scan (startdir,"") 
                         wg.Done()
                    }  ()

	      go func () { gofind.Scan (startdir,"")
	                 wg.Done() }  ()
	      wg.Wait()

	      gofind.Data.Dump("GO")
              htfind.Data.Dump("HTML")
}







