package main

// Need to show solution

func handleOperations(id string) {
	operations := getOperations(id)
	if operations != nil {
		// handling...
	}
}

func getOperations(id string) []float32 {
	opearations := make([]float32, 0)
	if id == "" {
		return opearations
	}

	// adding operations...
	return opearations
}
