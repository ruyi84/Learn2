package One

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (ms MyString)FindVowels() []rune{
	var vowele []rune
	for _,rune := range ms{
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune =='u'{
			vowele = append(vowele,rune)
		}
	}
	return  vowele
}

func main(){
	name := MyString("Sam Anderson")
	var v VowelsFinder
	v =name
	fmt.Printf("Vowels are %c",v.FindVowels())
}