package types

import (
	"time"
)

type EN struct{}
type n struct{}
type na struct{}

type B struct {
	I int
}

type A struct {
	EIa, EIb int
	EI       int
	ES       string
	EB       B
	EN
	ET time.Time
	time.Time
	EIS struct {
		I int
	}
	EISL    []int
	EIA     [2]int
	EMap    map[int]int
	Fn      func(int) int
	ESa, sb string
	is      struct {
		S string
	}
	isi struct {
		S string
		i int
	}
	i int
	n
	na  na
	nap *na
	string
	*int
	b     B
	bp    *B
	t     time.Time
	tp    *time.Time
	isl   []int
	islsl [][]int
	ia    [2]int
	iaa   [2][2]int
	m     map[int]string
	mm    map[int]map[int]string
	mn    map[int]n
	fnii  func(int) int
	fnin  func(int) n
}
