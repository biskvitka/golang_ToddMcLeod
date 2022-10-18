package main
import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"strings"
)

func main() {

	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

}