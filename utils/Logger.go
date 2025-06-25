package utils

import (
	"fmt"
	"log"
	"godp.abdanhafidz.com/config"
)

func LogError(errorLogged error) {
	fmt.Println("There is an error!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Error Log :", errorLogged)
}
