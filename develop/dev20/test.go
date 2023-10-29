package main

import (
	"fmt"
	"strings"
)

/*

=== Задача ===

Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».

*/

func main() {
	sequence := "snow dog sun"

	reverseSequence := reverseSeq(sequence)
	fmt.Printf("%v - %v\n", sequence, reverseSequence)

}

func reverseSeq(sequence string) string {
	args := strings.Split(sequence, " ")
	reversedSeq := make([]string, 0, len(args))

	for i := len(args) - 1; i >= 0; i-- {
		reversedSeq = append(reversedSeq, args[i])
	}

	return strings.Join(reversedSeq, " ")
}
