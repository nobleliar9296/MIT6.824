package main
import "fmt"



func main() {

	test := "go go gopher"

	singular(test)

	
}


func singular(test string) {

	table := make(map[string]int)

	i := 0;
	for j := 0; j< len(test); j++  {
		if ( test[j] == ' ' || j == len(test) -1 ) {

			if j == len(test) - 1 {
				j++
			}
			table[test[i:j]]++;
			i = j+1;
		} 
	}

	fmt.Println(table);

}