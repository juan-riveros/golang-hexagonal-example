package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	core "../../core"
)

type (
	// UI represents a web User Interface
	UI struct {
		generator core.SceneGenerator
	}
)

// NewWebUI creates a UI object
func NewWebUI(generator core.SceneGenerator) *UI {
	return &UI{generator}
}

// Run Initializes anything necessary for the UI to function
func (web *UI) Run() {
	fmt.Println("starting webserver")

	web.startServer()

}

func (web *UI) startServer() {
	http.HandleFunc("/", web.displayScene)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/add", web.addValue)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
	fmt.Println(time.Now().UnixNano(), "favicon")
}

func (web *UI) displayScene(w http.ResponseWriter, r *http.Request) {
	scn := web.generator.GenerateScene()

	fmt.Fprintln(w, "Locale:", scn.Locale)
	fmt.Fprintln(w, "Mood:", scn.Mood)
	fmt.Fprintln(w, "Time of Day:", scn.TimeOfDay)
	fmt.Fprintln(w, "Time Period:", scn.TimePeriod)
	fmt.Fprintln(w, "Characters:")
	for _, char := range scn.Characters {
		fmt.Fprintf(w, "\t%+v\n", web.displayCharacter(char))
	}
	fmt.Println(time.Now().UnixNano(), "displayScene")
}

func (web *UI) displayCharacter(character core.Character) string {
	return fmt.Sprintf("%s %s %s", character.AgeGroup, character.Gender, character.Profession)

}

func (web *UI) addValue(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		cat := r.FormValue("cat")
		val := r.FormValue("val")
		web.generator.AddValue(cat, val)
		http.Redirect(w, r, "/", http.StatusFound)
	} else if r.Method == "GET" {
		fmt.Fprintf(w, `<html>
		<script>
			function AddValue() {
				let formdata = new FormData();
				formdata.append('cat', document.getElementById('cat').value)
				formdata.append('val',document.getElementById('val').value)
				fetch('/add', {
					method: "post", 
					body: formdata
				})
			}
			
		</script>
		<form method='post' action='/' target='_self'>
			<p>cat:
			<input id='cat' list='cats' name='cat'>
			<datalist id='cats'>
				<option value='locale'>Locale</option>
				<option value='mood'>Mood</option>
				<option value='timeofday'>Time of Day</option>
				<option value='timeperiod'>Time Period</option>
				<option value='agegroup'>Age Group</option>
				<option value='gender'>Gender</option>
				<option value='profession'>Profession</option>
			</p>
			<p>val:<input id='val' type='text' name='val'></p><input id='addValueBtn' type='submit'>
		</form>
		<script>
		(() => {
			document.getElementById('addValueBtn').addEventListener('click',AddValue)
		})()
		</script>
		</html>`)
	}
}
