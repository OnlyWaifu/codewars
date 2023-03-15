// I love you, a little , a lot, passionately ... not at all
// Who remembers back to their time in the schoolyard, when girls would take a flower and tear its petals, saying each of the following phrases each time a petal was torn:

// "I love you"
// "a little"
// "a lot"
// "passionately"
// "madly"
// "not at all"
// If there are more than 6 petals, you start over with "I love you" for 7 petals, "a little" for 8 petals and so on.

// When the last petal was torn there were cries of excitement, dreams, surging thoughts and emotions.

// Your goal in this kata is to determine which phrase the girls would say at the last petal for a flower of a given number of petals. The number of petals is always greater than 0.

package main

import (
	"fmt"
	"os"
)

func main() {
	var n int

	fmt.Print("Введите колличество лепестков: ")
	fmt.Fscan(os.Stdin, &n)

	fmt.Println(HowMuchILoveYou1(n))
	fmt.Println(HowMuchILoveYou2(n))

}

func HowMuchILoveYou1(i int) string {
	a := ""
	r := []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}
	if i <= 6 {
		i -= 1
		a = r[i]
	} else {
		i -= 1
		i = i % 6
		a = r[i]
	}
	return a
}

func HowMuchILoveYou2(i int) string {
	return []string{"I love you", "a little", "a lot", "passionately", "madly", "not at all"}[(i-1)%6]
}
