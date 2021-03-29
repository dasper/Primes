package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

type PrimeInterface interface {
	PrintResults(bool, time.Duration, int)
	Run()
	Set(int)
}

type primeBase struct {
	size int
	sqrt int
}

type primeBool struct {
	primeBase
	isPrime []bool
}

type primeSieve struct {
	primeBase
	bits []uint8
}

var myDict = map[int]int{
	10:        4,
	100:       25,
	1000:      168,
	10000:     1229,
	100000:    9592,
	1000000:   78498,
	10000000:  664579,
	100000000: 5761455,
}

func main() {
	limit := flag.Int("l", 1000000, "limit")
	useBool := flag.Bool("b", true, "use bool instead of bits")
	flag.Parse()

	var p PrimeInterface
	passes := 0
	start := time.Now()
	if *useBool {
		p = &primeBool{}
	} else {
		p = &primeSieve{}
	}

	for time.Since(start).Seconds() < 10 {
		p.Set(*limit)
		p.Run()
		passes++
	}
	tD := time.Since(start)
	p.PrintResults(false, tD, passes)
}

func (p *primeBool) Set(index int) {
	p.size = index
	p.isPrime = nil
	p.isPrime = make([]bool, index+1)
	p.sqrt = int(math.Sqrt(float64(index)))
	p.isPrime[2] = true
	for i := 3; i <= index; i = i + 2 {
		if i%2 != 0 {
			p.isPrime[i] = true
		}
	}

}

func (p *primeBool) Run() {
	factor := 3
	for factor <= p.sqrt {
		for num := factor; num <= p.size; num = num + 2 {
			if p.isPrime[num] {
				factor = num
				break
			}
		}
		p.clear(factor)
		factor = factor + 2
	}
}

func (p *primeBool) clear(factor int) {
	for num := factor * 3; num <= p.size; num += factor * 2 {
		p.isPrime[num] = false
	}
}

func (p *primeBool) PrintResults(showResults bool, duration time.Duration, passes int) {
	primes := make([]int, 0)
	for num := 2; num <= p.size; num++ {
		if p.isPrime[num] {
			primes = append(primes, num)
		}
	}
	if showResults {
		fmt.Println(primes)
	}
	var avg time.Duration = duration / time.Duration(passes)
	fmt.Printf("Passes: %d, Time: %v, Avg: %v, Limit: %d, Count: %d, Valid: %t\n",
		passes,
		duration,
		avg,
		p.size,
		len(primes),
		p.validateResults())
}

func (p *primeBool) validateResults() bool {
	value, ok := myDict[p.size]
	if !ok {
		fmt.Println("Warning: validation quantity not in map")
		return false
	}
	return value == p.countPrimes()
}

func (p *primeBool) countPrimes() int {
	count := 0
	for i := 0; i < p.size; i++ {
		if p.isPrime[i] {
			count++
		}
	}
	return count
}

func (p *primeSieve) getBit(index int) bool {
	if index%2 == 0 {
		return false
	}
	index = index / 2
	return ((p.bits[index/8]) & (1 << (index % 8))) != 0
}

func (p *primeSieve) clearBit(index int) {
	if index%2 == 0 {
		fmt.Println("You're clearing even bits, which is sub-optimal.")
		return
	}
	index = index / 2
	p.bits[index/8] &^= (1 << (index % 8))
}

func (p *primeSieve) validateResults() bool {
	if _, ok := myDict[p.size]; !ok {
		return false
	}
	return myDict[p.size] == p.countPrimes()
}

func (p *primeSieve) Run() {
	factor := 3

	for factor < p.sqrt {
		for num := factor; num < p.size; num += 2 {
			if p.getBit(num) {
				factor = num
				break
			}
		}

		for num := factor * 3; num < int(p.size); num += factor * 2 {
			p.clearBit(num)
		}

		factor += 2
	}
}

func (p *primeSieve) countPrimes() int {
	count := 0
	for i := 0; i < p.size; i += 1 {
		if p.getBit(i) {
			count += 1
		}
	}
	return count
}

func (p *primeSieve) Set(size int) {
	bits := make([]uint8, size/8+1)
	for i := range bits {
		bits[i] = 0xFF
	}
	p.bits = bits
	p.size = size
	p.sqrt = int(math.Sqrt(float64(p.size)))
}

func (p *primeSieve) PrintResults(showResults bool, duration time.Duration, passes int) {
	if showResults {
		fmt.Printf("2, ")
	}

	count := 1
	for num := 3; num <= p.size; num += 1 {
		if p.getBit(num) {
			if showResults {
				fmt.Printf("%d, ", num)
			}
			count++
		}
	}

	if showResults {
		fmt.Println()
	}

	fmt.Printf("Passes: %d, Time: %.6f, Avg: %.6f, Limit: %d, Count: %d, Valid: %v\n",
		passes,
		float64(duration)/float64(time.Second),
		float64(duration)/float64(time.Second)/float64(passes),
		p.size,
		count,
		p.validateResults())
}
