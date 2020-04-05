package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func modInv(e int, phi int) int {
	e = e % phi
	for i := 0; i < phi; i++ {
		if e*i%phi == 1 {
			return i
		}
	}
	return 1
}

func modInvN(e int, phi int, n int) int {
	return int(math.Pow(float64(e), float64(phi-1))) % int(n)
}

func findPQ(n int) (int, int) {
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			p := i
			q := n / i
			fmt.Printf("Found p=%d, q=%d\n", p, q)
			return p, q
		}
	}
	return 1, n
}

func decryptFile(d int, n int, filename string) {
	toCrack, fileErr := ioutil.ReadFile(filename)

	if fileErr != nil {
		fmt.Println("Error reading encrypted file to decrypt")
		os.Exit(1)
	}

	alpha := map[int]string{2: "A", 3: "B", 4: "C", 5: "D", 6: "E", 7: "F", 8: "G", 9: "H", 10: "I", 11: "J", 12: "K", 13: "L", 14: "M", 15: "N", 16: "O", 17: "P", 18: "Q", 19: "R", 20: "S", 21: "T", 22: "U", 23: "V", 24: "W", 25: "X", 26: "Y", 27: "Z", 28: " "}

	splits := strings.Fields(string(toCrack))

	var toPrint string

	for _, s := range splits {
		floatS, convErr := strconv.ParseFloat(s, 64)

		if convErr != nil {
			fmt.Println("Cannot convert number in decrypt file\nPlease make sure it only contains base10 numbers in ASCII")
			os.Exit(1)
		}

		dmsg := int(math.Pow(floatS, float64(d))) % int(n)
		toPrint += alpha[dmsg]
	}

	fmt.Println(toPrint)
	filename += "Decrypted"
	writeErr := ioutil.WriteFile(filename, []byte(toPrint), 0644)

	if writeErr != nil {
		fmt.Println("There was a problem writing the file, but we have printed to the screen anyway")
		os.Exit(1)
	}

	fmt.Printf("File %s, written with decrypted data\n", filename)
}

func encryptFile(e int, n int, filename string) {
	toEncrypt, fileErr := ioutil.ReadFile(filename)

	if fileErr != nil {
		fmt.Println("Error reading decrypted file to encrypt")
		os.Exit(1)
	}

	alpha := map[string]int{"A": 2, "B": 3, "C": 4, "D": 5, "E": 6, "F": 7, "G": 8, "H": 9, "I": 10, "J": 11, "K": 12, "L": 13, "M": 14, "N": 15, "O": 16, "P": 17, "Q": 18, "R": 19, "S": 20, "T": 21, "U": 22, "V": 23, "W": 24, "X": 25, "Y": 26, "Z": 27, " ": 28}

	s := string(toEncrypt[:])

	var toPrint string

	for _, i := range s {
		emsg := (int(math.Pow(float64(alpha[strings.ToUpper(string(i))]), float64(e))) % int(n))
		toPrint += strconv.FormatInt(int64(emsg), 10) + " "
	}

	fmt.Println(toPrint)
	filename += "Encrypted"
	writeErr := ioutil.WriteFile(filename, []byte(toPrint), 0644)

	if writeErr != nil {
		fmt.Println("there was a problem writing the file, but we printed to the screen anyway")
		os.Exit(1)
	}

	fmt.Printf("File %s, written with encrypted data\n", filename)
}

func findD(p int, q int, e int) int {
	fmt.Printf("Using %d, %d, %d, to find `d` via the formula: d*e = 1 mod((p-1)(q-1))\n", p, q, e)
	phi := (p - 1) * (q - 1)
	d := modInv(e, phi)
	return d
}

func main() {
	fmt.Println("This program assumes that A->2,B->3,...,Z->27,' '->28\nIf this is not the case, please modify the program accordingly")
	userInput := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to (e)ncrypt or (d)ecrypt?")
	choice, readErr := userInput.ReadString('\n')

	choice = strings.Replace(choice, "\n", "", -1)

	var n int64

	if readErr != nil {
		fmt.Println("Error reading user input")
		os.Exit(1)
	}

	if choice == "e" || choice == "E" {
		fmt.Println("Do you have n (y/n)?")
		nOrPQ, readErr1 := userInput.ReadString('\n')

		if readErr1 != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		nOrPQ = strings.Replace(nOrPQ, "\n", "", -1)

		if nOrPQ == "y" || nOrPQ == "Y" {
			fmt.Println("Please enter n:")
			userN, readErr2 := userInput.ReadString('\n')

			if readErr2 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userN = strings.Replace(userN, "\n", "", -1)
			Tempn, convErr := strconv.ParseInt(userN, 10, 64)

			n = Tempn

			if convErr != nil {
				fmt.Println("Cannot convert `n` to integer")
				os.Exit(1)
			}

			// don't need to find p&q to decrypt, but incase you want to see
			// you can see the factors of n
			// p, q := findPQ(int(n))
			// fmt.Println("p=%d, q=%d", p, q)
		} else {
			fmt.Println("Please enter p: ")
			userP, readErr3 := userInput.ReadString('\n')

			if readErr3 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userP = strings.Replace(userP, "\n", "", -1)
			p, convErr := strconv.ParseInt(userP, 10, 64)

			if convErr != nil {
				fmt.Println("Cannot convert p to integer")
				os.Exit(1)
			}

			userQ, readErr4 := userInput.ReadString('\n')

			if readErr4 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userQ = strings.Replace(userQ, "\n", "", -1)
			q, convErr1 := strconv.ParseInt(userQ, 10, 64)

			if convErr1 != nil {
				fmt.Println("Cannot convert q to integer")
				os.Exit(1)
			}

			n = p * q
		}

		fmt.Println("Please enter e:")
		userE, readErr5 := userInput.ReadString('\n')

		if readErr5 != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		userE = strings.Replace(userE, "\n", "", -1)
		e, convErr2 := strconv.ParseInt(userE, 10, 64)

		if convErr2 != nil {
			fmt.Println("Cannot convert e to integer")
			os.Exit(1)
		}

		fmt.Println("Please enter the name of the file you would like to encrypt (must be in same directory as this program): ")
		filename, filenameReadErr := userInput.ReadString('\n')

		if filenameReadErr != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		filename = strings.Replace(filename, "\n", "", -1)
		encryptFile(int(e), int(n), filename)
	} else if choice == "d" || choice == "D" {
		fmt.Println("Do you have n (y/n)?")
		nOrPQ, readErr1 := userInput.ReadString('\n')

		if readErr1 != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		nOrPQ = strings.Replace(nOrPQ, "\n", "", -1)

		if nOrPQ == "y" || nOrPQ == "Y" {
			fmt.Println("Please enter n:")
			userN, readErr2 := userInput.ReadString('\n')

			if readErr2 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userN = strings.Replace(userN, "\n", "", -1)
			Tempn, convErr := strconv.ParseInt(userN, 10, 64)

			n = Tempn

			if convErr != nil {
				fmt.Println("Cannot convert `n` to integer")
				os.Exit(1)
			}

			// don't need to find p&q to decrypt, but incase you want to see
			// you can see the factors of n
			// p, q := findPQ(int(n))
			// fmt.Println("p=%d, q=%d", p, q)
		} else {
			fmt.Println("Please enter p: ")
			userP, readErr3 := userInput.ReadString('\n')

			if readErr3 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userP = strings.Replace(userP, "\n", "", -1)
			p, convErr := strconv.ParseInt(userP, 10, 64)

			if convErr != nil {
				fmt.Println("Cannot convert p to integer")
				os.Exit(1)
			}

			fmt.Println("Please enter q: ")
			userQ, readErr4 := userInput.ReadString('\n')

			if readErr4 != nil {
				fmt.Println("Error reading user input")
				os.Exit(1)
			}

			userQ = strings.Replace(userQ, "\n", "", -1)
			q, convErr1 := strconv.ParseInt(userQ, 10, 64)

			if convErr1 != nil {
				fmt.Println("Cannot convert q to integer")
				os.Exit(1)
			}

			n = p * q
		}

		fmt.Println("Please enter e:")
		userE, readErr5 := userInput.ReadString('\n')

		if readErr5 != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		userE = strings.Replace(userE, "\n", "", -1)
		e, convErr2 := strconv.ParseInt(userE, 10, 64)

		if convErr2 != nil {
			fmt.Println("Cannot convert e to integer")
			os.Exit(1)
		}

		fmt.Println("Please enter the name of the file you would like to decrypt (must be in same directory as this program): ")
		filename, filenameReadErr := userInput.ReadString('\n')

		if filenameReadErr != nil {
			fmt.Println("Error reading user input")
			os.Exit(1)
		}

		p, q := findPQ(int(n))
		d := findD(int(p), int(q), int(e))
		filename = strings.Replace(filename, "\n", "", -1)
		decryptFile(int(d), int(n), filename)
	} else {
		fmt.Println("Sorry, that was not ['d','D','e','E']")
		os.Exit(1)
	}

}
