// The primes command was automatically generated by Shenzhen Go.
package main

import (
	"fmt"
	"sync"
)

var (
	div2 = make(chan int, 0)
	div3 = make(chan int, 0)
	out  = make(chan int, 0)
	raw  = make(chan int, 0)
)

func main() {

	var wg sync.WaitGroup

	// Filter divisible by 2
	wg.Add(1)

	go func() {
		defer wg.Done()

		for x := range raw {

			if x == 2 || x%2 != 0 {
				div2 <- x
			}
		}

		close(div2)

	}()

	// Filter divisible by 3
	wg.Add(1)

	go func() {
		defer wg.Done()

		var multWG sync.WaitGroup
		multWG.Add(3)
		for n := 0; n < 3; n++ {
			go func(instanceNumber int) {
				defer multWG.Done()
				for x := range div2 {

					if x == 3 || x%3 != 0 {
						div3 <- x
					}
				}
			}(n)
		}
		multWG.Wait()

		close(div3)

	}()

	// Filter divisible by 5
	wg.Add(1)

	go func() {
		defer wg.Done()

		for x := range div3 {

			if x == 5 || x%5 != 0 {
				out <- x
			}
		}

		close(out)

	}()

	// Generate integers ≥ 2
	wg.Add(1)

	go func() {
		defer wg.Done()

		for i := 2; i < 49; i++ {
			raw <- i
		}
		close(raw)
	}()

	// Print output
	wg.Add(1)

	go func() {
		defer wg.Done()

		for n := range out {
			fmt.Println(n)
		}

	}()

	// Wait for the end
	wg.Wait()
}