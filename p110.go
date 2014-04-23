package main

func problem102() string {
	type Point struct {
		x, y int
	}

	// check if q lies on the same side of the origin p1--p2.
	sameSide := func(p1, p2, q Point) bool {
		// the line p1 <--> p2 is (p2x - p1x)y - (p2y - p1y)x = c
		// if c == 0 then origin lies on the line, thus origin is on the same side as q
		// given the inequality I(x0,y0) : c > (p2x - p1x)y0 - (p2y - p1y)x0,
		// then I(0,0) == I(qx,qy) iff q is on the same side as the origin.
		// equiv. same-side iff (c > 0) == (c > (p2x-p1x)qy - (p2y-p1y)qx)
		// equiv. same-side iff (c > 0) == (c - (p2x-p1x)qy + (p2y-p1y)qx > 0)
		dx, dy := p2.x - p1.x, p2.y - p1.y
		c := dx*p1.y - dy*p1.x
		if (c != dx*p2.y - dy*p2.x) { panic("Math error") }
		return c == 0 || ( (c > 0) == (c - dx*q.y + dy*q.x > 0) )
	}

	surroundsOrigin := func(a, b, c Point) bool {
		return sameSide(a, b, c) && sameSide(a, c, b) && sameSide(b, c, a)
	}

	toPoint := func(parts []string, i int) Point {
		xs, ys := parts[i], parts[i+1]
		return Point{atoi(xs),atoi(ys)}
	}

	count := 0
	for _, line := range ReadAllLines("data/p102.txt") {
		if len(line) == 0 { break }
		ps := splitByComma(line)
		a, b, c := toPoint(ps, 0), toPoint(ps, 2), toPoint(ps, 4)
		if surroundsOrigin(a, b, c) { count++ }
	}

	return itoa(count)
}
