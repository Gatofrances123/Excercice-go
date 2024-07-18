package tester

func Totallnconvenience(n int, height []int) int {
	// Create a slice to store the assigned seats
	seats := make([]int, n)

	// Initialize the total inconveniences to 0
	totalInconveniences := 0

	// Iterate through the students
	for i := 0; i < n; i++ {
		// Find the optimal seat for the current student
		optimalSeat := findOptimalSeat(height, seats, i)

		// Assign the seat to the student
		seats[i] = optimalSeat

		// Calculate the inconveniences for the current student
		inconveniences := calculateInconveniences(seats, i)

		// Add the inconveniences to the total
		totalInconveniences += inconveniences
	}

	// If the total inconveniences are 0, return -1
	if totalInconveniences == 0 {
		return -1
	}

	// Otherwise, return the total inconveniences
	return totalInconveniences
}

// Function to find the optimal seat for a student
func findOptimalSeat(height []int, seats []int, i int) int {
	// Find the first available seat
	optimalSeat := 0
	for ; optimalSeat < len(seats); optimalSeat++ {
		if seats[optimalSeat] == 0 {
			break
		}
	}

	// If the current student is taller than the student in the previous seat, shift the current student to the right
	if i > 0 && height[i] > height[i-1] {
		for j := optimalSeat + 1; j < len(seats); j++ {
			if seats[j] == 0 && height[i] > height[i-1] {
				optimalSeat = j
			}
		}
	}

	return optimalSeat
}

// Function to calculate the inconveniences for a student
func calculateInconveniences(seats []int, i int) int {
	// Initialize the inconveniences to 0
	inconveniences := 0

	// Iterate through the seats before the current student's seat
	for j := 0; j < i; j++ {
		// If a seat before the current student's seat is occupied, increment the inconveniences
		if seats[j] != 0 {
			inconveniences++
		}
	}

	return inconveniences
}
