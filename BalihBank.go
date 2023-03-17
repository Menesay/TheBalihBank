package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

//
//	The BalihBank Project
//       13.02.2023
//

func createWalletDB() {

	const walletDBname = "username-balance.db"
	const DBname = "username-pass.db"

	// f, err := os.Create(walletDBname)
	// if err != nil {
	// 	log.Fatal("Error while creating Wallet DB", err)
	// }
	// defer f.Close()

	// dbData, err := ioutil.ReadFile(DBname)
	// if err != nil {
	// 	log.Fatal("Error while reading DB", err)
	// }

	// len, err := f.WriteString(string(dbData))
	// len += 1 // make little noise

	walletf, err := os.OpenFile(walletDBname, os.O_APPEND|os.O_CREATE, 0000)
	if err != nil {
		log.Fatal(err)
	}

	//
	//
	//
	dataArray := []string{}
	//////////////////////////////////////////////
	file, err := os.Open(DBname)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		dataArray = append(dataArray, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	//////////////////////////////////////////////
	prefix := ""
	suffix := ""
	for _, val := range dataArray {

		for y, letter := range val {

			if string(letter) == ":" {

				prefix = string(val[:y])
				suffix = string(val[y+1:])

				if _, err := walletf.Write([]byte(prefix)); err != nil {
					log.Fatal(err)
				}
				if _, err := walletf.Write([]byte(":")); err != nil {
					log.Fatal(err)
				}
			}

		}

	}
	println(prefix, " ", suffix)
}

func sleep(millisec float64) {
	time.Sleep(time.Duration(millisec) * time.Millisecond)
}

func isItInDataBase(name_surname, passwd string) bool {
	const DataBase = "username-pass.db"
	dataArray := []string{}

	//////////////////////////////////////////////
	file, err := os.Open(DataBase)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		//fmt.Printf("%d) \"%s\"\n", i, scanner.Text())
		dataArray = append(dataArray, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	/////////////////////////////////////////////

	//fmt.Println(dataArray)

	prefix := ""
	suffix := ""
	for _, val := range dataArray {

		for y, letter := range val {

			//menes:123
			if string(letter) == ":" {
				prefix = string(val[:y])   //menes
				suffix = string(val[y+1:]) //123
			}

			if prefix == name_surname && suffix == passwd {
				return true
			}

		}

	}
	return false
}

// ARCHIVE
func readDataBase() {
	const DataBase = "username-pass.db"
	file, err := os.Open(DataBase)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("%d) \"%s\"\n", i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func writeDataBase(name_surname, passwd string) {

	// fullData := fmt.Sprintf(name_surname, ":", passwd)
	f, err := os.OpenFile("username-pass.db", os.O_APPEND|os.O_CREATE, 0000)
	if err != nil {
		log.Fatal(err)
	}
	//
	if _, err := f.Write([]byte("\n")); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(name_surname)); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(":")); err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(passwd)); err != nil {
		log.Fatal(err)
	}
	//
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

// should false
func isNameInDataBase(name_surname string) bool {

	const DataBase = "username-pass.db"
	dataArray := []string{}

	//////////////////////////////////////////////
	file, err := os.Open(DataBase)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		//fmt.Printf("%d) \"%s\"\n", i, scanner.Text())
		dataArray = append(dataArray, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	/////////////////////////////////////////////

	//fmt.Println(dataArray)

	prefix := ""
	for _, val := range dataArray {
		for y, letter := range val {
			//menes:123
			if string(letter) == ":" {
				prefix = string(val[:y]) //menes
			}
			if prefix == name_surname {
				return true
			}

		}
	}
	return false
}

// ARCHIVE
// should false
func containsSpace(s string) bool {
	for _, val := range s {

		println(string(val))
		if string(val) == " " {
			//
			println("contains space")
			//
			return true
		}
	}
	return false
}

// should true
func trueRegexp(name_surname string) bool {

	res, err := regexp.Match(`^[a-zA-Z]*$`, []byte(name_surname))
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func askPasswd(passwd *string) string {
	print("Password: ")
	fmt.Scanln(passwd)
	return *passwd
}

func askName(name_surname *string) string {
	print("Name-Surname: ")
	fmt.Scanln(name_surname)
	return *name_surname
}

func signup() {

	sleep(100)
	println("\n~~~~~~~~~~~~~~~~~~~~")
	sleep(100)
	print("       Signup        ")
	sleep(100)
	println("\n~~~~~~~~~~~~~~~~~~~~")
	sleep(1000)
	println("Name-Surname Conditions")
	sleep(1000)
	println("> No space")
	sleep(1000)
	println("> Only english letters")
	sleep(1000)
	println("> No number")
	sleep(1000)
	//
	var name_surname string
	askName(&name_surname)

	for isNameInDataBase(name_surname) == true {
		println(name_surname, "is exist. Please enter different one.")
		askName(&name_surname)
	}
	for trueRegexp(name_surname) == false {
		println(name_surname, "isn't suitable.")
		askName(&name_surname)
		for isNameInDataBase(name_surname) == true {
			println(name_surname, "is exist. Please enter different one.")
			askName(&name_surname)
		}
	}
	//

	println("")

	//
	println("Password Conditions")
	sleep(1000)
	println("> Only english letters")
	sleep(1000)
	println("> No space")
	sleep(1000)

	var passwd string
	print("Password: ")
	fmt.Scanln(&passwd)
	println("")
	//

	writeDataBase(name_surname, passwd)
	sleep(2000)
	println("Successfully Sign Upped\n")

	sleep(1000)
	mainInterface()
}

var incorrectLogin int

func login() {

	sleep(100)
	println("\n~~~~~~~~~~~~~~~~~~~~")
	print("       Login        ")
	println("\n~~~~~~~~~~~~~~~~~~~~")
	sleep(100)

	var name_surname string
	var passwd string

	askName(&name_surname)
	for isNameInDataBase(name_surname) == false {
		println(name_surname, "isn't exist. Please enter different one.")
		askName(&name_surname)
	}

	askPasswd(&passwd)

	if isItInDataBase(name_surname, passwd) == true {
		loginInterface(name_surname, passwd)
	}
	if isItInDataBase(name_surname, passwd) == false {

		incorrectLogin += 1
		if incorrectLogin >= 3 {
			sleep(100)
			println("Too many incorrect attempt...")
			sleep(500)
			incorrectLogin = 0
			mainInterface()
		}
		println("Password is incorrect. Please try again.")
		login()
	}
}

func deposit(balance, val float64) (float64, string) {

	if val < 0 {
		return 0, "vallesser"
	}

	return balance + val, "nil"
}
func withdraw(balance, val float64) (float64, string) {

	if val > balance {
		return 0, "valbigger"
	}
	if val < 0 {
		return 0, "vallesser"
	}
	return balance - val, "nil"
}

func roundBalance(balance float64) string {
	return fmt.Sprintf("%3.f", balance)
}

func loginInterface(name_surname, passwd string) {

	var balance float64
	balance = 20.0

	sleep(100)
	println("Login successful")
	sleep(100)
	println("")
	sleep(100)
	println("~~~~~~~~~~~~~~~~~~")
	println("# Logged in as:", name_surname)
	println("# Balance:", roundBalance(balance))
	println("~~~~~~~~~~~~~~~~~~")
	sleep(1000)
	println("1) Deposit")
	sleep(600)
	println("2) Withdraw")
	sleep(600)
	println("3) Log out")
	sleep(600)
	println("~~~~~~~~~~~~~~~~~~")
	sleep(600)

	for {

		var cmd string
		print("Option: ")
		fmt.Scanln(&cmd)

		switch cmd {

		case "1":
			var depoVal float64
			print("Value: ")
			fmt.Scanln(&depoVal)
			tmpbalance, err := deposit(balance, depoVal)
			if err == "nil" {
				balance = tmpbalance
			} else if err == "vallesser" {
				println("Value must not less than 0.")
			}

			println("~~~~~~~~~~~~~~~~~~")
			println("# Balance:", roundBalance(balance))
			println("~~~~~~~~~~~~~~~~~~")

		case "2":

			var withVal float64
			print("Value: ")
			fmt.Scanln(&withVal)

			tmpbalance, err := withdraw(balance, withVal)
			if err == "nil" {
				balance = tmpbalance
			} else if err == "valbigger" {
				println("Value must less than balance.")
			} else if err == "vallesser" {
				println("Value must not less than 0.")
			}

			println("~~~~~~~~~~~~~~~~~~")
			println("# Balance:", roundBalance(balance))
			println("~~~~~~~~~~~~~~~~~~")

		case "3":

			println("Logging out")
			sleep(500)
			mainInterface()

		default:
			println("Enter 1 or 2 or 3.")
		}

	}

}

func mainInterface() {

	println("~~~~~~~~~~~~~~~~~~~~")
	sleep(100)
	println("#                  #")
	sleep(100)
	println("#   THE BALİHBANK  #")
	sleep(100)
	println("#                  #")
	sleep(100)
	println("~~~~~~~~~~~~~~~~~~~~")
	println(" Written in Go 2023 ")
	println("     By Menesay     ")
	println("")
	sleep(500)
	println("1) Sign Up")
	println("2) Login")
	println("3) Exit")
	sleep(100)
	mainlogged := true
	var option string

	for mainlogged {

		print("Option: ")
		fmt.Scanln(&option)

		if option == "1" {
			signup()
			mainlogged = false
		}
		if option == "2" {
			login()
			mainlogged = false
		}
		if option == "3" {
			sleep(500)
			print("\n\r")
			sleep(140)
			print("\rB")
			sleep(140)
			print("\rBa")
			sleep(140)
			print("\rBal")
			sleep(140)
			print("\rBali")
			sleep(140)
			print("\rBalih")
			sleep(140)
			print("\rBalih ")
			sleep(140)
			print("\rBalih w")
			sleep(140)
			print("\rBalih wi")
			sleep(140)
			print("\rBalih wit")
			sleep(140)
			print("\rBalih with")
			sleep(140)
			print("\rBalih with ")
			sleep(140)
			print("\rBalih with y")
			sleep(140)
			print("\rBalih with yo")
			sleep(140)
			print("\rBalih with you")
			sleep(140)
			print("\rBalih with you.")
			sleep(140)
			print("\rBalih with you..")
			sleep(140)
			print("\rBalih with you...")
			sleep(1000)
			os.Exit(0)
		}
	}

}

func main() {

	mainInterface()

}
