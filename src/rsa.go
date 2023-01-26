package src

import (
	"fmt"
)

var (
	a int = 9001
	//prime number 2
	b     int = 23
	data  []byte
	input string = "testing"
	e     int    = 2
	d     int    = 1
	pie   int    = 1
	//edata []int
)

func AssignPrimes(x int, y int) {
	a = x
	b = y
}

func AssignInput(str string) {
	input = str
}

func InputData(str string) []byte {
	data = []byte(str)
	return data
}

func Encrypt(str string) []int {
	GenerateKeys()
	input = str
	return publicKeyEncryption1(a*b, e)
	//fmt.Println(edata)
}

func Decrypt(edatainput []int) []byte {
	// fmt.Println(edatainput)
	GenerateKeys()
	//fmt.Println(d, edata, a, b)
	return decryption1(d, edatainput, a*b)
}

func GenerateKeys() (int, int) {
	// //prime number 1
	// a = 9001
	// //prime number 2
	// b = 23
	// input = "testing"
	pie = (a - 1) * (b - 1)
	e = 2
	for e < pie {
		if gcd(pie, e) == 1 {
			break
		}
		e++
	}

	d = privateKey(pie, e)
	return e, d
}

// eucleadean remainder theorem
func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

// binary exponentiation
func power(a int, b int, mod int) int {
	val := a
	res := 1
	for b > 0 {
		if b%2 == 1 {
			res = res * val
			res = res % mod
		}
		b = b / 2
		val *= val
	}
	return res
}

func publicKeyEncryption1(n int, e int) []int {
	//message here
	data = InputData(input)
	edata := []int{}
	fmt.Println("data is ", data)
	//fmt.Println("data to be encrypted = ", data)
	for i := 0; i < len(data); i++ {
		edata = append(edata, power(int(data[i]), e, n))
	}

	return edata
}

func privateKey(pie int, e int) int {
	k := 2
	for k = 1; k < pie; k++ {
		if (k*pie+1)%e == 0 {
			return (pie*k + 1) / e
		}
	}
	return 2
}

func decryption1(d int, edata []int, n int) []byte {
	var res []byte
	//res := power(edata, d, n)
	for i := 0; i < len(edata); i++ {
		var res1 int
		res1 = 1
		counter := d
		edata1 := edata[i]
		for counter > 0 {
			res1 = res1 * edata1
			res1 = res1 % n
			counter--
		}
		res = append(res, byte(res1))
		//fmt.Print(res1, "|")
	}
	return res
}
