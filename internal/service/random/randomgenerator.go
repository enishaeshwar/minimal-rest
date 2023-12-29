package random

import (
	"math/rand"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) RandomNumbers(n int) []int {
	c := make(chan []int)

	var res []int

	go s.getRand(n, c, "routine1")
	go s.getRand(n, c, "routine2")
	go s.getRand(n, c, "routine3")
	go s.getRand(n, c, "routine4")
	go s.getRand(n, c, "routine5")

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

func (s *Service) getRand(n int, c chan []int, name string) {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Intn(100)
	}

	c <- res
}
