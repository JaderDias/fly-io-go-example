package main

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

//go:embed templates/*
var resources embed.FS

var t = template.Must(template.ParseFS(resources, "templates/*"))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	{

		db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		createStatement := `CREATE TABLE accounts (
			user_id serial PRIMARY KEY,
			username VARCHAR ( 50 ) UNIQUE NOT NULL,
			password VARCHAR ( 50 ) NOT NULL,
			email VARCHAR ( 255 ) UNIQUE NOT NULL,
			created_on TIMESTAMP NOT NULL,
				last_login TIMESTAMP 
		)`

		_, err = db.Exec(createStatement)
		if err != nil {
			fmt.Fprintf(os.Stderr, "createStatement failed: %v\n", err)
		}

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := map[string]string{
			"Region": os.Getenv("FLY_REGION"),
		}
		if r.Method != http.MethodPost {
			t.ExecuteTemplate(w, "index.html.tmpl", data)
			return
		}

		details := map[string]string{
			"username": r.FormValue("username"),
			"password": r.FormValue("password"),
			"email":    r.FormValue("email"),
		}

		db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer db.Close()

		sqlStatement := `insert into Accounts (username, password, email, created_on) values ($1, $2, $3, NOW())`
		_, err = db.Exec(
			sqlStatement,
			details["username"],
			details["password"],
			details["email"],
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		}
		t.ExecuteTemplate(w, "index.html.tmpl", data)
	})

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
