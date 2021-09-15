package main

import "fmt"

// s := ["q","e","","w","d"]
// i	|func		|i++
//----------------------
// 0	|s[0]="q"	|0++
// 1	|s[1]="e"	|1++
// 2	|skip		|skip
// 2	|s[2]="w"	|2++
// 3	|s[3]="d"	|3++
// 4	|return s[:4]=["q","e","w","d"]

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	strings := []string{"q", "e", "", "w", "d"}
	fmt.Println(nonempty(strings))
	fmt.Println(nonempty2(strings))
}
