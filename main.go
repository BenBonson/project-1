package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

//Topcom creates a struct for the output of top
type Topcom struct {
	TOP []string
}

//Logcom creates a struct for the output of log
type Logcom struct {
	LOG []string
}

//Syscom creates a struct for the output of sys
type Syscom struct {
	SYS []string
}

func main() {
	// Sets up a file server in current directory
	http.HandleFunc("/", index)
	http.HandleFunc("/top.html", topfunc)
	http.HandleFunc("/log.html", logfunc)
	http.HandleFunc("/sysinfo.html", sysfunc)
	http.HandleFunc("/localtop.html", localtop)
	http.HandleFunc("/locallog.html", locallog)
	http.HandleFunc("/localsys.html", localsys)

	//For each endpoint, call a function described below
	http.ListenAndServe(":7000", nil)
}

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	temp.Execute(response, nil)
}

//sysfunc runs system info
func sysfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/sysinfo.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	syst, err := exec.Command("ssh", "user1@192.168.56.102", "lsb_release", "-a").Output()
	sy := Syscom{SYS: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(syst); l = l + 1 {
		if syst[l] != 10 {
			sy.SYS[length] = sy.SYS[length] + string(syst[l])
		} else {
			sy.SYS = append(sy.SYS, "")
			length = length + 1
		}
	}

	temp.Execute(response, sy)
}

//sys runs system info
func localsys(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/localsys.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	syst, err := exec.Command("lsb_release", "-a").Output()
	sy := Syscom{SYS: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(syst); l = l + 1 {
		if syst[l] != 10 {
			sy.SYS[length] = sy.SYS[length] + string(syst[l])
		} else {
			sy.SYS = append(sy.SYS, "")
			length = length + 1
		}
	}

	temp.Execute(response, sy)
}

//top gets the results of the top command and sends the results to the html page
func localtop(response http.ResponseWriter, request *http.Request) {
	//connects to html and uses go templates
	temp, _ := template.ParseFiles("html/localtop.html")
	//this line makes it work w/html
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	//cmd := exec.Command("top")
	//temp.Execute(response, Signin)
	//top := exec.Command("top")

	//run top commamd and store the output as top
	top, err := exec.Command("top", "-b", "-n", "1").Output()
	//use the value t to put data into the struct
	t := Topcom{TOP: make([]string, 1)}
	//add the variable length which will be used to determine the row number
	length := 0

	//error checking for good practice
	if err != nil {
		fmt.Println(err)
	}
	//sets l to 0 and then adds 1 to said value every loop, stops looping when all lines have been looped
	for l := 0; l < len(top); l = l + 1 {
		//checks for breaks in the code (somehow 10 is the syntax for this new line)
		if top[l] != 10 {
			//if there is NOT a break in the code add the string to a slice in the struct = to the current length
			//value...this is the row number
			t.TOP[length] = t.TOP[length] + string(top[l])
		} else {
			//after the start of a new line/break append a blank for the next line to start in
			//example if it was (t.TOP, "lol") all text after the first line will have lol infront and an extra
			//lol will be printed after by itself
			t.TOP = append(t.TOP, "")
			//adds 1 to the length
			length = length + 1
		}
	}
	//runs output in html using go template {{range .TOP}}
	temp.Execute(response, t)
}

//top gets the results of the top command and sends the results to the html page
func topfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/top.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	top, err := exec.Command("ssh", "user1@192.168.56.102", "top", "-b", "-n", "1").Output()
	t := Topcom{TOP: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}
	for l := 0; l < len(top); l = l + 1 {
		if top[l] != 10 {
			t.TOP[length] = t.TOP[length] + string(top[l])
		} else {
			t.TOP = append(t.TOP, "")
			length = length + 1
		}
	}

	temp.Execute(response, t)
}

//logfunc gets the logs and sends the results to the html page
func logfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/log.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")

	log, err := exec.Command("ssh", "user1@192.168.56.102", "tail", "-n", "25", "/var/log/syslog").Output()
	lo := Logcom{LOG: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}

	for l := 0; l < len(log); l = l + 1 {
		if log[l] != 10 {
			lo.LOG[length] = lo.LOG[length] + string(log[l])
		} else {
			lo.LOG = append(lo.LOG, "")
			length = length + 1
		}
	}
	temp.Execute(response, lo)
}

//locallog gets the logs and sends the results to the html page
func locallog(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/locallog.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//cmd := exec.Command("top")
	//temp.Execute(response, Signin)
	log, err := exec.Command("tail", "-n", "25", "/var/log/syslog").Output()
	lo := Logcom{LOG: make([]string, 1)}
	length := 0

	if err != nil {
		fmt.Println(err)
	}
	//logs := string(log)
	//fmt.Println(tops)
	for l := 0; l < len(log); l = l + 1 {
		if log[l] != 10 {
			lo.LOG[length] = lo.LOG[length] + string(log[l])
		} else {
			lo.LOG = append(lo.LOG, "")
			length = length + 1
		}
	}
	temp.Execute(response, lo)
}
