package funcs


func Output (year int) (string){
	if  year%100 == 0 &&  year%400 != 0 {
		return ("В году насчитывается 365 дней")   
	}else {
		return ("В году насчитывается 366 дней")
	}
	
}