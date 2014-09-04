package main

import  "bufio" 
import  "testing"
import   "os"

func TestInit ( t * testing.T) {  

        tt:=TempFiles {}
	tt.Init()
	file,err := tt.OpenFile("tester")
        if err != nil  { t.Error("FAIL TO CREATE TEMP FILE")  }
        w:= bufio.NewWriter(file)
	code:=
`1234567890
abcdefghi`

        w.WriteString(code)
	w.Flush()

       f,err := os.Open(file.Name()) 
       scan := bufio.NewScanner(f)
   
       scan.Scan()
       str1 := scan.Text()
       scan.Scan()
       str2 := scan.Text()
       tt.Remove()
       _,rerr:=os.Stat(file.Name())

         switch  { case  str1 != "1234567890",str2 != "abcdefghi":
                         t.Error("Strings read write file")
		   case	 rerr == nil:
		         t.Error("The File is not removed")
                 }

}




