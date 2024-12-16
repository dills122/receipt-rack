package handlers

import (
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/dills122/receipt-rack/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	retailerPoints := countAlphaNumeric(receipt.Retailer)
	points += retailerPoints

	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if isWholeNumber(total) {
		points += 50
	}

	if isMultipleOfQuarter(total) {
		points += 25
	}

	itemPairsPoints := (len(receipt.Items) / 2) * 5
	points += itemPairsPoints

	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints := int(math.Ceil(price * 0.2))
			points += itemPoints
		}
	}

	parsedDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if parsedDate.Day()%2 != 0 {
		points += 6
	}

	parsedTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if parsedTime.Hour() == 14 || (parsedTime.Hour() == 15 && parsedTime.Minute() < 60) {
		points += 10
	}

	return points
}

func countAlphaNumeric(str string) int {
	count := 0
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			count++
		}
	}
	return count
}

func isWholeNumber(num float64) bool {
	return math.Mod(num, 1) == 0
}

func isMultipleOfQuarter(num float64) bool {
	return math.Mod(num, 0.25) == 0
}
