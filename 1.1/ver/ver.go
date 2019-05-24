package ver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//DataModel structure for JSON export
type DataModel struct {
	Name    string //`json:"Name"`
	Release string //`json:"Release"`
	Year    string //`json:"Year"`
}

//Vers sets and prints the version
func Vers() {

	v := ver()
	fmt.Println(v.Name, " version:", v.Release, " ", v.Year)

}

func ver() DataModel {

	f, err := os.Open("./ver/version.dat")
	if err != nil {
		fmt.Println("file create error")

	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	var d1 DataModel
	json.Unmarshal(byteValue, &d1)

	return d1

}
