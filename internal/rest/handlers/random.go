package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func RandomHandler(w http.ResponseWriter, r *http.Request) {
	res := randomNumbers(10)
	fmt.Println(res)

	resS := strings.Trim(strings.Join(strings.Split(fmt.Sprint(res), " "), ","), "[]")
	fmt.Println(resS)

	w.Write([]byte(resS))
}

func randomNumbers(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		rand.In
		r = append(r, rand.Int())
	}

	return r
}
