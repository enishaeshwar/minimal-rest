package random

import (
	"math/rand"
	"strconv"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) RandomNumbers(n int) []int {
	c := make(chan []int)

	var temp, res []int

	for i := 0; i < n; i++ {
		go s.getRand(n, c, "routine"+strconv.Itoa(n))
	}

	for i := 0; i < n; i++ {
		temp = <-c
		res = append(res, temp...)
	}

	return res
}

func (s *Service) getNoOfGoRoutines(n int) int {
	perRoutine := 10

	val := n / perRoutine
	if val == 0 {
		return 1
	}

	return val
}

func (s *Service) getRand(n int, c chan []int, name string) {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Intn(100)
	}

	c <- res
}
