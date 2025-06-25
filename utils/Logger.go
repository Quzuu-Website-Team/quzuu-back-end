package utils

import (
	"fmt"
)

func LogError(errorLogged error) {
	fmt.Println("There is an error!")
	fmt.Println("Error Log :", errorLogged)
	
}
