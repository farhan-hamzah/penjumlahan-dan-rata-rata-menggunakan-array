package main
import "fmt"
const NMAX int = 10
type tabInt [NMAX]int
func main(){
	var data tabInt
	var nData, jumlah int
	var rata float64

	baca(&data, &nData)
	cetak(data, nData)
	jumlah = penjumlahan(data, nData)
	fmt.Print(jumlah, " ")
	rata = rerata(data, nData)
	fmt.Print(rata)
}
func baca(A *tabInt, n *int){
	var x int
	for *n < NMAX{
		fmt.Scan(&x)
		if x < 0 {
			x = -x 
		}
		if x == 0{
			break
		}

		(*A)[*n] = x
		*n++
	}

}
func cetak(A tabInt, n int){
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}
func penjumlahan(A tabInt, n int)int{
	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	return sum
}
func rerata(A tabInt, n int)float64{
	if n == 0 {
		return 0.0
	}
	return float64(penjumlahan(A, n)) / float64(n)
}