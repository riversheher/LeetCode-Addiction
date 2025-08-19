package threesum

const TARGET int = 0

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	// If there are fewer than 3 numbers, return an empty result
	if len(nums) < 3 {
		return result
	}

	// Create a set of unique numbers, tracking how many times each number appears
	uniqueNums := make(map[int]int)
	for _, num := range nums {
		uniqueNums[num]++
	}

	// Iterate over the unique numbers
	for num1 := range uniqueNums {
		for num2 := range uniqueNums {
			if num1 == num2 {
				//check if we have enough of the same number to form a triplet
				if uniqueNums[num1] < 2 {
					continue
				}
			}

			//find what the required number is.
			required := 0 - (num1 + num2)

			//create a mini map of the numbers we require, tracking how many of them we need
			currNums := make(map[int]int)
			currNums[num1]++
			currNums[num2]++
			currNums[required]++

			// Check if we have enough of the required number
			for _, count := range currNums {
				if count > uniqueNums[required] {
					// If we need more of the required number than we have, skip this pair
					continue
				}
			}

			// Check if the required number is in the set of unique numbers
			if _, exists := uniqueNums[required]; exists {
				triplet := []int{num1, num2, required}

				// Sort the triplet to ensure uniqueness
				if triplet[0] > triplet[1] {
					triplet[0], triplet[1] = triplet[1], triplet[0]
				}
				if triplet[0] > triplet[2] {
					triplet[0], triplet[2] = triplet[2], triplet[0]
				}
				if triplet[1] > triplet[2] {
					triplet[1], triplet[2] = triplet[2], triplet[1]
				}

				// check if the triplet is already in the result
				exists := false

				for _, existingTriplet := range result {
					//we only need to check if two numbers are equal, since the third one is derived from the first two.
					if triplet[0] == existingTriplet[0] && triplet[1] == existingTriplet[1] {
						exists = true
					}
				}

				if !exists {
					result = append(result, triplet)
				}
			}

		}
	}

	return result
}
