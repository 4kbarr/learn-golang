package main

import "fmt"

func main() {
	var name string

	var place = "Jakarta" //var tidaklah wajib

	name = "Thomas Mueller"
	fmt.Println(name)
	fmt.Println(place)

	//nilai variable name dirubah
	name = "Kevin de Bruyne"
	fmt.Println(name)

	//menyederhanakan deklarasi variable
	lastname := "myLastName"
	fmt.Println(lastname)

	//membuat variable secara banyak
	var (
		namadpn  = "Xabi"
		namablkg = "Alonso"
	)

	fmt.Println(namadpn)
	fmt.Println(namablkg)
}
