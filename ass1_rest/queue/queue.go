package queue

import "fmt"

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(item int, mode bool) {
	if mode {
		if len(*q) < 1000000 {
			*q = append(*q, item)
		} //3
	} else {
		if len(*q) < 1000000 {
			*q = append([]int{item}, *q...)
		}

	}
	fmt.Println("Input: ", item)
	fmt.Println(*q)
}

func (q *Queue) Dequeue(mode bool) int {
	if q.IsEmpty() {
		panic("-")
	} else {
		if mode {
			item := (*q)[0]
			*q = (*q)[1:]
			fmt.Println(*q)
			return item
		} else {
			item := (*q)[len(*q)-1]
			*q = (*q)[:len(*q)-1]
			fmt.Println(*q)
			return item
		}
	}

}
