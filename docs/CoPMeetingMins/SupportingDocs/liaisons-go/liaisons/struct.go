package liaisons


//global variables

//struct for marshalling liaisons.json, Caps to export variable. The word marshalling is a software concept, probably not the more common word definition for police type marshall.
type Result struct {
    Filename string
    Version string
	Comments string
    Agencies []struct {
		Agency string
		AgencyRef string
		LastModified string
		Liaisons[]struct {
			Email string
			Name string
			Phone string
		}
		Offboarding[]struct {
			Title string
			Url string
		}
		Repos[]struct {
			Name string
			Url string
		}
	}
}

var r Result //global variable for r(R)esult.


var htmlTable string //same as r above, but string format. string format is easier to write out to file since no need to loop through all the variables in r.
