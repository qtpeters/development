
package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func factorial( i int ) int {
	if i == 0 {
		return 1
	} else {
		return factorial( i - 1 ) * i
	}
}

func main() {
	bytes, _ := ioutil.ReadAll( os.Stdin )
	inpt := strings.TrimSpace( string( bytes ) )
	fnum, _ := strconv.Atoi( inpt )
	fmt.Print( factorial( fnum ) )
}
