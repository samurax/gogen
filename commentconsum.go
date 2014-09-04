package main
import  "strings" 

type CommentEater struct { 
eating   bool
}

func (c *CommentEater ) Eat (codeline  string )  ( string  ) { 
	switch c.eating {   case true :
                             if  poscom := strings.Index(codeline,"*/") ; poscom >=0 {
		                            c.eating = false
					    return  codeline[poscom+2:] }
		             return ""
		             case false :
		             if  poscom := strings.Index(codeline,"/*") ; poscom >= 0 {
				             c.eating = true
                                             current:= codeline [:poscom] 
                                             if  np:= strings.Index(codeline,"*/"); np >=0 && np > poscom  { 
						                                             c.eating = false
											     current = current + codeline[np+2:]
						                                             } 
					     return current }
			     if  poscom := strings.Index(codeline,"//") ; poscom >=0  {
			     return codeline [:poscom] }
			  }
	return codeline	
}
