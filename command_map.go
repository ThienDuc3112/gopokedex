package main

import "fmt"

func commandMap(config *Config, _ []string) error {
	locationList, err := config.PokeClient.ListLocation(config.Next)
	if err != nil {
		return fmt.Errorf("error getting map: %v", err)
	}
	config.Next = locationList.Next
	config.Previous = locationList.Previous

	for _, v := range locationList.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(config *Config, _ []string) error {
	if config.Previous == nil {
		return fmt.Errorf("cannot go back, at first page")
	}
	locationList, err := config.PokeClient.ListLocation(config.Previous)
	if err != nil {
		return fmt.Errorf("error getting map: %v", err)
	}
	config.Next = locationList.Next
	config.Previous = locationList.Previous

	for _, v := range locationList.Results {
		fmt.Println(v.Name)
	}

	return nil
}
