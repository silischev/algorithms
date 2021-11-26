package arrays

func maxValueLinearComplexity(data []int) int {
	newPeriodDay := 0
	maxPricePeriods := 0
	maxPrice := 0

	for i := 0; i < len(data); i++ {
		price := data[i]

		if i == len(data)-1 {
			if price < data[i-1] {
				maxPricePeriods += price
			} else {
				if price >= maxPrice {
					maxPricePeriods = price * len(data)
				} else {
					maxPricePeriods += price * (newPeriodDay + 1)
				}
			}

			break
		}

		nextPrice := data[i+1]
		if price > nextPrice {
			maxPricePeriods += price * (newPeriodDay + 1)
			newPeriodDay = 0

			if price > maxPrice {
				maxPrice = price
			}
		} else {
			newPeriodDay++
		}
	}

	return maxPricePeriods
}

func maxValueQuadraticComplexity(data []int) int {
	maxPrice := 0
	partVal := 0
	maxPriceDay := 0

	for i := 0; i < len(data); {
		price := data[i]

		if price >= maxPrice {
			maxPrice = price
			maxPriceDay = i
		}

		for j := i + 1; j < len(data); j++ {
			nextDaysPrice := data[j]

			if nextDaysPrice >= maxPrice {
				maxPrice = nextDaysPrice
				maxPriceDay = j
			}
		}

		period := maxPriceDay - i + 1
		partVal += data[maxPriceDay] * period

		if maxPriceDay == len(data)-1 {
			break
		}

		i = maxPriceDay + 1
		maxPrice = 0
		maxPriceDay = 0

		if i == len(data)-1 && maxPrice == 0 {
			partVal += data[i]
			break
		}
	}

	return partVal
}
