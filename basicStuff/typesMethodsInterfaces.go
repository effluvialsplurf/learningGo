package main

import (
	"sort"
)

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Wins  map[string]int
	Teams []Team
}

func (l *League) MatchResult(name1, name2 string, score1, score2 int) {
	l.Wins[name1] = score1
	l.Wins[name2] = score2
}

func (l League) Ranking() []string {
	sortedMap := make(map[string]int)
	for name, score := range l.Wins {
		sortedMap[name] = score
	}

	sortKeys := make([]string, 0, len(sortedMap))

	for key := range sortedMap {
		sortKeys = append(sortKeys, key)
	}

	sort.SliceStable(sortKeys, func(i, j int) bool {
		return sortedMap[sortKeys[i]] > sortedMap[sortKeys[j]]
	})

	return sortKeys
}
