package main

import (
	"fmt"
	"net/http"
	//. "github.com/y0ssar1an/q"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"os"
	"time"
)

var page = `
<!DOCTYPE html>
<html>
	<head>
		<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://unpkg.com/purecss@1.0.0/build/pure-min.css" integrity="sha384-nn4HPE8lTHyVtfCBi5yW9d20FjT8BJwUXyWZT9InLYax14RDjBj46LmSztkmNP9w" crossorigin="anonymous">
<style>
body {
	text-align: center;
}
.checkmark {
  display: inline-block;
}
.checkmark:after {
  /*Add another block-level blank space*/
  content: '';
  display: block;
  /*Make it a small rectangle so the border will create an L-shape*/
  width: 3px;
  height: 6px;
  /*Add a white border on the bottom and left, creating that 'L' */
  border: solid #000;
  border-width: 0 2px 2px 0;
  /*Rotate the L 45 degrees to turn it into a checkmark*/
  transform: rotate(45deg);
}
</style>
	</head>
	<body>
	<table class="pure-table pure-table-bordered" style="width: 100%;">
			<tr>
				<th> Action </th>
				{{ range $day := .Days }}
					<th> {{ $day.Weekday }} </th>
				{{ end }}
			</tr>
			{{ range $action := .Actions }}
			<tr>
				<td> {{ $action.Name }} </td>
				{{ range $j := N 0 6 }}
					<td>
					{{ with $checks := index $action.Date $j }}
						{{ range $i := N 1 $checks }}
							<div class="checkmark"> </div>
						{{ end }}
					{{end}}
					</td>
				{{ end }}
			</tr>
			{{ end}}
		</table>
	</body>

</html>
`

type Action struct {
	id   int
	Name string
	Date []int
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:ambition@(mysql:3306)/model")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t, err := template.New("page").Funcs(template.FuncMap{"N": N}).Parse(page)
	if err != nil {
		fmt.Println(err)
	}

	oneday := time.Now()
	//oneday, err := time.Parse("2006-01-02", "2017-07-02")
	//if err != nil {
	//fmt.Println(err)
	//}
	preweek := last7Days(oneday)

	actions := Actions(db)

	for _, a := range actions {
		for _, day := range preweek {
			a.Date = append(a.Date, Occurrences(db, a.id, day))
		}
	}

	err = t.Execute(w, struct {
		Actions []*Action
		Days    []time.Time
	}{
		actions,
		preweek,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func last7Days(now time.Time) []time.Time {
	return []time.Time{
		now.AddDate(0, 0, -6),
		now.AddDate(0, 0, -5),
		now.AddDate(0, 0, -4),
		now.AddDate(0, 0, -3),
		now.AddDate(0, 0, -2),
		now.AddDate(0, 0, -1),
		now,
	}
}

func Occurrences(db *sql.DB, actionID int, date time.Time) int {
	const format = "2006-01-02"
	const query = `SELECT COUNT(*) FROM occurrences WHERE action_id=? AND Date(datetime)=?`
	resp := db.QueryRow(query, actionID, date.Format(format))
	var count int
	resp.Scan(&count)
	return count
}

func Actions(db *sql.DB) []*Action {
	const query = `SELECT id, action_name from actions`
	resp, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var actions []*Action
	for resp.Next() {
		var a Action
		err := resp.Scan(&a.id, &a.Name)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		actions = append(actions, &a)
	}
	return actions
}

// N allows for iterating over ints
// https://stackoverflow.com/a/22716709
func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		if start > end {
			for i := start; i >= end; i-- {
				stream <- i
			}
		} else {
			for i := start; i <= end; i++ {
				stream <- i
			}
		}
		close(stream)
	}()
	return
}
