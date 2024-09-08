package main

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func count(s string, b byte) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			c++
		}
	}
	return c
}
