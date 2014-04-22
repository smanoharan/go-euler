package main

func problem112() string {
	const max = 1000*1000*1000
	
	isBouncy := func(i int) bool {
		last := i % 10
		for i /= 10; i > 0 && last == i % 10; i /= 10 {}
		if i == 0 { return false }

		cur := i % 10
		isDesc := cur < last
		last = cur
		
		for i /= 10; i > 0; i /= 10 {
			cur = i % 10
			if (cur != last) && ((cur < last) != isDesc) { return true }
			last = cur
		}
		
		return false
	}

	bouncyCount := 0
	for i := 100; i < max; i++ {
		if isBouncy(i) { bouncyCount++ }

		if bouncyCount*100 == i*99 { 
			return itoa(i) 
		}
	}

	panic("No solution found.")
}
