package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		output = flag.String("o", "", "output format")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()

	var root string
	if len(flag.Args()) > 1 {
		root = flag.Args()[1]
	} else {
		root = "."
	}

	res, err := check(root)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch *output {
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		err := enc.Encode(res)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	default:
		for _, v := range res {
			for _, p := range v.Paths {
				fmt.Printf("%s\n", p)
			}
			fmt.Println()
		}
	}
}

type result struct {
	Paths []string `json:"paths"`
}

func check(root string) ([]result, error) {
	d, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var res []result
	seen := map[string]*result{} // normalized path -> real paths
	for _, f := range d {
		real := filepath.Join(root, f.Name())
		norm := strings.ToLower(real)

		if f.IsDir() {
			r, err := check(real)
			if err != nil {
				return nil, err
			}
			res = append(res, r...)
		}

		v, exists := seen[norm]
		if exists {
			v.Paths = append(v.Paths, real)
			continue
		}

		seen[norm] = &result{Paths: []string{real}}
	}

	for _, v := range seen {
		if len(v.Paths) > 1 {
			res = append(res, *v)
		}
	}

	return res, nil
}
