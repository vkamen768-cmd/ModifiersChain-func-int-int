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
	ModifiersChain []func(int) int
}

var opMap = map[string]func(int) int{
	"WEEKEND":   discountWeekend,
	"COUPON500": couponFixed,
	"FRAGILE":   feeFragileCargo,
	"HOLIDAY20": discountHoliday,
}

func processorBasket(order OrderCard) int {
	amount := order.Amount
	for _, mod := range order.ModifiersChain {
		if mod == nil {
			continue
		}
		amount = mod(amount)
	}
	return amount
}
func main() {
	orders := []OrderCard{
		{ID: "ORD-01", Amount: 5_000, ModifiersChain: []func(int) int{opMap["WEEKEND"], opMap["COUPON500"]}},
		{ID: "ORD-02", Amount: 600, ModifiersChain: []func(int) int{opMap["COUPON500"], opMap["COUPON500"]}},
		{ID: "ORD-03", Amount: 2_000, ModifiersChain: []func(int) int{opMap["UNKNOWN_PROMO"], opMap["FRAGILE"]}},
	}
	for _, order := range orders {
		fmt.Println(order.ID, order.Amount, processorBasket(order))
	}
}
