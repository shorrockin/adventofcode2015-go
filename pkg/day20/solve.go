package day20

import (
	"math"
)

func Solve(presentsTarget int, multiplier int, maxHouses int) int {
	houseNumber := 1
	for {
		if sumOfDivisors(houseNumber, multiplier, maxHouses) >= presentsTarget {
			return houseNumber
		}
		houseNumber++
	}
}

func sumOfDivisors(houseNumber int, presentsPerHouse int, maxDeliveries int) int {
	sum := 0
	sqrt := int(math.Sqrt(float64(houseNumber)))

	for divisor := 1; divisor <= sqrt; divisor++ {
		if houseNumber%divisor == 0 {
			// when true we know that i is a divisor and houseNumber/i
			// is also divisor. example, if house is 36 i and 3, then we
			// know we have 3 and 12 (house/3) divisors
			pairedDivisor := (houseNumber / divisor)

			if pairedDivisor <= maxDeliveries {
				sum += divisor * presentsPerHouse
			}

			// avoid adding the square root twice if n is a perfect square, example
			// if houseNumber is 36, and i is 6 then pairedDivisor would also
			// be 6
			if divisor != pairedDivisor {
				if (houseNumber / pairedDivisor) <= maxDeliveries {
					sum += pairedDivisor * presentsPerHouse
				}
			}
		}
	}
	return sum
}
