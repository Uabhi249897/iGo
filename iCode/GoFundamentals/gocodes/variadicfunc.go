package gocodes

import "fmt"

func Variadicfunc() {
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16) //undefined number of values
	fmt.Println(bestFinish)
}

//values passed in finishes as slice/list
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	return best
}
