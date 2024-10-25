package main

import (
	"fmt"
)

func mapF(cfg *config) error {
	locationsResp, err := cfg.client.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func mapB(cfg *config) error {
	locationsResp, err := cfg.client.ListLocations(cfg.previousURL)
	if err != nil {
		return err
	}
	cfg.nextURL = locationsResp.Next
	cfg.previousURL = locationsResp.Previous

	for _, result := range locationsResp.Results {
		fmt.Println(result.Name)
	}
	return nil
}
