package main

import "fmt"

func main() {

	//create map m
	m := map[string]int{
		"Abhay": 24,
		"Pritu": 23,
	}
	fmt.Println(m)

	fmt.Println(m["Abhay"])

	fmt.Println(m["Pritu"])

	//check if key exists in map m
	v, ok := m["Barbadosa"]
	fmt.Println(v)
	fmt.Println(ok)

	//adding element to map m
	m["Jack"] = 33

	//print value if key exists in map m
	if v, ok := m["Jack"]; ok {
		fmt.Println(v)
	}

	//print map m
	fmt.Println("\nMap m after adding")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete key in map m
	delete(m, "Jack")
	fmt.Println("\nMap m after delete")
	fmt.Println(m)

	//check if key exists then delete
	if v, ok := m["Abhay"]; ok {
		fmt.Println("value:", v)
		delete(m, "Abhay")
	}
	fmt.Println("\nMap m")
	fmt.Println(m)

}
