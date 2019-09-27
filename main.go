package main

import (
	"github.com/astaxie/beego"
	"github.com/najidroid/quizOfStudents/models"
	_ "github.com/najidroid/quizOfStudents/routers"

	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"github.com/astaxie/beego/orm"

	//	"log"
	"encoding/json"
	"io/ioutil"
	"os"

	//	"flag"
	//	"log"
	//	"net/http"

	//	"text/template"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "u0rjdhj3uers2hge:GK0UooOMJL9Yf2Jc1HAd@tcp(bfcegwfzeq1o1qcwfe2y-mysql.services.clever-cloud.com:3306)/bfcegwfzeq1o1qcwfe2y?charset=utf8")
}

func main() {
	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := false

	// Print log.
	verbose := false

	// Error.
	err := orm.RunSyncdb(name, force, verbose)

	if err != nil {
		fmt.Println(err)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	SetSchools()
	SetQuestions()

	beego.Run()

	//	//	this is for chat:?
	//	flag.Parse()
	//	go hub.run()
	//	http.HandleFunc("/", serveHome)
	//	http.HandleFunc("/ws", serveWs)
	//	err1 := http.ListenAndServe(*addr, nil)
	//	if err1 != nil {
	//		log.Fatal("ListenAndServe: ", err)
	//	}
}

func SetSchools() {
	var schools []*models.School
	orm.NewOrm().QueryTable(new(models.School)).All(&schools)
	if len(schools) != 0 {
		return
	}
	file, e := ioutil.ReadFile("./schools.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &schools)
	for _, item := range schools {
		_, err := orm.NewOrm().Insert(item)
		if err != nil {
			fmt.Println(err)
		}
		//		fmt.Println(item)
	}
	fmt.Println("schools added to DB")
}

func SetQuestions() {
	var questions []*models.Question
	orm.NewOrm().QueryTable(new(models.Question)).All(&questions)
	if len(questions) != 0 {
		return
	}
	file, e := ioutil.ReadFile("./questions.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	err1 := json.Unmarshal(file, &questions)
	if err1 != nil {
		fmt.Println(err1)
	}
	for _, item := range questions {
		_, err := orm.NewOrm().Insert(item)
		//		fmt.Println(item)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("questions added to DB")
}

////this is for chat
//var addr = flag.String("addr", ":8181", "http service address")
//var homeTemplate = template.Must(template.ParseFiles("home.html"))

//func serveHome(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.Error(w, "Not found", 404)
//		return
//	}
//	if r.Method != "GET" {
//		http.Error(w, "Method not allowed", 405)
//		return
//	}
//	w.Header().Set("Content-Type", "text/html; charset=utf-8")
//	homeTemplate.Execute(w, r.Host)
//}
