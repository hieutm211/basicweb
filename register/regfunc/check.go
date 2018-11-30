package regfunc

import (
	"fmt"
	"net/http"
)

const (
	UsernameRule = "Username must contain 5-12 characters (include Letter, Digit)"
	PasswordRule = "Password must contain 6-25 characters"
	FullnameRule = "Fullname must contain 5-20 characters (include Letter, Space)"
	BirthdayRule = "Birthday is Invalid"
)

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z' || 'A' <= c && c <= 'Z')
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func CheckUsername(name string) bool {
	if len(name) < 5 || 12 < len(name) {
		return false
	}

	for i := 0; i < len(name); i++ {
		if !isLetter(name[i]) && !isNumber(name[i]) {
			return false
		}
	}

	return true
}

func CheckFullname(name string) bool {
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

func CheckPassword(pw string) bool {
	if len(pw) < 6 || 25 < len(pw) {
		return false
	}
	return true
}

func CheckBirthday(date string) bool {
	if len(date) != 10 {
		return false
	}
	return true
}

func RegisterCheck(w http.ResponseWriter, r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fullname := r.FormValue("fullname")
	birthday := r.FormValue("birthday")

	check := true

	if !CheckUsername(username) {
		fmt.Fprintln(w, UsernameRule)
		check = false
	}

	if !CheckPassword(password) {
		fmt.Fprintln(w, PasswordRule)
		check = false
	}

	if !CheckFullname(fullname) {
		fmt.Fprintln(w, FullnameRule)
		check = false
	}

	if !CheckBirthday(birthday) {
		fmt.Fprintln(w, BirthdayRule)
		check = false
	}

	return check
}

func LoginCheck(w http.ResponseWriter, r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")

	check := true

	if !CheckUsername(username) {
		fmt.Fprintln(w, UsernameRule)
		check = false
	}

	if !CheckPassword(password) {
		fmt.Fprintln(w, PasswordRule)
		check = false
	}

	return check
}
