package main

import (
	"fmt"
	"math/rand"
	"flag"
	"time"
)

func getNeighbors(board [][] int, r int, c int) []int {
	res := make([] int, 0)
	moves := [][] int {{1,0}, {-1,0}, {0,1}, {0,-1}, {1,1}, {1,-1}, {-1,1}, {-1,-1}}
	
	for _, m := range moves {
		tempR := r + m[0]
		tempC := c + m[1]
		if tempR >= 0 && tempR < len(board) && tempC >= 0 && tempC < len(board[0]) {
			res = append(res, board[tempR][tempC])
		}
	}
	return res
}

func filter(vals [] int, f func(int) bool) [] int {
	res := make([] int, 0)
	for _, v := range vals {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func aliveCount(board [][] int, r int, c int) int {
	return len(filter(getNeighbors(board, r, c), func(v int) bool {
		return v == 1
	}))
}

func runGeneration(board, next [][] int) ([][] int, [][]int){
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			count := aliveCount(board, r, c)
			if board[r][c] == 1 {
				if count < 2 {
					next[r][c] = 0
				}
				if count == 2 || count == 3 {
					next[r][c] = 1
				}
				if count > 3 {
					next[r][c] = 0
				}
			} else {
				if count == 3 {
					next[r][c] = 1
				} else {
					next[r][c] = board[r][c]
				}
			}
		}
	}

	return next, board
}
 

func main() {
	start := time.Now()
	var rows, cols, seed, iters int
	var print bool

	flag.IntVar(&rows, "rows", 10, "How many rows the board should have.")
	flag.IntVar(&cols, "cols", 10, "How many cols the board should have.")
	flag.IntVar(&seed, "seed", 10, "Seed to initialize board.")
	flag.IntVar(&iters, "iters", 10, "How many iterations to execute.")
	flag.BoolVar(&print, "print", false, "Print to stdout.")
	
	flag.Parse()

	// Initialization
	rand.Seed(int64(seed))
	board := make([][] int, rows)
	nxt := make([][] int, rows)
	for r := 0; r < rows; r++ {
		board[r] = make([] int, cols)
		nxt[r] = make([] int, cols)
		for c := 0; c < cols; c++ {
			board[r][c] = rand.Intn(2)
		}
	}

	for i := 0; i < iters; i++ {
		if print {
			fmt.Println("GEN", i+1)
			for r := 0; r < rows; r++ {
				fmt.Println(board[r])
			}
		}
		board, nxt = runGeneration(board, nxt)
		
	}

	fmt.Println(time.Since(start))


}