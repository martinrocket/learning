package ver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//DataModel structure for JSON export
type DataModel struct {
	Name      string //`json:"Name"`
	Release   string //`json:"Release"`
	Year      string //`json:"Year"`
	StartTime string //`json:"Start Time"`
}

//Vers sets and prints the version
func Vers() {

	v := ver()
	fmt.Println(v.Name, " version:", v.Release, v.Year, " Start Time:", v.StartTime)

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
	t := time.Now()
	var d1 DataModel
	json.Unmarshal(byteValue, &d1)
	d1.StartTime = t.Format("January-01-2006 15:04:05")

	return d1

}
