package main

// sum from start to finish (inclusive), only including every step'th number.
// requires: step \neq 0 and start < finish
// no requirements are checked
func sum(start, step, finish int) int {
    count := (finish - start) / step
    actualFinish := start + (count * step)
    return (start + actualFinish) * (count + 1) / 2 // +1 to include start and finish
}

// the nth square pyramidal number
func squarePyramidal(n int) int {
    return (n * (n + 1) * (2*n + 1)) / 6
}

// sum all multiples of 3 or 5 under 1000
func problem1() string {
    start := 0
    finish := 999
    return itoa(sum(start, 3, finish) + sum(start, 5, finish) - sum(start, 15, finish))
}

// sum all even fibonacci terms not exceeding 4 million
func problem2() string {
    sum := 0
    max := 4000000
    for prev, cur := 1, 1; cur <= max; prev, cur = cur, prev+cur {
        if (cur & 1) == 0 { // is even
            sum += cur
        }
    }
    return itoa(sum)
}

// find the largest prime factor of 600851475143
func problem3() string {
    max := uint64(600851475143)
    sievemax := int(SqrtU64(max)) + 1
    bitset := NewBitSet(sievemax + 1)

    largest_prime := 0
    for p := 2; p <= sievemax; p++ {
        if !bitset.Get(p) {

            for pm := 2 * p; pm <= sievemax; pm += p {
                bitset.Set(pm)
            }

            if 0 == (max % uint64(p)) {
                largest_prime = p
            }
        }
    }

    return itoa(largest_prime)
}

func toDigits(number int) []int {
    var digits [10]int // 2^31 is at most 2B, i.e. 10 digits
    c := 0
    for n := number; n > 0; c++ {
        d := n / 10
        digits[c] = n - (d * 10)
        n = d
    }
    return digits[0:c]
}

func isPalindrome(number int) bool {
    digits := toDigits(number)
    for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
        if digits[i] != digits[j] {
            return false
        }
    }
    return true
}

// find the largest palindrome which is a product of two 3-digit numbers
func problem4() string {
    min, max := 100, 999

    maxP := 0
    for i := max; i >= min; i-- {
        for j := i; j >= min; j-- {

            p := i * j
            if p <= maxP {
                break // i*j is decreasing inside the inner loop
            } else if isPalindrome(p) {
                maxP = p
            }
        }

        if (i-1)*(i-1) <= maxP {
            break // (i-1)*(i-1) is the max of all future iterations
        }
    }

    return itoa(maxP)
}

// find the smallest positive number which is divisible by all numbers from 1 to 20
func problem5() string {
    const (
        min, max = 1, 20
        maxA     = max + 1
    )

    // compute prime factors of each number
    var maxPF [maxA]int    // max prime factor for each number
    var pF [maxA][maxA]int // First index is number, Second is prime factor, cell is count.

    bitset := NewBitSet(maxA)
    for p := 2; p <= max; p++ {
        if bitset.Get(p) {
            // p is composite, with at least one prime factor (stored in maxPF)
            f1 := maxPF[p]
            f2 := p / f1

            // f2 < p, so the prime decomposition of f2 is known.
            // prime decomposition of p is that of f2 with one more f1.
            for i := 0; i < maxA; i++ {
                pF[p][i] = pF[f2][i]
            }
            pF[p][f1]++
        } else {

            // p is prime
            for pm := 2 * p; pm <= max; pm += p {
                bitset.Set(pm)
                maxPF[pm] = p
            }

            // prime decomposition of p is just p
            pF[p][p] = 1
        }
    }

    // find lcm: find max number of each prime needed:
    lcm := int64(1)
    for p := 2; p <= max; p++ {
        if !bitset.Get(p) {
            maxF := 0
            for i := 1; i < maxA; i++ {
                maxF = Max2i(pF[i][p], maxF)
            }
            lcm *= PowI64(int64(p), int64(maxF))
        }
    }

    return i64toa(lcm)
}

// find the difference between sum of squares and square of sum of 1..100
func problem6() string {
    n := 100
    sumn := sum(0, 1, n)
    sqOfSum := sumn * sumn
    sumOfSq := squarePyramidal(n)  // \sum_k k^2 is the kth square pyramid number
    return itoa(sqOfSum - sumOfSq) // sqOfSum >= sumOfSq for all n >= 0
}

// find the 10,001st prime
func problem7() string {

    n := 10001
    sievemax := 1000000 // just a guess as to the required bound
    bitset := NewBitSet(sievemax + 1)

    c := 0
    for p := 2; p <= sievemax; p++ {
        if !bitset.Get(p) {
            c++
            if c >= n {
                return itoa(p)
            }

            for pm := 2 * p; pm <= sievemax; pm += p {
                bitset.Set(pm)
            }
        }
    }

    return "sievemax wasn't large enough (only found " + itoa(c) + " primes)"
}

const ZERO_CHAR = int('0')

func charToNum(c byte) int {
    return int(c) - ZERO_CHAR
}

func product(arr []int, size int) int {
    prod := 1
    for i := 0; i < size; i++ {
        prod *= arr[i]
    }
    return prod
}

func problem8() string {
    number :=
        "73167176531330624919225119674426574742355349194934" +
            "96983520312774506326239578318016984801869478851843" +
            "85861560789112949495459501737958331952853208805511" +
            "12540698747158523863050715693290963295227443043557" +
            "66896648950445244523161731856403098711121722383113" +
            "62229893423380308135336276614282806444486645238749" +
            "30358907296290491560440772390713810515859307960866" +
            "70172427121883998797908792274921901699720888093776" +
            "65727333001053367881220235421809751254540594752243" +
            "52584907711670556013604839586446706324415722155397" +
            "53697817977846174064955149290862569321978468622482" +
            "83972241375657056057490261407972968652414535100474" +
            "82166370484403199890008895243450658541227588666881" +
            "16427171479924442928230863465674813919123162824586" +
            "17866458359124566529476545682848912883142607690042" +
            "24219022671055626321111109370544217506941658960408" +
            "07198403850962455444362981230987879927244284909188" +
            "84580156166097919133875499200524063689912560717606" +
            "05886116467109405077541002256983155200055935729725" +
            "71636269561882670428252483600823257530420752963450"

    maxi := len(number)
    const M = 5

    var n [M]int // a circular array of the last M numbers

    maxp := 0
    c := 0
    for i := 0; i < maxi-M; i++ {
        n[c] = charToNum(number[i])
        c++
        if c == M {
            c = 0
        }
        maxp = Max2i(maxp, product(n[:], M))
    }

    return itoa(maxp)
}

// find (a,b,c) s.t. a+b+c = 1000 and a^2 + b^2 = c^2
func problem9() string {
    max := 1000

    maxa := max - 2
    for a := 1; a < maxa; a++ {
        maxb := max - a - 1
        for b := 1; b < maxb; b++ {
            c := max - a - b
            if a*a+b*b == c*c {
                return itoa(a * b * c)
            }
        }
    }
    return "Logic Error: Not Found"
}

// find the sum of all primes under 2 million
func problem10() string {
    sievemax := 2000000
    bitset := NewBitSet(sievemax)

    sump := int64(0)
    for p := 2; p < sievemax; p++ {
        if !bitset.Get(p) {
            sump += int64(p)
            for pm := 2 * p; pm < sievemax; pm += p {
                bitset.Set(pm)
            }
        }
    }

    return i64toa(sump)
}
