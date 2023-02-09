package main

import "fmt"
func tinhTong(n int) int {
	Tong := 0
	i := 1
	for i <= n {
		Tong = Tong + i
		i++
	}
	return Tong
}
func main() {
	var n int
	fmt.Printf("Nhap n: \n")
	fmt.Scanf("%d",&n)
	Tong := tinhTong(n)
	fmt.Printf("Tong can tinh la: %d",Tong)
}