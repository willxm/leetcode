package problems

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	var ret int
	ch := GenerateNatural()
	for i := 0; i < n/2; i++ {
		prime := <-ch
		ch = PrimeFilter(ch, prime)
		if prime < n {
			ret++
		} else {
			return ret
		}
	}
	return ret
}
