// 1-) studentName --> Johnny Depp, grade --> 77, isPassed --> true değişkenlerini
// 3 farklı yöntem ile oluşturup, çıktısını yazdırınız.

package main

import "fmt"

func main() {

	var studentName string = "Johnny Depp"
	var grade int = 77
	var isPassed bool = true

	// 	var studentName = "Johnny Depp"
	// 	var grade = 77
	// 	var isPassed = true

	// 	studentName := "Johnny Depp"
	// 	grade := 77
	// 	isPassed := true

	fmt.Println(studentName)
	fmt.Println(grade)
	fmt.Println(isPassed)
}

// ################################################################################################

// 2-) yukarıda belirtilen değişkenleri tek satır içerisinde tanımlayınız.

// package main

// import "fmt"

// func main() {

// 	var studentName string = "Johnny Depp"
// 	var grade int = 77
// 	var isPassed bool = true

// 	var studentName, grade, isPassed = "Johnny Depp", 77, true

// 	studentName, grade, isPassed := "Johnny Depp", 77, true

// 	fmt.Println(studentName)
// 	fmt.Println(grade)
// 	fmt.Println(isPassed)
// }

// ################################################################################################

// 3-) "Declaration", "Assign", "Initialization", "Initial Value" kavramlarını açıklayınız. (Terminoloji)

/* package main

import "fmt"

func main() {

	var studentName string = "Johnny Depp"

	studentName = "Jason Brown"

	fmt.Println(studentName)

} */

// 4-) "Statically Typed" vs "Dynamically Typed" ifadelerini GO ve Python üzerinden gösteriniz.

// GO'da değişkenin veri tipinin değiştirilmesine izin verilmez. Fakat değişkenin değeri güncellenebilir.

// Python'de değişkenin değeri ve tipi süre gelen satırlarda değiştirilebilir.

// ################################################################################################

// 5-) ":=" vs "=" aradaki farkı gösteriniz, double declaration

// package main

// import "fmt"

// func main() {

// 	/* 	var studentName string = "Johnny Depp"
// 	   	studentName = "Jason Brown" */

// 	studentName := "Johnny Depp"
// 	studentName = "Jason Brown"

// 	fmt.Println(studentName)

// }
