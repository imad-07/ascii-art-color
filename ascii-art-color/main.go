package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 && len(args) != 3 && len(args) != 4 {
		fmt.Printf("There is something missing: \nEX: go run . --color=<color> <letters to be colored> \"something\"\nThe rules of this program: If you write something in the banner incorrectly, it will return directly to Standard without leaving an error message")
		return
	}
	l := 1
	baner := ""
	doz := true
	if len(args) == 4 {
		l++
		baner = args[l+1]
	} else if len(args) == 3 && (args[l+1] == "shadow" || args[l+1] == "thinkertoy" || args[l+1] == "standard") {
		baner = args[l+1]
		doz = false
	} else if len(args)-1 == 2 {
		l++
	}

	for i := 0; i < len(args[l]); i++ {
		if args[l][i] < 32 || args[l][i] > 126 {
			fmt.Printf("error in input\n")
			return
		}
	}
	word := split(args[l])
	fileContent, err := os.ReadFile(banner(baner))
	if err != nil {
		fmt.Printf("error in stabdard file")
		return
	}

	lettres := getLettres(fileContent)
	if len(args) == 4 || len(args) == 3 && doz {
		writing(lettres, word, colors(args[0]), args[1])
	} else {
		writing(lettres, word, colors(args[0]), "")
	}
}

// There are three options for writing text and here we know your choice
func banner(arg string) string {
	banner := ""
	if arg == "shadow" {
		banner = "shadow.txt"
	} else if arg == "thinkertoy" {
		banner = "thinkertoy.txt"
	} else if arg == "standard" {
		banner = "standard.txt"
	} else {
		banner = "standard.txt"
	}
	return banner
}

// You write the code with some modifications
func writing(lettres [][]string, word []string, color string, selecletters string) {
	bl := false

	for l := 0; l < len(word); l++ {
		if word[l] == "" {
			continue
		}
		if word[l] == "\n" {
			if len(word)-2 == l {
				fmt.Println()
				continue
			}
			if bl && word[l+1] != "\n" {
				continue
			}
			fmt.Println()
			continue
		}
		str := strings.Split(word[l], " ")

		for i := 1; i < 9; i++ {
			index := 0
			Nbword := 0
			doz := true
			m := 0
			Youcan := 0
			for j := 0; j < len(word[l]); j++ {
				if len(str[index]) > len(selecletters) {
					if m+len(selecletters) <= len(str[index]) {
						if str[index][m:len(selecletters)+m] == selecletters {
							doz = false
							Youcan = len(selecletters)
						}
					}
					m++
					if Nbword == len(str[index]) {
						index++
						Nbword = 0
						doz = true
						m = 0
					} else {
						Nbword++
					}
					if doz == false || len(selecletters) == 0 || Youcan != 0 {
						fmt.Print(color + lettres[word[l][j]-32][i])
						doz = true
						Youcan--
					} else {
						fmt.Print(colors("--color=white") + lettres[word[l][j]-32][i])
					}

				} else {
					if index < len(str) && str[index] == selecletters {
						doz = false
					}

					if Nbword == len(str[index]) {
						index++
						Nbword = 0
						doz = true
					} else {
						Nbword++
					}
					if doz == false || len(selecletters) == 0 {
						fmt.Print(color + lettres[word[l][j]-32][i])
						doz = true
					} else {
						fmt.Print(colors("--color=white") + lettres[word[l][j]-32][i])
					}
				}
			}
			fmt.Println()
		}
		bl = true
	}
}

func split(str string) []string {
	word := ""
	splitedword := []string{}

	for i := 0; i < len(str); i++ {
		if i != len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if word != "" {
				splitedword = append(splitedword, word)
			}
			word = ""

			splitedword = append(splitedword, "\n")
			i++
			continue
		}
		word = word + string(str[i])
	}
	splitedword = append(splitedword, word)
	return splitedword
}

func getLettres(fileContent []byte) [][]string {
	lettres := [][]string{}
	lettre := []string{}
	line := []byte{}
	filtering := ""
	for i := 0; i < len(fileContent); i++ {
		if fileContent[i] != 13 {
			filtering = filtering + string(fileContent[i])
		}
	}
	for i := 0; i < len(filtering); i++ {
		if i != len(filtering)-1 && filtering[i] == '\n' && filtering[i+1] == '\n' {
			lettre = append(lettre, string(line))
			lettres = append(lettres, lettre)
			lettre = nil
			line = nil
			continue
		}
		if filtering[i] == '\n' {
			lettre = append(lettre, string(line))
			line = nil
			continue
		}
		line = append(line, filtering[i])
	}
	lettres = append(lettres, lettre)
	return lettres
}

// ANSI escape codes for text colors
func colors(arg string) string {
	if arg == "--color=reset" || arg == "--color=Reset" {
		arg = "\033[0m"
	} else if arg == "--color=red" || arg == "--color=Red" {
		arg = "\033[31m"
	} else if arg == "--color=green" || arg == "--color=Green" {
		arg = "\033[32m"
	} else if arg == "--color=yellow" || arg == "--color=Yellow" {
		arg = "\033[33m"
	} else if arg == "--color=blue" || arg == "--color=Blue" {
		arg = "\033[34m"
	} else if arg == "--color=purple" || arg == "--color=Purple" {
		arg = "\033[35m"
	} else if arg == "--color=cyan" || arg == "--color=Cyan" {
		arg = "\033[36m"
	} else if arg == "--color=white" || arg == "--color=White" {
		arg = "\033[37m"
	} else if arg == "--color=orange" || arg == "--color=Orange" {
		arg = "\033[38;5;208m" // ANSI escape code for orange
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"The rules of this program: If you write something in the banner incorrectly, it will return directly to Standard without leaving an error message")
		os.Exit(0)
	}
	return arg
}
