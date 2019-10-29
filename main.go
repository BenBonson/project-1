package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
)

//var top = exec.Command("top")

// type Topcom struct {
// 	TOP []string
// }

func main() {
	// Sets up a file server in current directory
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", index)
	http.HandleFunc("/top.html", topfunc)
	http.HandleFunc("/log.html", logfunc)

	//For each endpoint, call a function described below
	// http.HandleFunc("/hello", hello)
	// //http.HandleFunc("/logs", logs)
	// http.HandleFunc("/hellohtml", hellohtml)
	// http.HandleFunc("/formsubmit", formsubmit)
	// http.HandleFunc("/template", hellotemphandler)
	http.ListenAndServe(":7000", nil)
	//pretty sure l and s is what sets default index page to be localhost:9000
	// ssh()
}

func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//temp := template.Must(template.ParseFiles("html/index.html"))
	//template.Must(template.New(".").Parse(temp))
	temp.Execute(response, nil)
	//temp.Execute(response, temp)
}

//prints out all html with top included in commandline????
// func topfunc(response http.ResponseWriter, request *http.Request) {
// 	temp, _ := template.ParseFiles("html/top.html")
// 	response.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	//cmd := exec.Command("top")
// 	//temp.Execute(response, Signin)
// 	//top := exec.Command("top")

// 	top, err := exec.Command("top", "-b", "-n", "1").Output()
// 	if err != nil {
// 		fmt.Println("here")
// 		log.Fatal(err)
// 	}
// 	tops := string(top)
// 	//fmt.Println(tops)

// 	data := struct{ TOP string }{tops}

// 	err = temp.Execute(os.Stdout, data)
// 	if err != nil {
// 		fmt.Println("error two")
// 		panic(err)
// 	}
// 	temp.Execute(response, nil)
// }

//top gets the results of the top command and sends the results to the html page its ugly as sin but it works
func topfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/top.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//cmd := exec.Command("top")
	//temp.Execute(response, Signin)
	//top := exec.Command("top")

	top, err := exec.Command("top", "-b", "-n", "1").Output()
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}
	tops := string(top)
	//fmt.Println(tops)

	data := struct{ TOP string }{tops}

	err = temp.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("error two")
		panic(err)
	}
	temp.Execute(response, data)
}

// func topfunc(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("top.html")
// 	cmd := exec.Command("top", "-b", "-n", "1")
// 	stdout, err := cmd.Output()

// 	var count int = 0
// 	if err != nil {
// 		println(err.Error())
// 		return
// 	}

// 	t := Topcom{TOP: make([]string, 1)}

// 	for i := 1; i < len(stdout); i++ {
// 		if stdout[i] != 10 {
// 			t.TOP[count] += string(stdout[i])
// 		} else {
// 			count++
// 			t.TOP = append(t.TOP, "")
// 		}
// 	}
// 	// output := string(stdout)
// 	// cmd.Start()
// 	// cmd.Wait()
// 	// fmt.Println(cmd)
// 	temp.Execute(w, t)
// }

//combined w/gmac
// func topfunc(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("applications.html")
// 	cmd := exec.Command("top", "-b", "-n", "1")
// 	cmdout, stderr := cmd.Output()
// 	t := Topcom{TOP: make([]string, 1)}
// 	var count int = 0

// 	if stderr != nil {
// 		fmt.Println(stderr)
// 	}
// 	for i := 1; i < len(cmdout); i++ {
// 		if cmdout[i] != 10 {
// 			t.TOP[count] += string(cmdout[i])
// 		} else {
// 			count++
// 			t.TOP = append(t.TOP, "")
// 		}
// 	}

// 	temp.Execute(w, t)
// }

//prints out in command line but not html
// func topfunc(response http.ResponseWriter, request *http.Request) {
// 	temp, _ := template.ParseFiles("html/top.html")
// 	response.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	//cmd := exec.Command("top")
// 	//temp.Execute(response, Signin)
// 	//top := exec.Command("top")

// 	cmd := exec.Command("top", "-b", "-n", "1")
// 	stdout, err := cmd.Output()

// 	if err != nil {
// 		println(err.Error())
// 		return
// 	}

// 	tops := string(stdout)

// 	data := struct{ TOP string }{tops}

// 	err = temp.Execute(os.Stdout, data)
// 	if err != nil {
// 		fmt.Println("error two")
// 		panic(err)
// 	}
// 	print(string(stdout))

// 	//temp.Execute(response, top)
// 	//temp.Execute(response, nil)
// }

//logfunc gets the logs and sends the results to the html page its ugly as sin but it works
func logfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/log.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	//cmd := exec.Command("top")
	//temp.Execute(response, Signin)
	log, err := exec.Command("tail", "-n", "5", "/var/log/syslog").Output()
	if err != nil {
		fmt.Println("here")
		panic(err)
	}
	logs := string(log)
	//fmt.Println(tops)

	data := struct{ LOG string }{logs}

	err = temp.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("error two")
		panic(err)
	}
	temp.Execute(response, data)
	//temp.Execute(response, nil)
}

// func ssh() {
// 	cmd := exec.Command("ssh", "user1@192.168.56.102")
// 	reader, _ := cmd.StdoutPipe()
// 	scanner := bufio.NewScanner(reader)
// 	go func() {
// 		for scanner.Scan() {
// 			fmt.Println("user1> ", scanner.Text())
// 		}
// 	}()
// 	cmd.Start()
// 	cmd.Wait()
// }

// func hello(response http.ResponseWriter, request *http.Request) {
// 	navigation
// 	name := request.FormValue("name")
// 	output := []byte("Hello " + name)
// 	fmt.Println(name, "says hello!")
// 	response.Write(output)
// }

// //prints what is in html on this page
// func hellohtml(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "text/html")
// 	//output := []byte("<html><body><h1>Hello There!</h1></body></html>")
// 	//response.Write(output)
// 	io.WriteString(response, `
// 	<DOCTYPE html>
// 	<html>
// 	<head>
// 		<title>My Page</title>
// 	</head>
// 	<body>
// 		<h2>Welcome to My Page</h2>
// 		<p>This is a test of a go server</p>
// 	</body>
// 	</html>
// 	`)
// }

// //prints nothing by itself
// //however if localhost:9000 is run first and the form filled it will print welcome username and their password
// //in command line
// func formsubmit(response http.ResponseWriter, request *http.Request) {
// 	fmt.Println("Welcome", request.FormValue("user"))
// 	fmt.Println("Your password is", request.FormValue("password"))
// }

// //template stuff prints Hello, Bob!
// const hellotemplate = `
// <!DOCTYPE html>
// <html>
// 	<head>
// 		<title>Template Page</title>
// 	</head>
// 	<body>
// 		<h1>Hello, {{.Name}}!</h1>
// 	</body>
// </html>
// `

// var hellotmpl = template.Must(template.New(".").Parse(hellotemplate))

// func hellotemphandler(response http.ResponseWriter, request *http.Request) {
// 	myCount.Add(1)
// 	myStatus.Set("Good")
// 	hellotmpl.Execute(response, map[string]interface{}{
// 		"Name": "Bob",
// 	})
// }
