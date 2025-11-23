package main

import (
	"fmt"
	"math"
)

// O if e else é recomendavel usar apenas quando são poucas expressões...
// 1 ou 2 expressões, acima disso use switch

func main() {
	fruit := "Tv"
	switch fruit {
	case "Tv":
		fmt.Println("Teste")
		fmt.Println(math.MaxFloat64)
		fmt.Println("Tv")
		fmt.Println(math.MaxFloat32)
	case "Banana":
		fmt.Println("Banana")
	// default rodará um código onde nenhum dos cases anteriores foram validados!
	default:
		fmt.Println("Teste")
	}
	// O switch case do go usa break por padrão
	// se desejas que continue verificando mesmo após uma validação
	// ...utilize: fallthrough
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Its a week day")
	case "Sunday":
		fmt.Println("Its sunday!")
	default:
		fmt.Println("Invalid day!")
	}

	number := 17

	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
		// uso de fallthrough: permite verificar a condição abaixo
		fallthrough
	case number == 15 || number == 17:
		fmt.Println("Number is 15 or 17")
		// sem fallthrough aqui: ou seja: a verificação de condição para aqui!
	// default roda um código caso nenhum dos cases sejam realizados!
	default:
		fmt.Println("number is 20 or more!")

	}
	checkType(10)
	checkType(12.1)
	checkType("João Víctor")
	checkType(true)

}

// interface seria uma variavel que desejamos utilizar qualquer tipo
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println(("Its a integer"))
	case float64:
		fmt.Println("Its a float64")
	case float32:
		fmt.Println("Its a float32")
	case string:
		fmt.Println("Its a string")
	default:
		fmt.Println("Unknown type")
	}

}
