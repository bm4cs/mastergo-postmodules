package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		input float64
		result float64
		fromUnit string
		toUnit string
	)

	flag.StringVar(&fromUnit, "from", "", "The unit to convert from")
	flag.StringVar(&toUnit, "to", "", "The unit to convert to")
	flag.Parse()

	if toUnit == "" {
		fmt.Println("-to unit not provided")
		return
	}

	if fromUnit == "" {
		fmt.Println("-from unit not provided")
		return
	}

	if len(flag.Args()) == 0 {
		fmt.Println("No value argument provided")
		return
	}

	if _, err := fmt.Sscanf(flag.Args()[0], "%f", &input); err != nil {
		fmt.Println("Could not parse argument as floating point value")
		return
	}

	switch {
	case fromUnit == "meters" && toUnit == "feet":
		result = input * 3.281
	case fromUnit == "feet" && toUnit == "meters":
		result = input / 3.281
	case fromUnit == "meters" && toUnit == "inches":
		result = input * 39.37
	case fromUnit == "inches" && toUnit == "meters":
		result = input / 39.37
	default:
		fmt.Println("Sorry, the units you specified are not supported")
		return
	}

	fmt.Println("Result:", result)
}
