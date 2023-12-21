package buytwochocolates

import "math"

func buyChoco(prices []int, money int) int {

	minPrice := math.MaxInt
	secondMinPrice := math.MaxInt

	for _, price := range prices {
		if price < minPrice {
			secondMinPrice = minPrice
			minPrice = price
		} else if price < secondMinPrice {
			secondMinPrice = price
		}
	}

	if secondMinPrice+minPrice <= money {
		return money - (secondMinPrice + minPrice)
	}

	return money

}
