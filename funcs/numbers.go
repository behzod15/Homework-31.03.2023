package funcs

func IdNum (num int)(string){
	if num < 0 && num % 2 == 0{
		return ("Чётное отрицательное число")
	}else if num > 0 && num % 2 == 0 {
		return ("Чётное положительное число")
	}else if num > 0 && num % 2 != 0 {
		return ("Нечётное положительное число")
	}else if num < 0 && num % 2 != 0 {
		return ("Нечётное отрицательное число")
	}else {
		return ("Нуливое число")
	}
}