package main

import (
	"fmt"
	"math"
	"time"
)

type primeBase struct {
	size uint
	sqrt uint
}

type primeBool struct {
	primeBase
	isPrime [1000001]bool
}

var myDict = map[uint]uint{
	10:         4,
	100:        25,
	1000:       168,
	10000:      1229,
	100000:     9592,
	1000000:    78498,
	10000000:   664579,
	100000000:  5761455,
	1000000000: 50847534,
}

var limit uint = 1000000
var listing [1000001]bool
var duration float64 = 10

func main() {
	var passes uint
	start := time.Now()
	p := primeBool{}

	for {
		p.Set(limit)
		p.Run()
		passes++
		if time.Since(start).Seconds() >= duration {
			break
		}
	}
	tD := time.Since(start)
	p.PrintResults(false, tD, passes)
}

func (p *primeBool) Set(index uint) {
	p.size = index
	p.isPrime = listing
	p.sqrt = uint(math.Sqrt(float64(index)))
	p.isPrime[2] = true
	var i uint
	for i = 3; i <= index; i = i + 2 {
		if i%2 != 0 {
			p.isPrime[i] = true
		}
	}

}

func (p *primeBool) Run() {
	var factor uint = 3
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

func (p *primeBool) clear(factor uint) {
	for num := factor * 3; num <= p.size; num += factor * 2 {
		p.isPrime[num] = false
	}
}

func (p *primeBool) PrintResults(showResults bool, duration time.Duration, passes uint) {
	primes := make([]uint, 0)
	var num uint
	for num = 2; num <= p.size; num++ {
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

func (p *primeBool) countPrimes() uint {
	var count uint
	var i uint
	for i = 0; i < p.size; i++ {
		if p.isPrime[i] {
			count++
		}
	}
	return count
}
