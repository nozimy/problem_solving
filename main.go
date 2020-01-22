package main

import (
	"flag"
	"problem_solving/problems"
)

func main() {
	var problemNum int
	flag.IntVar(&problemNum, "problemNum", 1, "problem number")
	flag.Parse()

	problemsArr := []func(){
		problems.RunFindCommonNumberProblem,
		problems.RunNumGameProblem,
		problems.RunGetMaxDepthProblem,
		problems.RunReverseLinkedListProblem,
		problems.RunMaxVisitorsProblem,
		problems.SwapWithoutAdditionalMemory,
	}

	problemsArr[problemNum]()
}
