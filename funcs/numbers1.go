package funcs

func IdNum1 (num int)(string){
	if num < 0 && num % 2 == 0 && num < 10{
		return ("Чётное однозначное отрицательное число")
	}else if num > 0 && num % 2 == 0 && num < 10{
		return ("Чётное однозначное положительное число")
	}else if num < 0 && num % 2 == 0 && num < 100{
		return ("Чётное двухзначное отрицательное число")
	}else if num > 0 && num % 2 == 0 && num < 100{
		return ("Чётное двухзначное положительное число")
	}else if num < 0 && num % 2 == 0 && num < 1000{
		return ("Чётное трёхзначное отрицательное число")
	}else if num > 0 && num % 2 == 0 && num < 1000{
		return ("Чётное трёхзначное положительное число")
	}else if num < 0 && num % 2 != 0 && num < 10{
		return ("Нечётное однозначное отрицательное число")
	}else if num > 0 && num % 2 != 0 && num < 10{
		return ("Нечётное однозначное положительное число")
	}else if num < 0 && num % 2 != 0 && num < 100{
		return ("Нечётное двухзначное отрицательное число")
	}else if num > 0 && num % 2 != 0 && num < 100{
		return ("Нечётное двухзначное положительное число")
	}else if num < 0 && num % 2 != 0 && num < 1000{
		return ("Нечётное трёхзначное отрицательное число")
	}else if num > 0 && num % 2 != 0 && num < 1000{
		return ("Нечётное трёхзначное положительное число")
	}else {
		return ("Нуливое число")
	}
}