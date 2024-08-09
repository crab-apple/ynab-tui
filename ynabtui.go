package main

func main() {

	token, err := readAccessToken()
	if err != nil {
		panic(err)
	}

	budgets, err := readBudgets(token)
	if err != nil {
		panic(err)
	}

	println(budgets)
}
