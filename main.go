package main 

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
		//var card string = "ace of spades"
	//	card := newCard()

	// b := []string{"google", "gmail", "yahoo", "twitter", "facebook"}
	// b = append(b, "new string")

	// for i, list := range b {
	// 	fmt.Println(i, list)
	// 	newCard()
	// }

	// var s = []int{2,4,6,8,10,12}

	// t := []int{1,3,5,7,9}

	// fmt.Println("s = ", s)
	// fmt.Println("s = ", t)

	var arr = [5]string{"google", "apple", "facebook", "yahoo", "twitter"}

	    sli := arr[0:3]

	for i, so := range sli {
		fmt.Println(i, so)
	}

	fmt.Println("array", arr)

	var lang = [5]string{"Go", "javascript", "python", "c++", "java"}

	sil1 := lang[1:5]
	sil2 := lang[:3]
	sil3 := lang[2:]
	sil4 := lang[:]

	fmt.Println("slice1", sil1)
	fmt.Println("slice2", sil2)
	fmt.Println("slice3", sil3)
	fmt.Println("slice4", sil4)
	
	fmt.Println("the sum is ", add(20, 15));
	fmt.Printf("%d - %b \n", 42, 42)
} 