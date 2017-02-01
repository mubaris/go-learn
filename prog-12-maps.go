package main

import "fmt"

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["Ne"] = "Neon"

	flashFamily := map[string]string{
		"Flash":       "Barry Allen",
		"Kid Flash":   "Wally West",
		"Impulse":     "Bart Allen",
		"First Flash": "Jay Garrick",
	}

	elementsDetails := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "Solid",
		},
	}

	fmt.Println(elements["He"])
	fmt.Println(flashFamily["Impulse"])
	fmt.Println(elementsDetails)

	delete(elements, "Ne")
}
