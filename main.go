package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: go run . <input> <output>")
		os.Exit(1)
	}

	input, output := os.Args[1], os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Fprintln(os.Stderr, "erreur : fichier introuvable :", input)
		} else {
			fmt.Fprintln(os.Stderr, "erreur de lecture :", err)
		}
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		lines[i] = Process(line)
	}
	result := strings.Join(lines, "\n")

	if err = os.WriteFile(output, []byte(result), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "erreur d'écriture :", err)
		os.Exit(1)
	}
}
