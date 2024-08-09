// functions for reading, sorting, and writing to an .l14 file

package liaisons

import (
	"fmt"
    "encoding/json"
	"os"
	"sort"
	 "strconv"
)

func check(e error) {
	if e != nil {
        panic(e)
    } //else {
	//	fmt.Println("\nno error for e")
	//}
}

func RemoveOldL14(path *string) { //if an old .l14 is present, remove it
	if _, err := os.Stat(*path); err == nil {
		check(err)
		fmt.Println("\nRemoveOldL14(): old .l14 exists planning to remove it")
		e := os.Remove(*path)
		check(e)
	}
}

func ReadFile(path *string) {
	fmt.Println("\nIn package liaisons, func ReadFile() was called to open JSON file: ", *path)

	//path relative to module root, not this source file
	contents, err := os.ReadFile(*path)
	check(err)

    err2 := json.Unmarshal([]byte(contents), &r)
	check(err2) //check no eror
}


func SortStruct(sortCode *int) {
	fmt.Println("\nfunc SortStruct() called")

	// Sort by one way by removing comments

	if *sortCode == 1 { //Agency name
		fmt.Println("SortStruct() sortCode == 1 for Agency name")
		sort.Slice(r.Agencies, func(i, j int) bool {
			return r.Agencies[i].Agency < r.Agencies[j].Agency
		})
	} else if *sortCode == 2 { //AgencyRef aka legal reference
		fmt.Println("SortStruct() sortCode == 2 for legal reference")
		sort.Slice(r.Agencies, func(i, j int) bool {
			return r.Agencies[i].AgencyRef < r.Agencies[j].AgencyRef
		})
	} else {
		//fmt.Println("SortStruct() received unknown sortCode.")
		panic("SortStruct() received unknown sortCode.")
	}




}

func PrintStruct() {
	fmt.Println("\nfunc PrintStruct() called")

	htmlTable = "" //clear global variable

	//metadata
		htmlTable += "\n<p>source file:" + r.Filename + "</p>"     //fmt.Printf("\nr.Filename=%s", r.Filename)
		htmlTable += "\n<p>version:" + r.Version + "</p>"    //fmt.Printf("\nr.Version=%s", r.Version)
		htmlTable += "\n<p>comments:" + r.Comments + "</p>"    //fmt.Printf("\nr.Comments=%s", r.Comments)

	//table setup
		htmlTable += "\n<table border=1><caption>Agencies and liaisons as processed by Go language (Google)</caption>"
		htmlTable += "\n\t<tr><th>Agency</th>"
		htmlTable += "\n\t<th>Legal Ref</th>"
		htmlTable += "\n\t<th>liaisons</th>"
		htmlTable += "\n\t<th>repositories</th>"
		htmlTable += "\n\t<th>lastModified</th>"
		htmlTable += "\n\t<th>offboarding</th>"
		htmlTable += "\n</tr>"

	//fmt.Printf("\nr.Agencies=%s", r.Agencies)

	//loop thru agencies one at a time
	for i := 0; i < len(r.Agencies); i++ {
		htmlTable += "\n<tr>"
		//fmt.Printf("\na.Agency=%s; LegalRef=%s", r.Agencies[i].Agency, r.Agencies[i].AgencyRef)
		htmlTable += "\n\t<td>" + r.Agencies[i].Agency + "</td>" //name
		htmlTable += "\n\t<td>" + r.Agencies[i].AgencyRef + "</td>" //legal ref

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Liaisons); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Liaisons[j].Name + "; " +  r.Agencies[i].Liaisons[j].Email + "; " +  r.Agencies[i].Liaisons[j].Phone
				if j != len(r.Agencies[i].Liaisons) { htmlTable += "<br>" }
			}
			htmlTable += "\n\t</td>"

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Repos); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Repos[j].Name + "; " +  r.Agencies[i].Repos[j].Url
				if j != len(r.Agencies[i].Repos) { htmlTable += "<br>" }
			}
		htmlTable += "\n\t</td>"

		htmlTable += "\n\t<td>" + r.Agencies[i].LastModified + "</td>"

		htmlTable += "\n\t<td>"
			for j := 0; j < len(r.Agencies[i].Offboarding); j++ {
				htmlTable += "\n\t" + strconv.Itoa(j) + ": " + r.Agencies[i].Offboarding[j].Title + "; " +  r.Agencies[i].Offboarding[j].Url
				if j != len(r.Agencies[i].Offboarding) { htmlTable += "<br>" }
			}
		htmlTable += "\n\t</td>"

		htmlTable += "\n</tr>"
	}
		htmlTable += "\n</table>"

}

func WriteL14(path *string) { //write data to a .l14 file so a web page can load it
	err := os.WriteFile(*path, []byte(htmlTable), 0644)
	check(err)
}
