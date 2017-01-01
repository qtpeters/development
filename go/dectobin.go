
package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

func glbn( target int, lastDouble int, times int ) ( int, int ) {

	var retval int

	nextDouble := lastDouble * 2
	if target == 0 {
		retval = target
	} else if nextDouble > target {
		retval = lastDouble
	} else {
		retval, times = glbn( target, nextDouble, times + 1 )
	}

	return retval, times
}

func getBin( num int, bs []int ) {

	if num > 0 {
		value, times := glbn( num, 1, 0 )
		bslen := len( bs )
		nbs := make( []int, bslen + 1 )
		copy( nbs, bs[:bslen] )
		nbs[ bslen + 1 ] = times
		getBin( num - value, nbs )
	}
}

func main() {

	bytes, _ := ioutil.ReadAll( os.Stdin )
	numStr := strings.TrimSpace( string( bytes ) )
	num, _ := strconv.Atoi( numStr )

	var bs [0]int
	getBin( num, bs[0:] )
	for _, e := range bs {
		fmt.Printf( "Times: %d\n", e )
	}
}
