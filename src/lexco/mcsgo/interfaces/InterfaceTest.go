package main

import "fmt"

type Container []interface{}

//--retornara una referencia dle tipo especificado por lo que puedo;deferr alternativa para genrericos
//-container.Put
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

func (c *Container) Get(pos int) interface{} {
	elem := (*c)[pos]
	*c = (*c)[1:]
	return elem
}


func main() {
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)
	elem, corr := intContainer.Get(1).(int)
	if !corr {
		fmt.Println("imposible obtener")
	}
	fmt.Printf("result : %d (%T)\n", elem, elem)
}