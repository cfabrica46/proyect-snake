package main

import "fmt"

func main() {

	a := []int{1}

	b := [][]int{a, {1}, a}

	b[0][0] = 5

	fmt.Println(b)

	//------------------------------------
	seaNames := [][]string{{"shark", "octopus", "squid", "mantis shrimp"}, {"Sammy", "Jesse", "Drew", "Jamie"}}

	fmt.Println(seaNames)

	seaNames[0][1] = "pulpo"

	fmt.Println(seaNames)

}
