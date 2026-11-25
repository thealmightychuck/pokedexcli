package main

import "fmt"

func commandMapb(cfg *config, args ...string) error {

	if cfg.previous == nil || *cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	resp, err := cfg.client.GetPokeAPI(cfg.previous)
	cfg.next = resp.Next
	cfg.previous = resp.Previous

	if err != nil {
		return err
	}

	for _, item := range resp.Results {
		fmt.Println(item.Name)
	}

	//fmt.Printf("Running mapb command\n")
	return nil
}
