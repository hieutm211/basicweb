
package register

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"crypto/md5"

	"github.com/hieutm211/basicweb/register/regfunc"
	"github.com/hieutm211/basicweb/config"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	var err error

	username := r.FormValue("username")
	pw := md5.New()
	io.WriteString(pw, r.FormValue("password"))
	password := fmt.Sprintf("%x", pw.Sum(nil))
	fullname := r.FormValue("fullname")
	birthday := r.FormValue("birthday")
	email := r.FormValue("email")

	if !regfunc.Check(w, r) {
		return
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}
	defer db.Close()

	fmt.Fprintln(w, username, password, fullname, birthday, email)

	sqlStmt := `
		INSERT INTO users (username, password, name, birthday, email)
		VALUES($1, $2, $3, $4, $5);`
	_, err = db.Exec(sqlStmt, username, password, fullname, birthday, email)
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
