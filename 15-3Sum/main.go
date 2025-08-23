package threesum

const TARGET int = 0

func threeSum(nums []int) [][]int {
	sortedNums := quickSort(nums)

	//To be returned
	result := [][]int{}

	//Iterate through the sorted array using three pointers, smallest, middle, and largest
	for small := 0; small < len(sortedNums)-2; small++ {
		//Skip duplicates for the smallest pointer
		if small > 0 && sortedNums[small] == sortedNums[small-1] {
			continue
		}
		middle, largest := small+1, len(sortedNums)-1

		//While the middle pointer is less than the largest pointer (this skips duplicate processing)
		for middle < largest {
			sum := sortedNums[small] + sortedNums[middle] + sortedNums[largest]
			if sum == TARGET {
				result = append(result, []int{sortedNums[small], sortedNums[middle], sortedNums[largest]})

				//move middle and largest pointers
				middle++
				largest--

				//Skip duplicates for the middle pointer
				for middle < largest && sortedNums[middle] == sortedNums[middle-1] {
					middle++
				}
				//Skip duplicates for the largest pointer
				for middle < largest && sortedNums[largest] == sortedNums[largest+1] {
					largest--
				}
			} else if sum < TARGET {
				//If the sum is less than the target, move the middle pointer up
				middle++
			} else {
				//If the sum is greater than the target, move the largest pointer down
				largest--
			}
		}

	}

	return result
}

// This would be quicker using the sort package, but this is a custom implementation of quicksort
func quickSort(nums []int) []int {
	//Return if len(nums) is less than 2 (only one element)
	if len(nums) < 2 {
		return nums
	}

	//Indices for the quicksort algorithm
	left, right, pivotIdx := 0, len(nums)-1, len(nums)/2
	//Swap pivot into the right
	nums[pivotIdx], nums[right] = nums[right], nums[pivotIdx]

	for i := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	//Swap the pivot back to its correct position
	nums[left], nums[right] = nums[right], nums[left]

	//Recursively sort the left and right partitions
	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}
