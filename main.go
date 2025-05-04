package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	fmt.Println("Приложение паролей")
MenuLabel:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break MenuLabel
		}
	}
}

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или LOGIN")
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)
}

func getMenu() (variant int) {
	fmt.Println("Выберите вариант:")
	fmt.Println("1: Создать аккаунт")
	fmt.Println("2: Найти аккаунт")
	fmt.Println("3: Удалить аккаунт")
	fmt.Println("4: Выход")
	fmt.Scan(&variant)
	return
}

func findAccount() {

}

func deleteAccount() {

}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
