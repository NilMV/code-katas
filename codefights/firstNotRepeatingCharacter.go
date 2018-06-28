package codefights

import "fmt"

type RuneState struct {
	fstAppearance int
	numAppearance int
}

var lessPointer *int

type Runes map[rune]*RuneState

func firstNotRepeatingCharacter2(s string) string {
	var repeats [26]int
	for i, c := range s {
		if repeats[c-97] == 0 {
			repeats[c-97] = i + 1
		} else if repeats[c-97] > 0 {
			repeats[c-97] = -1
		}
	}
	r := "_"
	x := len(s) + 1
	for c, i := range repeats {
		if i > 0 && i < x {
			r = string(c + 97)
			x = i
		}
	}
	return r
}

func firstNotRepeatingCharacter(s string) string {
	mapRunes := make(Runes)
	for i, r := range s {
		_, ok := mapRunes[r]
		if ok {
			mapRunes[r].numAppearance += 1
		} else {
			mapRunes[r] = &RuneState{fstAppearance: i, numAppearance: 1}
		}
	}

	for k, _ := range mapRunes {
		if lessPointer == nil && mapRunes[k].numAppearance == 1 {
			lessPointer = &mapRunes[k].fstAppearance

		} else if lessPointer != nil && *lessPointer > mapRunes[k].fstAppearance && mapRunes[k].numAppearance == 1 {
			lessPointer = &mapRunes[k].fstAppearance
		}
	}

	if lessPointer == nil {
		return "_"
	} else {
		return string(s[*lessPointer])
	}

}

func main() {
	fmt.Println(firstNotRepeatingCharacter("bcccccccb"))                                             //_
	fmt.Println(firstNotRepeatingCharacter("abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"))           //d
	fmt.Println(firstNotRepeatingCharacter("xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc"))   //d
	fmt.Println(firstNotRepeatingCharacter("ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof"))  //g
	fmt.Println(firstNotRepeatingCharacter2("bcccccccb"))                                            //_
	fmt.Println(firstNotRepeatingCharacter2("abcdefghijklmnopqrstuvwxyziflskecznslkjfabe"))          //d
	fmt.Println(firstNotRepeatingCharacter2("xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc"))  //d
	fmt.Println(firstNotRepeatingCharacter2("ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof")) //g
}
