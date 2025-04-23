package main

import "fmt"

func commandExplore(cfg *config) error {
	fmt.Println(cfg.args)

	return nil
}
