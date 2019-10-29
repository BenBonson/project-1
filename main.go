package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

//var top = exec.Command("top")
// type Topstru struct {
// 	Name string
// }

//Topcom creates a strust for the output of top
type Topcom struct {
	TOP []string
}

//Logcom creates a strust for the output of log
type Logcom struct {
	LOG []string
}

//Indexcom creates a strust for the output of index
type Indexcom struct {
	IND []string
}

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

//index runs the index page
func index(response http.ResponseWriter, request *http.Request) {
	temp, _ := template.ParseFiles("html/index.html")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	ind := exec.Command("lsb_release", "-a")
	indout, stderr := ind.Output()
	in := Indexcom{IND: make([]string, 1)}
	var count int = 0

	if stderr != nil {
		fmt.Println(stderr)
	}
	for i := 0; i < len(indout); i++ {
		if indout[i] != 10 {
			in.IND[count] += string(indout[i])
		} else {
			count++
			in.IND = append(in.IND, "")
		}
	}

	temp.Execute(response, in)
	//temp := template.Must(template.ParseFiles("html/index.html"))
	//template.Must(template.New(".").Parse(temp))
	//temp.Execute(response, nil)
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

	top := exec.Command("top", "-b", "-n", "1")
	topout, stderr := top.Output()
	t := Topcom{TOP: make([]string, 1)}
	var count int = 0

	if stderr != nil {
		fmt.Println(stderr)
	}
	for i := 0; i < len(topout); i++ {
		if topout[i] != 10 {
			t.TOP[count] += string(topout[i])
		} else {
			count++
			t.TOP = append(t.TOP, "")
		}
	}

	temp.Execute(response, t)

	// for "\n"; tops {
	// 	topbreak = strings.Replace(tops, "\n", "<br>", -1)
	// }

	//topbreak := strings.Replace(tops, "\n", "<br/><br/>", -1)

	// data := struct{ TOP string }{tops}

	// err = temp.Execute(os.Stdout, data)
	// if err != nil {
	// 	fmt.Println("error two")
	// 	panic(err)
	// }
	// temp.Execute(response, data)
}

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
// 	data := WriteToFile("top.txt", tops)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// data := struct{ TOP string }{tops}

// 	// err = temp.Execute(os.Stdout, data)
// 	// if err != nil {
// 	// 	fmt.Println("error two")
// 	// 	panic(err)
// 	// }
// 	temp.Execute(response, data)
// }

// func WriteToFile(filename string, data string) error {
// 	cmd := exec.Command("rm", "top.txt")
// 	cmd.Start()
// 	cmd.Wait()
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.WriteString(file, data)
// 	if err != nil {
// 		return err
// 	}
// 	return file.Sync()
// }

// func (toop *Topcom) AddItem(item Topstru) []Topstru {
// 	toop.Items = append(toop.Items, item)
// 	return toop.Items
// }

// func main() {

// 	item1 := Topstru{Name: "Test Item 1"}

// 	items := []Topstru{}
// 	toop := Topcom{items}

// 	toop.AddItem(item1)

// 	fmt.Println(len(toop.Items))
// }

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
	log := exec.Command("tail", "-n", "25", "/var/log/syslog")
	logout, stderr := log.Output()
	l := Logcom{LOG: make([]string, 1)}
	var count int = 0

	if stderr != nil {
		fmt.Println(stderr)
	}
	//logs := string(log)
	//fmt.Println(tops)
	for i := 0; i < len(logout); i++ {
		if logout[i] != 10 {
			l.LOG[count] += string(logout[i])
		} else {
			count++
			l.LOG = append(l.LOG, "")
		}
	}

	//data := struct{ LOG string }{logs}

	// err = temp.Execute(os.Stdout, data)
	// if err != nil {
	// 	fmt.Println("error two")
	// 	panic(err)
	// }
	temp.Execute(response, l)
	//temp.Execute(response, nil)
}

//temp.Execute(response, t)

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
