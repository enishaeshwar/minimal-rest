package service

import (
	"math/rand"
)

func RandomNumbers(n int) []int {
	c := make(chan []int)

	var res []int

	go getRand(n, c, "routine1")
	go getRand(n, c, "routine2")
	go getRand(n, c, "routine3")
	go getRand(n, c, "routine4")
	go getRand(n, c, "routine5")

	t1, t2, t3, t4, t5 := <-c, <-c, <-c, <-c, <-c

	res = append(
		append(
			append(
				append(
					append(res, t1...),
					t2...),
				t3...),
			t4...),
		t5...)

	return res
}

func getRand(n int, c chan []int, name string) {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = rand.Intn(100)
	}

	c <- r
}
