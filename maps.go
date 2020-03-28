package main

import "fmt"

func main() {
	deleteValueFromMap("iOS")
}

func printPersonSalary() {
	var personSalary = make(map[string]int)
	personSalary["Steve"] = 12000
	personSalary["Harley"] = 2000
	personSalary["Tom"] = 70000

	fmt.Println("Person salary map contents: ", personSalary)
}

func printStudentsID() {
	var studentsID = map[int]string{
		1233331291233123: "Arifin Firdaus",
		1119123812331231: "Steve Harley",
		9998882221241244: "Tom Ausley",
	}

	fmt.Println("Students ID map contents: ", studentsID)
}

func printOneItemFromMap() {
	var platformProgrammingLanguage = map[string]string{
		"iOS":     "Swift",
		"Android": "Kotlin, Java",
		"Backend": "Go",
	}

	fmt.Println("Programming Language: ", platformProgrammingLanguage["Android"])
}

func checkKeyFromMap() {
	var learnedUIArchitecture = make(map[string]bool)
	learnedUIArchitecture["MVC"] = true
	learnedUIArchitecture["MVVM"] = true
	learnedUIArchitecture["MVP"] = true
	learnedUIArchitecture["VIPER"] = false
	learnedUIArchitecture["MVI"] = false

	var searchKey = "MVC"
	var value, isExist = learnedUIArchitecture[searchKey]

	if isExist {
		fmt.Println("Search key :", searchKey, "value :", value)
		return
	}
	fmt.Println("search key", searchKey, " not found")
}

func printAllValues() {
	var platformProgrammingLanguage = map[string]string{
		"iOS":     "Swift",
		"Android": "Kotlin, Java",
		"Backend": "Go",
	}

	for key, value := range platformProgrammingLanguage {
		fmt.Printf("[%s] = [%s] \n", key, value)
	}
}

func deleteValueFromMap(key string) {
	var platformProgrammingLanguage = map[string]string{
		"iOS":     "Swift",
		"Android": "Kotlin, Java",
		"Backend": "Go",
	}
	fmt.Println("length before delete: ", len(platformProgrammingLanguage))

	delete(platformProgrammingLanguage, key)

	fmt.Println("length before delete: ", len(platformProgrammingLanguage))
}
