type OrderCard struct {  
		ID             string  
		Amount         int  
		ModifiersChain []func(int) int  
}

Паттерн Pipeline / Middleware

old - как ненадо

Ты очень круто реализовал математику правил, защитил купон от ухода в минус и
написал конвейер!

Но давай сделаем этот код по-настоящему гибким и архитектурным в парадигме
First-Class Functions (Функции как значения). Здесь есть два важных момента:

1. Архитектура: хранение функций внутри структуры

Посмотри на структуру OrderCard. Сейчас ты хранишь там срез строк []string:

ModifiersChain []string

Из-за этого твой processorBasket жестко привязан к глобальной мапе opMap.
А теперь вспомни тему: функция — это такой же тип данных, как строка или число!

Что, если сделать поле заказа прямо срезом функций:

ModifiersChain []func(int) int

Чувствуешь разницу? Тогда карточка заказа хранит сам исполняемый пайплайн. А
процессор processorBasket становится абсолютно независимым от любых мап и строк:

func processorBasket(order OrderCard) int {
    amount := order.Amount
    for _, mod := range order.ModifiersChain {
        amount = mod(amount)
    }
    return amount
}

Это и есть паттерн Middleware / Pipeline, на котором строятся микросервисы в Go!

