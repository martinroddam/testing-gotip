package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/martinroddam/go-tip"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/utilitywarehouse/go-operational/op"
)

type RadioButton struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type PageVariables struct {
	PageTitle         string
	PageRadioButtons  []RadioButton
	PageRadioButtons2 []RadioButton
	Answer            string
}

func main() {
	http.HandleFunc("/", DisplayRadioButtons)
	http.HandleFunc("/selected", UserSelected)
	http.HandleFunc("/fruitselected", UserFruitSelected)
	dm := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "dummy_metric",
		Help: "Dummy counter",
	})
	http.Handle("/__/", op.NewHandler(
		op.NewStatus("My application", "application that does stuff").
			AddOwner("team x", "#team-x").
			SetRevision("7470d3dc24ce7876a9fc53ca7934401273a4017a").
			AddChecker("db check", func(cr *op.CheckResponse) { cr.Healthy("dummy db connection check succeesed") }).
			AddMetrics(dm).
			ReadyUseHealthCheck(),
	),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DisplayRadioButtons(w http.ResponseWriter, r *http.Request) {
	// Display some radio buttons to the user

	Title := "Which do you prefer?"
	MyRadioButtons := []RadioButton{
		RadioButton{"animalselect", "cats", false, false, "Cats"},
		RadioButton{"animalselect", "dogs", false, false, "Dogs"},
	}

	MyRadioButtons2 := []RadioButton{
		RadioButton{"fruitselect", "apple", false, false, "Apple"},
		RadioButton{"fruitselect", "orange", false, false, "Orange"},
	}

	MyPageVariables := PageVariables{
		PageTitle:         Title,
		PageRadioButtons:  MyRadioButtons,
		PageRadioButtons2: MyRadioButtons2,
	}

	t, err := template.ParseFiles("select.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}

}

func UserSelected(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
	// so get the animal which has been selected
	youranimal := r.Form.Get("animalselect")

	Title := "Your preferred animal"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer:    youranimal,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("select.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	gotip.Verify("User Logged In")
}

func UserFruitSelected(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// r.Form is now either
	// map[animalselect:[cats]] OR
	// map[animalselect:[dogs]]
	// so get the animal which has been selected
	yourfruit := r.Form.Get("fruitselect")

	Title := "Your preferred fruit"
	MyPageVariables := PageVariables{
		PageTitle: Title,
		Answer:    yourfruit,
	}

	// generate page by passing page variables into template
	t, err := template.ParseFiles("select.html") //parse the html file homepage.html
	if err != nil {                              // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, MyPageVariables) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                     // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	gotip.Verify("Order Submitted")
}
