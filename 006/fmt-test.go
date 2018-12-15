package main
import ("fmt")
var y = 23
func main(){
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	y=911
	//fmt.Printf("%#x\n", y)
	fmt.Printf("%b\t%x\t%#x\t\n", y, y, y)

	str := fmt.Sprintf("%b\t%x\t%#x\t\n", y, y, y)
	fmt.Println(str)

	fmt.Printf("%v\n", y)
}