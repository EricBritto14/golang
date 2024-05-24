package main

import "fmt"

//Criando uma função, que em python seria o def
func soma(a, b int){
	final := a + b
	fmt.Println("Soma final: ", final)
}

//Função base do programa, como o main do java. Precisa dessa sempre, aí pode criar outras por fora.
func main(){
	//Declaração de variáveis
	pp := 10
	vd := 20

	//Outro modo de declarar variáveis, se quiser declarar o tipo delas
	var a float64
	var b string
	var c bool
	var d int
	const l = 30

	//Colocando valores nas variáveis 
	a = 10.0
	b = "oi"
	c = true
	d = 10

	//Declarando varias do mesmo tipo mas com valores diferentes em uma mesma linha
	var e, f, g int
    e, f, g = 1, 2, 3

	h, i, j := 1, 2, 3 //Assim também funciona, melhorando o exemplo acima
	fmt.Println(a, b, c, d, e, f, g, h, i, j)

	//Também, invez de passar o valor na mão como acima, podemos passar pelo Scanln, que seria com o Scanf do java
	var name string
	fmt.Println("Digite seu nome, garoto: ")
	fmt.Scanln(&name)
	fmt.Println("Olá, " + name)


	//If com else if, igual em java... que em python é elif 
	if pp > vd{
		fmt.Println("pp maior que vd")
	}else if pp < vd{
		fmt.Println("vd maior que pp")
	}else{
		fmt.Println("As duas são iguais")
	}

	//fmt.Println como padrão para printar algo, esse fmt (formato). Também tem o Print normal, mas não pula uma linha que nem o ln
	fmt.Println(pp)
	//Chamando uma função e passando os valores necessários que a função declara
	soma(pp, vd)

	//For baseado no java, com criação da variável dentro do loop.
	for x := 10; x < 20; x++{
		fmt.Println("For imprimindo 10x")
	}

	//For com a variável do lado de fora, e que em teoria seria o While do Go
	mm := 40
	for mm < 50{
		fmt.Println("For que seria o while imprimindo 10x")
		mm++
	}

	// Arrays - Listas
    var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr)
    
    // Slices
    var slice []int = []int{1, 2, 3}
	fmt.Println(slice)

    // Structs - Estrutuas de informações
    type Person struct {
        Name string
        Age  int
    }
    var p Person = Person{Name: "Alice", Age: 30} //Passando os vlaores para a struct
	fmt.Println(p)
    
    // Maps
    var m map[string]int = map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println(m)
}

