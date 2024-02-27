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

	//wins, losses := solver.Minimax(ctx)

	wins, losses, moveCounts, totalBestScore, totalLowestScore, maxTileCounts := solver.MonteCarlo(ctx)
	//wins, losses := solver.PureRandom(ctx)

	fmt.Printf("Wins: %d, Losses: %d\n", wins, losses)
	fmt.Printf("\nBest Score: %.2f\n", float64(totalBestScore)/float64(wins+losses))
	fmt.Printf("Lowest Score: %.2f\n", float64(totalLowestScore)/float64(wins+losses))

	fmt.Println("\nMove Counts:")
	for move, count := range moveCounts {
		fmt.Printf("%s: %dx\n", move, count)
	}

	fmt.Println("\nMax Tile Counts:")
	for maxTile, count := range maxTileCounts {
		fmt.Printf("%d: %dx\n", maxTile, count)
	}

	elapsed := time.Since(start)
	fmt.Println("\nFinished in:", elapsed)

}
