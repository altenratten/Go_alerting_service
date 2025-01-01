package main

import (
	"fmt"
	"strings"
)

// Queue описывает очередь.
type Queue struct {
	first *QueueItem // указатель на первый элемент очереди
	last  *QueueItem // указатель на последний элемент очереди
}

// QueueItem описывает элемент очереди.
type QueueItem struct {
	value string     // данные
	next  *QueueItem // указатель на следующий элемент
}

// Pop удаляет первый элемент из очереди и возвращает хранимую там строку.
// вставьте здесь метод Pop() для типа *Queue
// ...
func (q *Queue) Pop() (string, bool) {
	if q.first == nil {
		return "", false
	}
	item := q.first
	q.first = item.next
	item.next = nil

	if q.first == nil {
		q.last = nil
	}

	return item.value, true
}

// Push добавляет в конец очереди элемент с указанной строкой.
func (queue *Queue) Push(value string) {
	item := &QueueItem{value: value}
	if queue.first == nil { // нет элементов
		// очередь пустая, поэтому добавляемый элемент
		// станет и первым и последним
		queue.first = item
		queue.last = item
		return
	}
	queue.last.next = item // текущий последний элемент должен указывать
	// на добавленный элемент
	queue.last = item // item становится последним элементом
}

func main() {
	list := []string{"На", "золотом", "крыльце", "сидели:", "царь,", "царевич,",
		"король,", "королевич."}
	queue := &Queue{}

	for _, v := range list {
		queue.Push(v)
	}
	list = list[:0]
	for {
		v, ok := queue.Pop()
		if !ok {
			break
		}
		list = append(list, v)
	}
	fmt.Print(strings.Join(list, " "))
}
