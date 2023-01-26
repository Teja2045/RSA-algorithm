## Cobra-installation

Use the following command to install the cobra package:

####    go get -u github.com/spf13/cobra/cobra


Import the package in your Go file:

    import "github.com/spf13/cobra"
    
use the following commands to build CLI for your application

#### go build
#### go install

##    Quick Introduction To The Alogithm

The function "GenerateKeys()" generates the encryption and decryption keys, "e" and "d", respectively. It uses the "gcd(pie, e)" function to find the greatest common divisor of "pie" (the product of "a-1" and "b-1") and "e". If the gcd is 1, the keys are generated.

The function "gcd(a int, b int)" uses the Euclidean algorithm to find the greatest common divisor of two integers.

The function "power(a int, b int, mod int)" uses binary exponentiation to raise a number "a" to the power of "b" and returns the result modulo "mod".

The function "publicKeyEncryption1(n int, e int)" takes the product of "a" and "b" and the encryption key "e" as input and encrypts the plaintext message using the RSA encryption algorithm. It returns the encrypted message as a slice of integers.
# 

The function "privateKey(pie int, e int)" calculates the decryption key "d" using the values of "pie" (the product of "a-1" and "b-1") and "e".

The function "decryption1(d int, edata []int, n int)" takes the decryption key "d", the encrypted message "edata" and the product of "a" and "b" as input and decrypts the message using the RSA decryption algorithm. It returns the decrypted message as a byte slice.


## Usage of CLi

The code can be run from the command line using the go run command or cobra root command rsa.
### Generating Keys

To generate the public and private keys needed for encryption and decryption, use the generate command:

#### go run main.go generate
#### rsa generate

    This command will print the values of e and d, which are the public and private keys respectively.
    
    
### Encrypting a Message

To encrypt a message, use the encrypt command followed by the message you want to encrypt:

#### go run main.go encrypt <message>
#### rsa encrypt <message>

    This command will print the encrypted message in the form of an array of integers.

    
### Decrypting a Message

To decrypt an encrypted message, use the decrypt command followed by the array of integers representing the encrypted message:

#### go run main.go decrypt <encrypted message>
#### rsa decrypt <encrypted message>

    This command will print the decrypted message.
    Note

    The code assumes that the two prime numbers used in the encryption process have been assigned in the src package.
