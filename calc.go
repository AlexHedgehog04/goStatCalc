package main

import (
    "flag" //работа с аргументами CLI
    "fmt" //StdOut
    "strconv" //конвертация
    "strings" //операции со строками
	"github.com/montanaflynn/stats" //статистические функции
)


//Принимаем -numbers string, конвертим в []float64 циклом через make()+for, считаем показатели  
func main() {
    numbers := flag.String("numbers", "", "числа через запятую (ex: 1.5,2.3,4.1)")
    flag.Parse()
	
    //Если нет ввода -numbers=...
    if *numbers == "" {
        fmt.Println("Используй: go run . -numbers=1.5,2.3,4.1")
        return
    }
    
    // Парсим строку в []float64
    strSlice := strings.Split(*numbers, ",")
    nums := make([]float64, 0, len(strSlice))
    
    for _, s := range strSlice {
        // обработка ввода не числа
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64) //Парсим []string -> []float64
        if err != nil {
            fmt.Printf("Ошибка: %s не число\n", s)
            return
        }
        nums = append(nums, f)

    }
    
	
    
    mean, _ := stats.Mean(nums) //blank identifier - беру среднее, ошибку игнорирую
	median, _ := stats.Median(nums)
	min, _ := stats.Min(nums)
	max, _ := stats.Max(nums)

	fmt.Printf("Вход: %v\n", nums)
    fmt.Printf("Среднее: %.2f\n", mean)
    fmt.Printf("Медиана: %.2f\n", median)
    fmt.Printf("Мин: %.2f\n", min)
    fmt.Printf("Макс: %.2f\n", max)
}
