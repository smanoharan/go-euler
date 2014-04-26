package main

func problem92() string {
    const max = 10*1000*1000 + 1
    visited := NewBitSet(max + 1)
    ends89 := NewBitSet(max + 1)

    NextLink := func(i int) int {
        next := 0
        for _, d := range toDigits(i) {
            next += d * d
        }
        return next
    }

    count := 0
    for i := 2; i < max; i++ {
        for j := i; true; j = NextLink(j) {
            if j < max && visited.Get(j) {
                if ends89.Get(j) {
                    count++
                }
                break
            }
            if j == 89 || j == 1 {
                if j == 89 {
                    if j < max {
                        ends89.Set(j)
                    }
                    count++
                }
                if j < max {
                    visited.Set(j)
                }
                break
            }
        }
    }

    return itoa(count)
}

func problem95() string {
    const max = 1000*1000 + 1

    type Num struct {
        DivSum, ChainLen, ChainStart int
        Visited                      bool
    }

    nums := make([]Num, max)
    for i := range nums {
        if i < 1 {
            nums[i].Visited = true
            continue
        }
        for j := i * 2; j < max; j += i {
            nums[j].DivSum += i // since i divides j
        }
    }

    // look for chains
    chain := make([]int, max)
    for i := range nums {
        n := nums[i]

        if n.Visited {
            continue
        }

        chi, cur := 0, i
        chain[chi] = cur

        for true {
            if nums[cur].Visited {
                if nums[cur].ChainLen == 0 {
                    // if chain len is 0, then cur is in chain (but not i)
                    // the prefix of the chain until the first occurence of cur
                    // must be considered a sink.
                    // The rest is a proper chain.
                    cs := 1
                    for ; chain[cs] != cur; cs++ {
                    }
                    cLen := chi - cs
                    nums[chain[cs]].ChainLen = cLen
                    for ; chi > cs; chi-- {
                        nums[chain[chi]].ChainLen = cLen
                    }
                }
                break
            }

            // advance to next link in the chain
            nums[cur].Visited = true
            cur = nums[cur].DivSum
            if cur == i || cur == 0 || cur >= max {
                // returned to start or next value is not valid
                break
            }

            chi++
            chain[chi] = cur
        }

        // apply found chain length to all members of the chain
        cLen := chi + 1
        if cur != i {
            cLen = -1 // reached a sink
            chi--
        }
        for ; chi >= 0; chi-- {
            nums[chain[chi]].ChainLen = cLen
        }

    }

    maxCL, minI := -1, 0
    for i, n := range nums {
        if n.ChainLen > maxCL {
            maxCL, minI = n.ChainLen, i
        }
    }
    return itoa(minI)
}

func problem97() string {
    const mod = int64(10 * 1000 * 1000 * 1000)

    quickExpMod := func(base, exp int) int64 {
        prod := int64(1)
        for mult, mask := int64(base), 1; mask <= exp; mult, mask = SafeMultMod(mult, mult, mod), mask<<1 {
            if (exp & mask) > 0 {
                prod = SafeMultMod(prod, mult, mod)
            }
        }
        return prod
    }

    prime := quickExpMod(2, 7830457)
    prime = (28433*prime + 1) % mod
    return i64toa(prime)
}

func problem99() string {

    max := float64(0)
    maxLN := 0

    lines := ReadAllLines("data/p99.txt")
    for i := 0; i < len(lines); i++ {
        parts := splitByComma(lines[i])
        base, exp := atoi(parts[0]), atoi(parts[1])
        curVal := logExp(base, exp)
        if curVal > max {
            max = curVal
            maxLN = i + 1 // 1-based line numbers
        }
    }

    return itoa(maxLN)
}
