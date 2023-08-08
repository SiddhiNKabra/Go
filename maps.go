package main

import "fmt"

func main() {
	studentdata := make(map[string]int)

	studentdata = map[string]int{
		"Siddhi":  31425,
		"Anushka": 31305,
		"Saniya":  31306,
		"Riddhi":  31470,
	}

	sd := studentdata

	fmt.Println(sd)
	studentdata["Charul"] = 31453
	sd1 := studentdata
	fmt.Println(sd1)
	delete(studentdata, "Riddhi")
	fmt.Println(studentdata)

	_, ok := sd["Riddhi"]
	fmt.Println("Presence of Riddhi: ", ok)
	_, test := sd["Anushka"]
	fmt.Println("Presence of Anushka: ", test)
}
