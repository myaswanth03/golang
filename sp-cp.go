package main
import "fmt"
func main() {
	var sellingprice, costprice float64

	fmt.Println("enter the costprice : ")
	fmt.Scanln(&costprice)

	fmt.Println("enter the sellingprice : ")
	fmt.Scanln(&sellingprice)

	if sellingprice > costprice {
		profit := sellingprice - costprice
		fmt.Printf("Profit: %.2f\n", profit)
	} else if costprice > sellingprice {
		loss := costprice - sellingprice
		fmt.Printf("Loss: %.2f\n", loss)
	} else {
		fmt.Println("No Profit, No Loss.")
	}

}
