package main

import "math/rand/v2"

type WordleDict struct {
	elems []string
}

func initWordleDict() *WordleDict {
	var wordleDict WordleDict
	wordleDict.elems = []string{"BUGRA", "NESLI", "TUGAY", "BARIS", "KADIR", "BUSRA"}
	return &wordleDict
}

func (w *WordleDict) pickRandom() string {
	return w.elems[rand.IntN(len(w.elems))]
}
