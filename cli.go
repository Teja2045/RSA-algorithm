package main

import (
	"fmt"
	"rsa/src"
	"strconv"

	"github.com/spf13/cobra"
)

// Note: short description tells the inputs for the commands, Long description, about the command functionality
func main() {

	rootCmd := &cobra.Command{
		Use:   "rsa",
		Short: "RSA algorithm",
		Long:  `RSA alogirithm`,
	}

	// primesCmd := &cobra.Command{
	// 	Use:   "primes",
	// 	Short: "| prime number 1 | prime number 2 |",
	// 	Long:  "give 2 prime numbers",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		var a, b int
	// 		aa, _ := strconv.ParseInt(args[0], 10, 64)
	// 		a = int(aa)
	// 		aa, _ = strconv.ParseInt(args[1], 10, 64)
	// 		b = int(aa)
	// 		src.AssignPrimes(a, b)
	// 	},
	// }

	generatekeysCmd := &cobra.Command{
		Use:   "generate",
		Short: "it geneates public and private keys",
		Long:  "no arguments",
		Run: func(cmd *cobra.Command, args []string) {
			e, d := src.GenerateKeys()
			fmt.Println("e value is :", e, ", ", "d value is", d)
		},
	}

	encryptCmd := &cobra.Command{
		Use:   "encrypt",
		Short: "| Message |",
		Long:  "takes 1 input which the a message string",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("not enough arguments provided")
				return
			}
			str := args[0]
			edata := src.Encrypt(str)
			fmt.Println("data is encrypted", edata)
		},
	}

	decryptCmd := &cobra.Command{
		Use:   "decrypt",
		Short: "It decryptes the previous encrypted message ",
		Long:  "Array of encrypted data",
		Run: func(cmd *cobra.Command, args []string) {
			edata := []int{}
			// fmt.Println(args[0])
			// fmt.Println(args[1])
			for i := 0; i < len(args); i++ {
				num, _ := strconv.ParseInt(args[i], 10, 64)
				edata = append(edata, int(num))
			}
			//fmt.Println(edata)
			data := src.Decrypt(edata)
			fmt.Println("data is decrypted: ", string(data))
		},
	}

	//rootCmd.AddCommand(primesCmd)
	rootCmd.AddCommand(generatekeysCmd)
	rootCmd.AddCommand(encryptCmd)
	rootCmd.AddCommand(decryptCmd)
	rootCmd.Execute()

}

//_________________________________________________________________________
