type OrderCard struct {
	ID             string
	Amount         int
	ModifiersChain []func(int) int
}
