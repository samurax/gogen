package main
import "testing"

func GoGen ( ) {
	stock1 := Stock {Name:"Stock1",file:"stock1.gob",
	                Prds:[]Product{ {"Laptop. Tochino 30400","0003404040"},
			                 {"TV felipo R455","00049590404"},
					 {"Multimedia Player","94995999"}, 
				       },
			 }
        stock1.Dump()
	stock1.Save()

        stock2 :=  Stock { file:"stock1.gob"  }
	stock2.Load()
	stock2.Dump()
        for  i,pr:= range stock1.Prds  {
              GAssert(pr == stock2.Prds[i],"Products are diferents") 
	}

	stock3 := Stock {Name:"Stock3",file:"stock1.gob",
	                Prds:[]Product{ {"Led Lamp. Jion","00303404040"},
			                 {"Raspberry","000445590404"},
				       },
			 }
	 stocks := new(StockGroup)
	 stocks.Stocks = []Stock{stock1,stock2,stock3}
	 stocks.file="teststocks.gob"
	 stocks.Save()

         newstocks := new (StockGroup)
	 newstocks.file = stocks.file
         newstocks.Load()
	 
	 newstocks.Dump()
}


func  TestGoGen ( t * testing.T ) {
GoGen ()
}
