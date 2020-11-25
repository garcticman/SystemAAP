package tasks

func getArraySum(array[] int) (sum int) {
	for _, value := range array {
		sum += value
	}
	return
}

func getPositiveAndNegativeSum(array[] int) (positiveSum, negativeSum int) {
	for _, value := range array {
		if value > 0 {
			positiveSum += value
		} else if value < 0 {
			negativeSum += value
		}
	}
	return
}