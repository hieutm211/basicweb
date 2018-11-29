
package regfunc 

import (
	"fmt"
	"net/http"
)

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z')
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9';
}

func checkUsername(name string) bool {
	if len(name) < 5 || 12 < len(name) {
		return false
	}

	for i := 0; i < len(name); i++ {
		if !isLetter(name[i]) && !isNumber(name[i]) {
			return false;
		}
	}

	return true;
}

func checkFullName(name string) bool {
	if len(name) < 5 || 20 < len(name) {
		return false
	}

	if name[0] == ' ' || name[len(name)-1] == ' ' {
		return false
	}
	for i := 1; i < len(name); i++ {
		if !(isLetter(name[i]) || name[i] == ' ') || (name[i-1] == ' ' && name[i] == ' ') {
			return false
		}
	}
	return true
}

func checkPassword(pw string) bool {
	if len(pw) < 6 || 25 < len(pw) {
		return false
	}
	return true
}

func checkBirthday(date string) bool {
	if len(date) != 10 {
		return false
	}
	return true
}

func Check(w http.ResponseWriter, r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fullname := r.FormValue("fullname")
	birthday := r.FormValue("birthday")

	check:= true

	if !checkUsername(username) {
		fmt.Fprintln(w, "Username must contain 5-12 characters (include Letter, Digit)")
		check = false
	}

	if !checkPassword(password) {
		fmt.Fprintln(w, "Password must contain 6-25 characters")
		check = false
	}

	if !checkFullName(fullname) {
		fmt.Fprintln(w, "Fullname must contain 5-20 characters (include Letter, Space)")
		check = false
	}

	if !checkBirthday(birthday) {
		fmt.Fprintln(w, "Birthday is Invalid")
		check = false
	}

	return check
}
