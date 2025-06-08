package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // hilangkan newline di akhir
	fmt.Println("Nama:", name)
}
