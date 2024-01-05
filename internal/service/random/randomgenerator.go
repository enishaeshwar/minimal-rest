package random

import (
	"math"
	"math/rand"
	"strconv"
)

const (
	PerRoutineCount = 10
)

type Service struct{}

type goRoutineConfig struct {
	count     int
	remainder int
}

type Result struct {
	Values []int `json:"values"`
	Sum    int   `json:"sum"`
}

func New() *Service {
	return &Service{}
}

func (s *Service) RandomNumbers(n int) Result {
	var temp, res []int
	c := make(chan []int)

	cfg := s.goRoutineConfig(n)

	// start go routines
	for i := 0; i < cfg.count-1; i++ {
		go s.getRand(PerRoutineCount, c, "routine_"+strconv.Itoa(i))
	}
	// last go routine with remainder
	go s.getRand(cfg.remainder, c, "routine_"+strconv.Itoa(cfg.count-1))

	for i := 0; i < cfg.count; i++ {
		temp = <-c
		res = append(res, temp...)
	}

	return Result{
		Values: res,
		Sum:    s.sum(res),
	}
}

func (s *Service) goRoutineConfig(n int) goRoutineConfig {
	var c, r int

	if c = int(math.Ceil(float64(n) / PerRoutineCount)); c == 0 {
		c = 1
	}

	r = n % PerRoutineCount

	return goRoutineConfig{
		count:     c,
		remainder: r,
	}
}

func (s *Service) getRand(n int, c chan []int, name string) {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = rand.Intn(100)
	}

	c <- res
}

func (s *Service) sum(arr []int) int {
	var sum int

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}
