package main
import "fmt"

func PublicFunc() {
	fmt.Println("Ашық функция (PublicFunc)")
}
func privateFunc() {
	fmt.Println("Жеке функция (privateFunc)")
}
func main() {
	PublicFunc()  
	privateFunc() 
}
