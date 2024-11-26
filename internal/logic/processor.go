package logic

import (
	"math"
	"strconv"
	"strings"
	"time"

	"receipt-processor/internal/models"
)

func CalculatePoints(receipt *models.Receipt) int {
	points := 0

	// 1. One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// 2. 50 points if the total is a round dollar amount with no cents
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 3. 25 points if the total is a multiple of 0.25
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items on the receipt
	points += (len(receipt.Items) / 2) * 5

	// 5. Points for items with descriptions whose trimmed length is a multiple of 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			itemPrice, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(itemPrice * 0.2))
		}
	}

	// 6. 6 points if the day in the purchase date is odd
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if date.Day()%2 != 0 {
		points += 6
	}

	// 7. 10 points if the time of purchase is between 2:00pm and 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 {
		points += 10
	}

	return points
}

func countAlphanumeric(input string) int {
	count := 0
	for _, c := range input {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			count++
		}
	}
	return count
}
