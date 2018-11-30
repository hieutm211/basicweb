package login

import (
	"crypto/md5"
	"fmt"
	"github.com/hieutm211/basicweb/config"
	"github.com/hieutm211/basicweb/register/regfunc"
	_ "github.com/lib/pq"
	"net/http"
)

type user struct {
	userid   int
	username string
	name     string
	email    string
	birthday string
	password string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	username := r.FormValue("username")
	password := fmt.Sprintf("%x", md5.Sum([]byte(r.FormValue("password"))))

	if !regfunc.LoginCheck(w, r) {
		return
	}
	fmt.Fprintln(w, username, password)

	db, err := config.InitDB()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer db.Close()

	id := 0
	sqlStmt := `SELECT userid FROM users WHERE username=$1 AND password=$2;`
	err = db.QueryRow(sqlStmt, username, password).Scan(&id)
	if id == 0 {
		fmt.Fprintln(w, "Wrong username or Password")
	} else {
		fmt.Fprintln(w, "Login Succesfully, id = ", id)
	}
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
}
