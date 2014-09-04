package main
import  "errors"
import  "bufio"
import  "strings"

func   saveInlineCode ( constname string , code string ,temps *TempFiles ) error {
       gerr :=errors.New("Inline conditions fails")
       pos:=strings.Index(constname,"=")
       if pos > 0 { constname =   constname[:pos] }
       constname = strings.TrimSpace(constname)
       idens:=strings.Split(constname,"_")
       switch {  case len(idens) != 3 ,idens[0] != "generic" , idens[2] != "template":
                  return gerr }
       codeiden := idens[1]
       file,err:= temps.OpenFile (codeiden)
       if err != nil {  return err  }
       w:= bufio.NewWriter(file)
       w.WriteString(code)
       w.Flush()
       return nil
}

func    ReadInlineCode ( constname string ,  initcode string , scan * bufio.Scanner ,temps *TempFiles) {
	poseq := strings.Index(initcode,"=")
	if poseq < 0 { return  }
        var code string
	start := false
	pos   :=  strings.Index(initcode,"`")
	if pos > 0 { code = initcode[pos+1:]
	             start = true 
	     } else { code = initcode[poseq+1:]
                    }
	code = initcode[pos+1:]
	for scan.Scan() {
                codeline :=  scan.Text()
                pos = strings.Index(codeline,"`")
		switch {  case  start == false && pos >= 0:
		                start = true
			        code = codeline[pos+1:]
		          case  start == false :
                          case pos < 0 :
			  code += codeline +"\n"
			  default:
			  code += codeline[:pos] +"\n"
			  saveInlineCode(constname,code,temps)
			  return
		       }
	}
}
