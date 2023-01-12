package main

import (
	"fmt"
	//"math"
)

const (
	//iota is a counter variable in go 
	enuconst = iota
	enuconst1 = iota
	enuconst2 = iota
)

const (
	//Compiler infers the values of the next two variables even though we didn't initialize them like we did for first one
	compinferiota = iota
	compinferiota2
	compinferiota3
)

const (
	// iota is scoped to a constant block
	a2 = iota
)

const (
	_ = iota + 5
	catSpecialist
	dogSpecialist 
	snakeSpecialist
)

const (
	_ = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)







// Constants can be shadowed similar to variables 
// Inner scoped constant wins against constant defined with same name at package level

const shadowed int16 = 27

func main(){
	const myConst int = 42
	// Cannot set constants to something that has to be determined at runtime
	//(Compiler has to compute the sin of pi/2 in this case while running)
	// const myConst float64 = math.Sin(1.57)
	fmt.Printf("%v, %T\n",myConst,myConst)
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	const shadowed int = 14
	var bv int = 27
	fmt.Printf("%v\n",a)	
	fmt.Printf("%v\n",b)	
	fmt.Printf("%v\n",c)	
	fmt.Printf("%v\n",d)	
	fmt.Printf("%v,%T\n",shadowed,shadowed)
	// Can perform arithmetic operations on constants similar to variables
	fmt.Printf("%v,%T\n",a+bv,a+bv)
	// Compiler can infer constant type when we don't specify the type explicitly
	const inferType = 42
	var cv int16 = 28
	fmt.Printf("%v,%T\n",a,a)
	// Here it infers the type of the answer same as that of cv cause we didnt specify the type for inferType variable
	// Compiler interprets it as 42+cv 
	fmt.Printf("%v,%T\n",inferType+cv,inferType+cv)
	fmt.Printf("%v,%T\n",enuconst,enuconst)
	fmt.Printf("%v\n",enuconst1)
	fmt.Printf("%v\n",enuconst2)
	fmt.Printf("%v\n",compinferiota)
	fmt.Printf("%v\n",compinferiota2)
	fmt.Printf("%v\n",compinferiota3)
	fmt.Printf("%v\n",a2)

	var specialistType int
	fmt.Printf("%v\n", specialistType == catSpecialist)
	fmt.Printf("%v\n",catSpecialist)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n",roles)
	fmt.Printf("Is Admin? %v\n",isAdmin & roles == isAdmin) 
	fmt.Printf("Is HQ? %v",isHeadquarters & roles == isHeadquarters) 


}