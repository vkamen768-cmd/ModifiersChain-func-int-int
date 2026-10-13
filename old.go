package main

import "fmt"

func discountWeekend(amount int) int {
	return amount - (amount * 10 / 100)
}
func couponFixed(amount int) int {
	finalAmount := amount - 500
	if finalAmount < 0 {
		return 0
	}
	return finalAmount
}
func feeFragileCargo(amount int) int {
	return amount + 300
}
func discountHoliday(amount int) int {
	return amount - (amount * 20 / 100)
}

type OrderCard struct {
	ID             string
	Amount         int
	ModifiersChain []string
}

var opMap = map[string]func(int) int{
	"WEEKEND":   discountWeekend,
	"COUPON500": couponFixed,
	"FRAGILE":   feeFragileCargo,
	"HOLIDAY20": discountHoliday,
}

func processorBasket(order OrderCard) int {
	for _, mod := range order.ModifiersChain {
		if op, ok := opMap[mod]; ok {
			order.Amount = op(order.Amount)
		} else {
			op = opMap["FRAGILE"]
			order.Amount = op(order.Amount)
		}
	}
	return order.Amount
}

func main() {
	orders := []OrderCard{{ID: "ORD-01", Amount: 5_000, ModifiersChain: []string{"WEEKEND", "COUPON500"}},
		{ID: "ORD-02", Amount: 600, ModifiersChain: []string{"COUPON500", "COUPON500"}},
		{ID: "ORD-03", Amount: 2_000, ModifiersChain: []string{"UNKNOWN_PROMO"}},
	}
	for _, order := range orders {
		fmt.Println(order.ID, order.Amount, processorBasket(order))
	}
}
