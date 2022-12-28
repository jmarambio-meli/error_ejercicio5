package main

import (
	"fmt"
)

type customError5 struct {
	msg string
}

func (c customError5) Error() string {
	return c.msg
}

func NewCustom5(message string) error {
	return &customError5{msg: message}
}

func main() {
	salario, err := SalarioMensual(20, 10000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("El salario correspondiente es %.1f\n", salario)
	}
}

func SalarioMensual(horas, valor float64) (salario float64, err error) {

	if horas < 80 {
		return 0, NewCustom5("Custom: Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	}
	salario = valor * horas
	if salario > 150000 {
		salario = salario * 0.9
	}
	return
}
