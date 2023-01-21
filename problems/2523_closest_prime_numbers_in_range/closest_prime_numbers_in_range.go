package main

import "math"

func closestPrimes(left int, right int) []int {
	primes := make([]int, right+1)
	for i := 0; i < right+1; i++ {
		primes[i] = i
	}
	primes[1] = 0
	//remove x*x + n*x
	for i := 0; i <= int(math.Sqrt(float64(right))); i++ {
		if primes[i] != 0 {
			j := i * i
			for j < right+1 {
				primes[j] = 0
				j += i
			}
		}
	}
	var ps []int
	for i := 0; i < right+1; i++ {
		if primes[i] != 0 && primes[i] >= left {
			ps = append(ps, primes[i])
		}
	}
	min := math.MaxInt
	res1, res2 := 0, 0
	for i := 1; i < len(ps); i++ {
		if ps[i]-ps[i-1] < min {
			res1, res2 = ps[i-1], ps[i]
			min = res2 - res1
		}
	}
	if res1 == res2 {
		return []int{-1, -1}
	}
	return []int{res1, res2}
}
