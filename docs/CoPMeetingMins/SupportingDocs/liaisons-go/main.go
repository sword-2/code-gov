// main.go. Reads in file liaison.json and prints out the data to a html table, a partial html file (*.l14)

package main

import (
	"fmt"
	"liaisons-go/liaisons"
)

/*
    "encoding/json"
    "io/ioutil"
	"strconv"
	"os"
*/


//	"strings"
//	"reflect" //useful to get type
//	"log"

//Caps to export variable
//type Result struct {
//    Filename string
//    Version string
//	Comments string
//    Agencies []interface{}
//}


func main() {
	fmt.Println("go's main function started v2\n")
	pathL14 := "htmlTable.go.l14"
	liaisons.RemoveOldL14(&pathL14) //if an old .l14 is present, remove it
	pathLiaisons := "liaisons.json"
	liaisons.ReadFile(&pathLiaisons)

	var sortCode int // variable declaration
	sortCode = 2 //1 sort on agency name, 2 sort on legal reference.
	liaisons.SortStruct(&sortCode)
	liaisons.PrintStruct()
	liaisons.WriteL14(&pathL14)

}

/*

htmlTable := "" //a variable to be able to write a file with html table content


//remove old output file if present
	e := os.Remove("htmlTable.go.l14")
		if e != nil {
			fmt.Println("error removing last output file %s\n", e)
	}

contents, _ := ioutil.ReadFile("../../liaisons.json")

var result Result
json.Unmarshal(contents, &result)
//metadata
	htmlTable += "\n<p>source file:" + result.Filename + "</p>"     //fmt.Println(result.Filename)
	htmlTable += "\n<p>version:" + result.Version + "</p>"    //fmt.Println(result.Version)
	htmlTable += "\n<p>comments:" + result.Comments + "</p>"    //fmt.Println(result.Comments)

htmlTable += "\n<table border=1><caption>Agencies</caption>"
htmlTable += "\n\t<tr><th>Agency</th>"
htmlTable += "\n\t<th>Legal Ref</th>"
htmlTable += "\n\t<th>liaisons</th>"
htmlTable += "\n\t<th>repositories</th>"
htmlTable += "\n\t<th>lastModified</th>"
htmlTable += "\n\t<th>offboarding</th>"
htmlTable += "\n</tr>"
//array of agencies to loop thru:     //fmt.Println(result.Agencies)
for _, val := range result.Agencies { //idx1 is just like an array index, 0 - last number
        //fmt.Printf("%d, %s\n", key, val)
		//fmt.Printf("idx1=%d, val type=%s\n", idx1, reflect.TypeOf(val)) //This was for debugging 
		if agency, ok := val.(map[string]interface{}); ok {  //an object for each agency
		//The order of variables can change so load all to output more predictably
			ag := "" //agency name
			ar := "" //agency, legal ref
			li := "" //liaisons
			re := "" //repos
			lm := "" //last modified
			ob := "" //offboarding

        for obj2, v2 := range agency { //object name, object value
			if  obj2 == "agency" { //simple string
				//fmt.Printf("\nagency=%s", v2)
				//htmlTable += "\n" + idx1.(string) + " " + fmt.Sprint(v2)
				ag = "<td>" + agency["agency"].(string) + "</td>" //to have index strconv.Itoa(idx1)
			} else if obj2 == "agencyRef" { //simple string
				// code to be executed if condition-2 is true
				//fmt.Printf("\n\tLegal ref=%s", v2)
				ar = "<td>" + agency["agencyRef"].(string) + "</td>"
			} else if obj2 == "liaisons" {
				//fmt.Printf("\n\tliaisons object, its type=%s\n", reflect.TypeOf(v2))
				lia2 := v2.([]interface{})	// assert a type for for loop
				li = "<td>"
				for idx3, message := range lia2 { //idx is a counter
					//fmt.Printf("\n\tidx3=%d; message=%s\n", idx3, message) //debugging
					nm := ""
					email := ""
					phone := ""
					record := message.(map[string]interface{})  //assert fpr,at
					for idx4, v4 := range record {
						//fmt.Printf("\n\tj=%d, rec2=%s", j, rec2) //debugging
						if idx4 == "name" {
							nm = v4.(string)
						} else if idx4 == "email" {
							email = v4.(string)
						} else if idx4 == "phone" {
							phone = v4.(string)
						} else { //unexpected value
							fmt.Printf("\n\tunexpected object; idx4=(%d), v4=(%s)\n", idx4, v4)
						}
					}
					//fmt.Printf("\n\tliaison %d: name=%s; email=%s; phone=%s", idx3, nm, email, phone)
					li += "\n\t\t<br>" + strconv.Itoa(idx3) + ", " + nm + ", " + email + ", " + phone
				}
				li += "\n\t</td>"
			} else if obj2 == "repos" {
				//fmt.Printf("\n\trepos object, its type=%s\n", reflect.TypeOf(v2))
				repos2 := v2.([]interface{})	// assert a type for for loop
				re = "<td>"
				for idx3, message := range repos2 { //idx is a counter, message has all text for 1 liaison
					//fmt.Printf("\n\tidx3=%d; message=%s\n", idx3, message) //debugging
					nm := ""
					url := ""
					record := message.(map[string]interface{})
					for idx4, v4 := range record { //goes thru each thing in map, name, email, phone
						//fmt.Printf("\n\tj=%d, rec2=%s", j, rec2) //debugging
						if idx4 == "name" {
							nm = v4.(string)
						} else if idx4 == "url" {
							url = v4.(string)
						} else { //unexpected value
							fmt.Printf("\n\tunexpected object; idx4=(%d), v4=(%s)\n", idx4, v4)
						}
					}
					//fmt.Printf("\n\trepo %d: name=%s; url=%s", idx3, nm, url)
					re += "\n\t\t<br>" + strconv.Itoa(idx3) + ", " + nm + ", " + url
				}
				re += "\n\t</td>"
			} else if obj2 == "lastModified" { //simple string
				//fmt.Printf("\n\tlastModified=%s", v2)
				lm = "<td>" + agency["lastModified"].(string) + "</td>"
			} else if obj2 == "offboarding" {
				offboard2 := v2.([]interface{})	// assert a type for for loop
				ob = "<td>"
				for idx3, message := range offboard2 { //idx is a counter, message has all text for 1 liaison
					//fmt.Printf("\n\tidx3=%d; message=%s\n", idx3, message) //debugging
					title := ""
					url := ""
					record := message.(map[string]interface{})
					for idx4, v4 := range record { //goes thru each thing in map, name, email, phone
						//fmt.Printf("\n\tj=%d, rec2=%s", j, rec2) //debugging
						if idx4 == "title" {
							title = v4.(string)
						} else if idx4 == "url" {
							url = v4.(string)
						} else { //unexpected value
							fmt.Printf("\n\tunexpected object; idx4=(%d), v4=(%s)\n", idx4, v4)
						}
					}
					//fmt.Printf("\n\trepo %d: title=%s; url=%s", idx3, title, url)
					ob += "\n\t\t<br>" + strconv.Itoa(idx3) + ", " + title + ", " + url
				}
				ob += "\n\t</td>"
			} else {
				fmt.Printf("\nunknown key (k)/value(v): obj1=%s; v=%s", obj2, v2)
			}
        }
		htmlTable += "\n<tr>" + ag + "\n\t" + ar + "\n\t" + li + "\n\t" + re + "\n\t" + lm + "\n\t" + ob + "\n</tr>"
    } else {
        fmt.Printf("\nval not a map[string]interface{}: %v\n", val)
    }
    }
	htmlTable += "\n</table>"
	//write out contents from variable to .l14 file for verse Luke 14:30
	ioutil.WriteFile("htmlTable.go.l14", []byte(htmlTable), 0644)
*/
