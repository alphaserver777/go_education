package main

import "fmt"

type Builder interface { // интерфейс
	Build()
}

type Person struct { // структура
	Name string
	Age  int
}

type WorkExperience struct { // структура
	Name string
	Age  int
}

func (p Person) printName() { // метод структуры Person
	fmt.Println(p.Name)
}

type WoodBuilder struct { // структура
	Person
	//Name string
	//WorkExperience
}

/*func (wb WoodBuilder) printName() {
	fmt.Println(wb.Name)
}*/

func (wb WoodBuilder) Build() { // метод структуры WoodBuilder
	fmt.Println("Строю дом из дерева")
}

type BrickBuilder struct {
	Person
}

func (bb BrickBuilder) Build() {
	fmt.Println("Строю из кирпича")
}

type Building struct { // структура
	Builder
	Name string
}

func main() {
	//explanation()
	usecase()
}

/*func explanation() {
	//builder := WoodBuilder{Person{Name: "Вася", Age: 30}}
	//builder := WoodBuilder{Person{Name: "Вася", Age: 30}, "Боб"}
	builder := WoodBuilder{
		Person{Name: "Вася", Age: 30},
		"Боб",
		WorkExperience{Name: "Таксист", Age: 3}}
	fmt.Printf("Type: %T Value: %#v\n", builder, builder)

	// shorthands
	fmt.Println(builder.Person.Age)
	fmt.Println(builder.WorkExperience.Age)

	//shadowing
	fmt.Println(builder.Name)
	fmt.Println(builder.Person.Name)

	builder.printName()
}*/

// usecase demonstrates the usage of the Building struct and its associated builders.
// It creates two instances of Building, one using a WoodBuilder and the other using a BrickBuilder,
// and then calls the Build method on each instance to print the construction process.
func usecase() {
	woodenBuilding := Building{
		Builder: WoodBuilder{Person{
			Name: "Вася",
			Age:  40,
		}},
		Name: "Деревянная изба",
	}
	woodenBuilding.Build()

	brickBuilding := Building{
		Builder: BrickBuilder{
			Person{
				Name: "Петя",
				Age:  30,
			},
		},
		Name: "Кирпичный дом",
	}
	brickBuilding.Build()
}
