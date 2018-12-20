package days

import (
	"adventofcode/inputs"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type claim struct {
	id int
	x int
	y int
	xend int
	yend int
}

type PoI struct {
	ID    int
	N     int
	Start bool
}

type ByN []PoI

func (p ByN) Len() int           { return len(p) }
func (p ByN) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByN) Less(i, j int) bool { return p[i].N < p[j].N }

func VertOverlap(vpoi ByN) int {
	active := 0
	last := 0
	overlap := 0

	for _, v := range vpoi {
		if active > 1 {
			overlap += v.N - last
		}

		if v.Start {
			active++
		} else {
			active--
		}

		last = v.N
	}

	return overlap
}

func Day3Part1() string {
	input := inputs.Input3
	claims := make(map[int]claim, 0)
	horPoI := make(ByN, 0)

	for _,v := range strings.Split(input, "\n") {
		id := 0
		x := 0
		y := 0
		xend := 0
		yend := 0

		_,_ = fmt.Sscanf(v, "#%d @ %d,%d: %dx%d", &id, &x, &y, &xend, &yend)
		claims[id] = claim{id, x, y, xend + x,yend + y}
		horPoI = append(horPoI, ByN{{id, x, true}, {id, x + xend, false}}...)
	}

	sort.Sort(horPoI)

	overlap := 0
	lastactivity := 0
	active := make(map[int]bool)

	for _,v := range horPoI {
		vpoi := make(ByN, 0)
		for k := range active {
			c := claims[k]
			vpoi = append(vpoi, ByN{{c.id, c.y, true}, {c.id, c.yend, false}}...)
		}

		sort.Sort(vpoi)
		overlap += VertOverlap(vpoi) * (v.N - lastactivity)

		if v.Start {
			active[v.ID] = true
		} else {
			delete(active, v.ID)
		}
		lastactivity = v.N
	}

	return strconv.Itoa(overlap)
}