package main

import (
	"io"
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
	sortedKeys := make([]string, 0, len(l.Wins))

	for key := range l.Wins {
		sortedKeys = append(sortedKeys, key)
	}

	sort.SliceStable(sortedKeys, func(i, j int) bool {
		return l.Wins[sortedKeys[i]] > l.Wins[sortedKeys[j]]
	})

	return sortedKeys
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, wrt io.Writer) {
	for _, team := range r.Ranking() {
		wrt.Write([]byte(team + "\n"))
	}
}
