// Develop a program which solves the following task:
// In the office each visitor’s arrival and departure times are registered. So at the end of the day for N
// visitors there are N pairs of values: the first value in pair is arrival time and the second – departure time.
// Need to find time interval during the day when there are maximum visitors in the office.

package problems

import (
	"fmt"
	"sort"
	"time"
)

const (
	pointTypeBegin int = 0
	pointTypeEnd   int = 1
)

type Pair struct {
	Arrival   time.Time
	Departure time.Time
}

type Point struct {
	pointType int
	time      time.Time
}

func RunMaxVisitorsProblem() {
	var pairs = []Pair{
		Pair{
			Arrival:   time.Date(2020, 01, 23, 10, 0, 0, 0, time.UTC),
			Departure: time.Date(2020, 01, 23, 11, 0, 0, 0, time.UTC),
		},
		Pair{
			Arrival:   time.Date(2020, 01, 23, 10, 30, 0, 0, time.UTC),
			Departure: time.Date(2020, 01, 23, 11, 0, 0, 0, time.UTC),
		},
		Pair{
			Arrival:   time.Date(2020, 01, 23, 12, 0, 0, 0, time.UTC),
			Departure: time.Date(2020, 01, 23, 13, 0, 0, 0, time.UTC),
		},
		Pair{
			Arrival:   time.Date(2020, 01, 23, 10, 0, 0, 0, time.UTC),
			Departure: time.Date(2020, 01, 23, 15, 0, 0, 0, time.UTC),
		},
	}

	maxCount, begin, end := findIntervalWithMaxVisitors(pairs)

	fmt.Printf("Max visitors: %d, Time Interval with max visitors: %+v - %+v\n", maxCount, begin, end)
}

func findIntervalWithMaxVisitors(pairs []Pair) (maxCount int, maxBegin, maxEnd time.Time) {
	points := make([]Point, 0)

	for _, p := range pairs {
		points = append(points, Point{
			pointType: pointTypeBegin,
			time:      p.Arrival,
		})
		points = append(points, Point{
			pointType: pointTypeEnd,
			time:      p.Departure,
		})
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].time.Before(points[j].time)
	})

	maxCount = -1
	visitors := 0

	for _, p := range points {
		if p.pointType == pointTypeBegin {
			visitors++
		} else {
			if visitors == maxCount {
				maxEnd = p.time
			}
			visitors--
		}

		if visitors > maxCount {
			maxCount = visitors
			maxBegin = p.time
		}
	}

	return maxCount, maxBegin, maxEnd
}
