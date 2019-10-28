package main

import (
	"html/template"
	"net/http"
	"os/exec"
)

//var top = exec.Command("top")

func main() {
	// Sets up a file server in current directory
	//http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/", index)
	http.HandleFunc("/top", topfunc)

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
	//temp.Execute(response, Signin)
	temp.Execute(response, nil)
	//temp.Execute(response, temp)
}

func topfunc(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/top.html")
	//cmd := exec.Command("top")
	//temp.Execute(response, Signin)
	top := exec.Command("top")
	temp.Execute(response, top)
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
