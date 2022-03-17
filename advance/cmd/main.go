package main

func main() {
	go func() {
		sum := 0
		for i := 0; i < 100; i += 1 {
			sum += i
		}
	}()
	sum0 := 0
	for i0 := 0; i0 < 100; i0 += 2 {
		sum0 += i0
	}
}
