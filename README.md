# project-1




# records 
1. started with navigation, tried multiple html pages but figured that this should be more go centered
2. restarted and made an index html
3. index displays source code when in its own file but does it properly in the same one
4. server stopped working
5. in confusion tested older projects and examples, they don't work either...eventhough I haven't touched them...very confused
6. localhost:9000 wont work tried 3000 and 7000 and they work ???
7. Now index displays its source code in the same file ???
8. Lots of testing and things only broke more now attempting to connect to local host says connection was reset
9. found and fixed the issue but it still displays the source code not in html format
10. added response.Header().Set("Content-Type", "text/html; charset=utf-8") now html displays correctly
11. Tried to work on getting top to work, its not displayig anything
12. main.go now displays all the html with top included, however it does so in the command line
13. Got the results to print in HTML, its pretty ugly
14. tried to print output into text file then transfer text to html, ended up with same problem as before
15. Got it to print in a somewhat readable way via golang templates
16. Made table look a little better
17. Added functionality to all current functions
18. added descriptions to pages
19. added lots of html to try and ssh
20. added a .gitignore for passcheck for security
21. everything works but you are prompted in command line for a password every time you ssh
22. added ssh-key so no login is required
23. made html look nicer
24. removed unused code from main.go and explained the process in localtop function


# notes/references

https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/

https://stackoverflow.com/questions/1877045/how-do-you-get-the-output-of-a-system-command-in-go

https://stackoverflow.com/questions/46655138/golang-print-html-lines-dynamically


ssh-keygen -t rsa
ssh-copy-id user1@ipaddress
now i can use ssh user1@ipaddress w/out a password

# unused/broken code

//var top = exec.Command("top")
// type Topstru struct {
// 	Name string
// }

//http.Handle("/", http.FileServer(http.Dir(".")))
// http.HandleFunc("/pass.html", passfunc)
// http.HandleFunc("/hello", hello)
// //http.HandleFunc("/logs", logs)
// http.HandleFunc("/hellohtml", hellohtml)
// http.HandleFunc("/formsubmit", formsubmit)
// http.HandleFunc("/template", hellotemphandler)

// if Remote {

	// } else {

	// }
	//temp := template.Must(template.ParseFiles("html/index.html"))
	//template.Must(template.New(".").Parse(temp))

//temp.Execute(response, temp)

//temp := template.Must(template.ParseFiles("html/index.html"))
//template.Must(template.New(".").Parse(temp))
//temp.Execute(response, nil)
//temp.Execute(response, temp)


// //pass runs system info
// func passfunc(response http.ResponseWriter, request *http.Request) {
// 	temp, _ := template.ParseFiles("html/pass.html")
// 	response.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	// if Remote {

// 	// } else {

// 	// }
// 	//temp := template.Must(template.ParseFiles("html/index.html"))
// 	//template.Must(template.New(".").Parse(temp))
// 	temp.Execute(response, nil)
// 	//temp.Execute(response, temp)

// }

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

//cmd := exec.Command("top")
//temp.Execute(response, Signin)
//top := exec.Command("top")




// if _, err := client.Exec("cat /tmp/somefile"); err != nil {
//     log.Println(err)
// }

// return nil

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

//cmd := exec.Command("top")
//temp.Execute(response, Signin)
//password := passcheck.Pass()

//logs := string(log)
//fmt.Println(tops)





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
