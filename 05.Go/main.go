package main

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() {
	fmt.Printf("Hola, mi nombre es %s y tengo %d años.\n", p.Nombre, p.Edad)
}

func main() {
	println("Hello, World!")
	a := 10
	var b int = 20
	var c int8 = 10
	var d int16 = 20
	var e int32 = 30
	var f int64 = 40

	s := "hola mundo"

	if a < b {
		println("a es menor que b")
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	res, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Resultado:", res)
	}

}

func sum(x int, y int) int {
	return x + y
}

func duplicar(x int) int {
	return x * 2
}

func dosNumeros(x int, y int) (int, int) {
	return x + y, x * y
}

func dividir(x int, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return x / y, nil
}
