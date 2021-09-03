package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var root string
	if len(os.Args) > 1 {
		root = os.Args[1]
	} else {
		root = "."
	}

	fmt.Println(check(root))
}

type result struct {
	paths []string
}

func check(root string) ([]result, error) {
	d, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	seen := map[string]string{} // normalized path -> real path
	res := []result{}           // make map to unify matches
	for _, f := range d {
		p := filepath.Join(root, f.Name())

		v, exists := seen[p]
		if exists {
			res = append(res, result{paths: []string{v, p}})
		}

		seen[strings.ToLower(p)] = p
	}

	if res == nil {
		res = []result{}
	}

	return res, nil
}
