package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type logType int

const (
	_ logType = iota
	shiftStart
	sleep
	wake
)

func (t logType) String() string {
	switch t {
	case shiftStart:
		return "shift"
	case sleep:
		return "sleep"
	case wake:
		return "wake"
	default:
		return "<error>"
	}
}

type log struct {
	time    time.Time
	kind    logType
	guardId string
}

type shift struct {
	sleeps [60]bool
	guard  string
}

func (s shift) minutesAsleep() int {
	var sum = 0
	for _, asleep := range s.sleeps {
		if asleep {
			sum++
		}
	}
	return sum
}

type guard struct {
	id     string
	shifts []shift
}

func main() {
	logs := ScanLogs()

	shifts := make([]shift, 0, 100)

	var curShift shift
	curShift.guard = logs[0].guardId
	var i = 1
L:
	for ; ; i++ {
		l := logs[i]
		switch l.kind {
		case shiftStart:
			break L
		case sleep:
			for j := l.time.Minute(); j < 60; j++ {
				curShift.sleeps[j] = true
			}
		case wake:
			for j := l.time.Minute(); j < 60; j++ {
				curShift.sleeps[j] = false
			}
		}
	}

	for ; i < len(logs); i++ {
		l := logs[i]
		switch l.kind {
		case shiftStart:
			shifts = append(shifts, curShift)
			curShift = shift{guard: l.guardId}
		case sleep:
			for j := l.time.Minute(); j < 60; j++ {
				curShift.sleeps[j] = true
			}
		case wake:
			for j := l.time.Minute(); j < 60; j++ {
				curShift.sleeps[j] = false
			}
		}
	}
	shifts = append(shifts, curShift)

	var guards = make(map[string]guard)
	for _, s := range shifts {
		g := guards[s.guard]
		g.id = s.guard
		g.shifts = append(g.shifts, s)
		guards[s.guard] = g
	}

	var bg guard
	var bmin, slept int
	for _, g := range guards {
		var minutes [60]int
		for _, shift := range g.shifts {
			for minute, asleep := range shift.sleeps {
				if asleep {
					minutes[minute]++
				}
			}
		}
		for min, s := range minutes {
			if s > slept {
				slept = s
				bmin = min
				bg = g
			}
		}
	}
	id, err := strconv.ParseInt(bg.id, 0, 0)
	check(err)
	fmt.Println(int(id) * bmin)
	return

}

var lineFormat = regexp.MustCompile(`\[(.*)] (.*)`)
var guardId = regexp.MustCompile(`Guard #(\d+) begins shift`)

func ScanLogs() []log {
	logs := make([]log, 0, 1200)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		l := log{}

		m := lineFormat.FindStringSubmatch(s.Text())
		if m == nil {
			panic(fmt.Sprintf("Bad text format: %#v", s.Text()))
		}

		t, err := time.Parse("2006-01-02 15:04", m[1])
		check(err)
		l.time = t

		msg := m[2]
		if msg == "wakes up" {
			l.kind = wake
		} else if msg == "falls asleep" {
			l.kind = sleep
		} else {
			l.kind = shiftStart
			m = guardId.FindStringSubmatch(msg)
			if m == nil {
				panic(fmt.Sprintf("Missing guard id: %#v", s.Text()))
			}
			l.guardId = m[1]
		}

		logs = append(logs, l)
	}
	check(s.Err())

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].time.Before(logs[j].time)
	})

	return logs
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
