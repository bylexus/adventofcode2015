package main

import (
	"os"
	"sync"
	"time"

	"alexi.ch/aoc2015/lib"
	"alexi.ch/aoc2015/lib/types"
	"alexi.ch/aoc2015/problems"
)

type Problems map[string]types.AocProblem

type Durations struct {
	duration1 time.Duration
	duration2 time.Duration
}

func runProblem(problem types.AocProblem) (duration Durations) {
	problem.Init()
	duration.duration1 = lib.MeasureTime(problem.Run1)
	duration.duration2 = lib.MeasureTime(problem.Run2)
	return
}

func runAll(probs Problems) {
	wg := sync.WaitGroup{}
	results := make([]types.AoCResult, 0, 25)
	resultChan := make(chan types.AoCResult)
	start := time.Now()

	for key, problem := range probs {
		wg.Add(1)
		go func(p types.AocProblem, k string) {
			p.Init()
			d := runProblem(p)
			res := types.AoCResult{
				Problem:       p,
				Key:           k,
				TimeSolution1: d.duration1,
				TimeSolution2: d.duration2,
			}
			resultChan <- res
			wg.Done()
		}(problem, key)
	}

	// collect results:
	doneChannel := make(chan bool)
	go func() {
		for res := range resultChan {
			results = append(results, res)
		}
		close(doneChannel)
	}()
	wg.Wait()
	totalDuration := time.Since(start)
	close(resultChan)
	<-doneChannel

	lib.OutputResultList(results, totalDuration)
}

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the problem name (e.g. 'day01') as parameter.")
	}

	problemName := os.Args[1]

	problemMap := make(Problems)
	problemMap["day01"] = &problems.Day01{}
	problemMap["day02"] = &problems.Day02{}
	problemMap["day03"] = &problems.Day03{}
	problemMap["day04"] = &problems.Day04{}
	problemMap["day05"] = &problems.Day05{}
	problemMap["day06"] = &problems.Day06{}

	if problemName == "all" {
		runAll(problemMap)
	} else {
		problem, defined := problemMap[problemName]
		if defined == true {
			lib.OutputTitle(problem.GetName())
			solution := runProblem(problem)
			lib.OutputSolution(1, solution.duration1, problem.GetSolution1())
			lib.OutputSolution(2, solution.duration2, problem.GetSolution2())
		} else {
			panic("Oops - Problem not found - is it defined in main?")
		}
	}
}
