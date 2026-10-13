type OrderCard struct {</br>
	ID             string
	Amount         int
	ModifiersChain []func(int) int
}
