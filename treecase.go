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

	seen := map[string]*result{} // normalized path -> real paths
	for _, f := range d {
		p := filepath.Join(root, f.Name())

		v, exists := seen[p]
		if exists {
			v.paths = append(v.paths, p)
			continue
		}

		seen[strings.ToLower(p)] = &result{paths: []string{p}}
	}

	var res []result
	for _, v := range seen {
		if len(v.paths) > 1 {
			res = append(res, *v)
		}
	}

	return res, nil
}
