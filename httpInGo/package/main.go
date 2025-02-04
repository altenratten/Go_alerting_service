package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Counter(count int, t time.Time) string {
	date := t.Format("02.01.2006")
	return fmt.Sprintf("%d %s", count, date)
}

func ParseCounter(input string) (int, time.Time, error) {
	// допишите код функции
	// ...
	vals := strings.Split(input, " ")
	counter, err := strconv.Atoi(vals[0])
	t, err := time.Parse("02.01.2006", vals[1])
	return counter, t, err
}

func main() {
	now := time.Now()
	for count := 1; count < 4; count++ {
		s := Counter(count, now)
		counter, t, err := ParseCounter(s)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println(counter, t.Format("02.01.2006"))
	}
}
