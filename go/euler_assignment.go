package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
 * Even Fibonacci numbers
 *
 * By considering the terms in the Fibonacci sequence whose values do not
 * exceed four million, find the sum of the even-valued terms.
 */
func problem2() int {
	var fib func(int, int, int) int
	fib = func(prev, cur, sum int) int {
		if cur < 4000000 {
			if cur%2 == 0 {
				sum += cur
			}
			return fib(cur, prev+cur, sum)
		}
		return sum
	}
	return fib(1, 2, 0)
}

/*
 * Largest palindrome product
 *
 * A palindromic number reads the same both ways. The largest palindrome made
 * from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */
func problem4() int {
	// Unlike most languages, Go supports multiple encodings. Hence no unified reverse.
	// UTF-8 is used by default.
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	isPalindrome := func(n int) bool {
		s := strconv.Itoa(n)
		if s == reverse(s) {
			return true
		}
		return false
	}

	var max int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			prod := i * j
			if isPalindrome(prod) && prod > max {
				max = prod
			}
		}
	}
	return max
}

/*
 * Special Pythagorean triplet
 *
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for
 * which, a^2 + b^2 = c^2
 *
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 *
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.

Dickson's method

a² + b² = c²
a < b < c

a = r + s
b = r + t
c = r + s + t

c must be odd
r must be even
*/
func problem9() int {
	// Perhaps a bit naive implementation.
	factor := func(n int) (ret []int) {
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				ret = append(ret, i)
			}
		}
		return
	}
	type Triplet struct {
		a, b, c int
	}
	prodTriplet := func(r, s, t int) Triplet {
		return Triplet{r + s, r + t, r + s + t}
	}
	genAllTriplets := func(r int, sl, tl []int) (ret []Triplet) {
		for i := 0; i < len(tl); i++ {
			tr := prodTriplet(r, sl[i], tl[i])
			if tr.c%2 == 1 && (tr.a < tr.b && tr.a < tr.c) && (tr.b < tr.c) {
				ret = append(ret, tr)
			}
		}
		return
	}
	// Because Dickson's method!
	dickson := func(r int) int { return (r * r) / 2 }

	// Arbitrary, but sufficient, limit. r must always be even.
	for r := 2; r < 200; r += 2 {
		fl := factor(dickson(r))
		// Grouping the smallest and largest values to be used together.
		ss, ts := fl[:len(fl)/2], fl[len(fl)/2:]
		sort.Sort(sort.Reverse(sort.IntSlice(ts)))
		all := genAllTriplets(r, ss, ts)
		for i := 0; i < len(all); i++ {
			tr := all[i]
			if tr.a+tr.b+tr.c == 1000 {
				return tr.a * tr.b * tr.c
			}
		}
	}
	// No value found...
	return 0
}

/*
 * Maximum path sum I
 *
 * By starting at the top of the triangle below and moving to adjacent numbers
 * on the row below, the maximum total from top to bottom is 23.
 *
 *      3
 *     7 4
 *    2 4 6
 *   8 5 9 3
 *
 * That is, 3 + 7 + 4 + 9 = 23.
 *
 * Find the maximum total from top to bottom of the given triangle with 15
 * rows:
 */
func problem18() int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	type Data []int
	var data = []Data{}
	snl := strings.Split(input18, "\n")
	for i := 1; i < len(snl)-1; i++ {
		sints := strings.Split(snl[i], " ")
		var nums []int
		for _, n := range sints {
			cn, _ := strconv.Atoi(n) // There is nothing but ints in data, ignore err.
			nums = append(nums, cn)
		}
		data = append(data, nums)
	}
	for line := len(data) - 2; line >= 0; line-- {
		for pos := 0; pos < len(data[line]); pos++ {
			e1, e2 := data[line+1][pos], data[line+1][pos+1]
			cur := data[line][pos]
			cur += max(e1, e2)
			data[line][pos] = cur
		}
	}

	return data[0][0]
}

/*
 * Maximum path sum II
 *
 * By starting at the top of the triangle below and moving to adjacent numbers
 * on the row below, the maximum total from top to bottom is 23.
 *
 *    3
 *   7 4
 *  2 4 6
 * 8 5 9 3
 *
 * That is, 3 + 7 + 4 + 9 = 23.
 *
 * Find the maximum total from top to bottom in the given triangle with
 * one-hundred rows.
 *
 * NOTE: This is a much more difficult version of Problem 18. It is not
 * possible to try every route to solve this problem, as there are 2^99
 * altogether! If you could check one trillion (10^12) routes every second it
 * would take over twenty billion years to check them all. There is an
 * efficient algorithm to solve it. ;o)
 */
func problem67() int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	type Data []int
	var data = []Data{}
	snl := strings.Split(input67, "\n")
	for i := 1; i < len(snl)-1; i++ {
		sints := strings.Split(snl[i], " ")
		var nums []int
		for _, n := range sints {
			cn, _ := strconv.Atoi(n) // There is nothing but ints in data, ignore err.
			nums = append(nums, cn)
		}
		data = append(data, nums)
	}
	for line := len(data) - 2; line >= 0; line-- {
		for pos := 0; pos < len(data[line]); pos++ {
			e1, e2 := data[line+1][pos], data[line+1][pos+1]
			cur := data[line][pos]
			cur += max(e1, e2)
			data[line][pos] = cur
		}
	}

	return data[0][0]
}

func resetColours()  { fmt.Print("\033[0m") }
func passedColours() { fmt.Print("\033[32m") }
func failedColours() { fmt.Print("\033[31;1m") }

func main() {
	type Tfun struct {
		fn  func() int
		ret int
	}
	// Could use runtime and reflect to get the func name, but that'd not be pretty.
	probs := map[string]Tfun{
		"problem2":  Tfun{problem2, 4613732},
		"problem4":  Tfun{problem4, 906609},
		"problem9":  Tfun{problem9, 31875000},
		"problem18": Tfun{problem18, 1074},
		"problem67": Tfun{problem67, 7273},
	}
	for n, t := range probs {
		switch out := t.fn(); out {
		case t.ret:
			passedColours()
			fmt.Printf("Passed %s\n", n)
			resetColours()
		default:
			failedColours()
			fmt.Printf("%s → %v, but expected %v\n", n, out, t.ret)
			resetColours()
		}
	}
}
