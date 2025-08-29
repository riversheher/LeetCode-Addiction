package perfectsquares

func numSquares(n int) int {
	coins := make([]int, 0)

	for i := 1; i*i <= n; i++ {
		coins = append(coins, i*i)
	}

	minCoins := make([]int, n+1)
	minCoins[0] = 0
	for i := 1; i <= n; i++ {
		min := n + 1
		for _, coin := range coins {
			//break if coin is greater than the amount
			if coin > i {
				break
			}
			result := minCoins[i-coin] + 1
			if result < min {
				min = result
			}
		}

	}

	return minCoins[n]
}
