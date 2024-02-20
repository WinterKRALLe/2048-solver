package main

import "Solver2048/solver"

func main() {
	//start := time.Now()
	//
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//sigCh := make(chan os.Signal, 1)
	//signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	//
	//go func() {
	//	<-sigCh
	//	cancel()
	//}()
	//
	//wins, losses := solver.Minimax(ctx)
	//
	//fmt.Printf("Wins: %d, Losses: %d\n", wins, losses)
	//
	//elapsed := time.Since(start)
	//fmt.Println("Finished in:", elapsed)
	solver.MonteCarlo()
}
