package main

import (
	"Solver2048/solver"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		cancel()
	}()

	minimaxWins, minimaxLosses, minimaxMoveCounts, minimaxTotalBestScore, minimaxTotalLowestScore, minimaxMaxTileCounts := solver.Minimax(ctx)
	monteCarloWins, monteCarloLosses, monteCarloMoveCounts, monteCarloTotalBestScore, monteCarloTotalLowestScore, monteCarloMaxTileCounts := solver.MonteCarlo(ctx)
	pureRandomWins, pureRandomLosses, pureRandomMoveCounts, pureRandomTotalBestScore, pureRandomTotalLowestScore, pureRandomMaxTileCounts := solver.PureRandom(ctx)

	fmt.Println("\n########## MINIMAX ##########")

	fmt.Printf("Wins: %d, Losses: %d\n", minimaxWins, minimaxLosses)
	fmt.Printf("\nBest Score: %.2f\n", float64(minimaxTotalBestScore)/float64(minimaxWins+minimaxLosses))
	fmt.Printf("Lowest Score: %.2f\n", float64(minimaxTotalLowestScore)/float64(minimaxWins+minimaxLosses))

	fmt.Println("\nMove Counts:")
	for move, count := range minimaxMoveCounts {
		fmt.Printf("%s: %dx\n", move, count)
	}

	fmt.Println("\nMax Tile Counts:")
	for maxTile, count := range minimaxMaxTileCounts {
		fmt.Printf("%d: %dx\n", maxTile, count)
	}

	fmt.Println("\n########## MONTE CARLO ##########")

	fmt.Printf("\nWins: %d, Losses: %d\n", monteCarloWins, monteCarloLosses)
	fmt.Printf("\nBest Score: %.2f\n", float64(monteCarloTotalBestScore)/float64(monteCarloWins+monteCarloLosses))
	fmt.Printf("Lowest Score: %.2f\n", float64(monteCarloTotalLowestScore)/float64(monteCarloWins+monteCarloLosses))

	fmt.Println("\nMove Counts:")
	for move, count := range monteCarloMoveCounts {
		fmt.Printf("%s: %dx\n", move, count)
	}

	fmt.Println("\nMax Tile Counts:")
	for maxTile, count := range monteCarloMaxTileCounts {
		fmt.Printf("%d: %dx\n", maxTile, count)
	}

	fmt.Println("\n########## PURE RANDOM ##########")

	fmt.Printf("\nWins: %d, Losses: %d\n", pureRandomWins, pureRandomLosses)
	fmt.Printf("\nBest Score: %.2f\n", float64(pureRandomTotalBestScore)/float64(pureRandomWins+pureRandomLosses))
	fmt.Printf("Lowest Score: %.2f\n", float64(pureRandomTotalLowestScore)/float64(pureRandomWins+pureRandomLosses))

	fmt.Println("\nMove Counts:")
	for move, count := range pureRandomMoveCounts {
		fmt.Printf("%s: %dx\n", move, count)
	}

	fmt.Println("\nMax Tile Counts:")
	for maxTile, count := range pureRandomMaxTileCounts {
		fmt.Printf("%d: %dx\n", maxTile, count)
	}

	elapsed := time.Since(start)
	fmt.Println("\nFinished in:", elapsed)

}
