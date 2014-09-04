package main
import  "fmt"
import  "os"
import  "bufio"
import  "testing"

const generic_flaco_template=
`texto
decontrol
withcode`


const generic_dumper_template =`
package main
import "fmt"
dkjkdjkd

func Test(rso *dt)
`
const generic_going_template=`
func Test<data>(d<data>) {
costa()
}`

/*
func tokenizeChar(ch string, token string) []string {
	ks := strings.Split(token, ch)
	nks := make([]string, 0)
	for _, t := range ks {
		s := strings.TrimSpace(t)
		if s == "" {
			continue
		}
		nks = append(nks, s)
	}
	return nks
}

func tokenizeSpec(okeys []string) []string {
	keys := make([]string, 0)
	for _, k := range okeys {
		ke := strings.TrimSpace(k)
		if ke == "" {
			continue
		}
		comma := tokenizeChar(ch_comma, ke)
		for _, kp := range comma {
			keys = append(keys, kp)
		}
	}
	return keys
}

func tokenizeCodeline(textline string) []string {
	ks := strings.Split(textline, " ")
	keys := make([]string, 0)
	for _, k := range ks {
		cl := strings.TrimSpace(k)
		if cl == "" {
			continue
		}
		keys = append(keys, cl)
	}
	return tokenizeSpec(keys)
}


func isGenericName(name string) bool {
	generictype := "generic"
	l := len(name)
	if l > 8 && name[0:8] == generictype+"_" {
		return true
	}
	return false
}

*/
func  TestReadInlineCode (t *testing.T) {
	temps := TempFiles {}
	temps.Init()

	filename:="inlinecode_test.go"
	f,err:= os.Open(filename)
	if err  != nil {  t.Error("Fail opening file "+filename)  }
	scan:= bufio.NewScanner(f)
	for scan.Scan() {
                  codeline:=scan.Text()
		  keys := tokenizeCodeline (codeline)
                  l:=len(keys)
		  switch   {    case  l==0:  
		                   continue 
		                   case  keys[0] == "const":
                                   ReadInlineCode ( keys[1] , codeline ,  scan , &temps )

		            }
	 }

	 for  n,v := range temps.tfiles {
              fmt.Println("-----------",n)
	          f,err:=os.Open(v.Name())
		  if err != nil { t.Error("Reopeing file for scanning fails")  }
	          scan:= bufio.NewScanner(f)
	          for scan.Scan() {
	              txt:=scan.Text()
		      fmt.Println("-----",txt)
		  }
	 }

         temps.Remove() 

}

