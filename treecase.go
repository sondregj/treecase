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
		real := filepath.Join(root, f.Name())
		norm := strings.ToLower(real)

		v, exists := seen[norm]
		if exists {
			v.paths = append(v.paths, real)
			continue
		}

		seen[norm] = &result{paths: []string{real}}
	}

	var res []result
	for _, v := range seen {
		if len(v.paths) > 1 {
			res = append(res, *v)
		}
	}

	return res, nil
}
