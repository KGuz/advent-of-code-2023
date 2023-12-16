package dbg

import "fmt"

const (
	CYAN      = "\033[96m"
	PURPLE    = "\033[95m"
	BLUE      = "\033[94m"
	YELLOW    = "\033[93m"
	GREEN     = "\033[92m"
	RED       = "\033[91m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	END       = "\033[0m"
)

func Print(graph [][]byte) {
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			fmt.Printf("%c", graph[i][j])
		}
		fmt.Println()
	}
}
