package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/najidroid/quizOfStudents/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title Get
// @Description get all Users
// @Success 200 {object} models.User
// @router / user [get]
func (u *UserController) GetAll() {
	users := models.SetUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get schools by city
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :city is empty
// @router /getschools [get]
func (u *UserController) GetSchools() {
	fmt.Println("***************GET SCHOOLS***************")
	schools := models.GetSchools()
	//	fmt.Println(schools)
	u.Data["json"] = schools
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /addstudent [post]
func (u *UserController) AddStudent() {
	fmt.Println("***************ADD STUDENT***************")
	var Ob models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	student, ok := models.AddStudent(Ob)
	if ok {
		u.Data["json"] = student
	} else {
		u.Data["json"] = nil
	}
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getstudent [post]
func (u *UserController) GetStudent() {
	fmt.Println("***************GET STUDENT***************")
	var Ob models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println("student0 is: ", Ob)
	student := models.GetStudent(Ob)
	u.Data["json"] = student
	fmt.Println("student is: ", student)
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getmyschool [post]
func (u *UserController) GetMySchool() {
	fmt.Println("***************GET MY SCHOOL***************")
	var Ob models.School
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	student := models.GetMySchool(Ob)
	u.Data["json"] = student
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getschool [post]
func (u *UserController) GetSchool() {
	fmt.Println("***************GET SCHOOL***************")
	var Ob models.School
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	school := models.GetSchool(Ob)
	u.Data["json"] = school
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /friendrequest [post]
func (u *UserController) FriendRequest() {
	fmt.Println("***************FRIEND REQUEST***************")
	var Ob models.Friend
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	friend := models.RequestFriendShip(Ob)
	fmt.Println(Ob)
	u.Data["json"] = friend
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /acceptfriendship [post]
func (u *UserController) AcceptFriendship() {
	fmt.Println("***************ACCEPT FRIENDSHIP***************")
	var Ob models.Friend
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	friend := models.AcceptFriendShip(Ob)
	u.Data["json"] = friend
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getmyfriends [post]
func (u *UserController) GetMyFriends() {
	fmt.Println("***************GET MY FRIENDS***************")
	var Ob models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	friends := models.GetMyFriends(Ob)
	fmt.Println("my friends: ", friends)
	u.Data["json"] = friends
	u.ServeJSON()
}

// @Title Get
// @Description start match
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /ismatch [post]
func (u *UserController) IsMatch() {
	fmt.Println("***************IS MATCH***************")
	var Ob models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	match := models.IsMatch(Ob)
	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description start match
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /startmatch [post]
func (u *UserController) StartMatch() {
	fmt.Println("***************START MATCH***************")
	var Ob models.Match
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println("error:", err)
	fmt.Println(Ob)
	match := models.StartMatch(Ob)
	fmt.Println("*****match is:", match)
	fmt.Println("*****match students is:", match.Students[0])
	fmt.Println("*****match submatches is:", match.SubMatches[0])

	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description start match
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /startmatchforfriendreq [post]
func (u *UserController) StartMatchForFriendReq() {
	fmt.Println("***************START MATCH FOR FRIEND REQUEST***************")
	var Ob models.Match
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println("error:", err)
	fmt.Println(Ob)
	match := models.StartMatchForFriendReq(Ob)
	fmt.Println("*****match is:", match)
	fmt.Println("*****match students is:", match.Students)
	fmt.Println("*****match submatches is:", match.SubMatches[0])

	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getsubmatch [post]
func (u *UserController) GetSubMatch() {
	fmt.Println("***************GET SUBMATCH***************")
	var Ob models.SubMatch
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	subMatch := models.GetSubMatch(Ob)
	fmt.Println("*****submatch is:", subMatch)
	u.Data["json"] = subMatch
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /acceptmatch [post]
func (u *UserController) AcceptMatch() {
	fmt.Println("***************ACCEPT MATCH***************")
	var Ob models.Match
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	subMatch := models.AcceptMatch(Ob)
	fmt.Println("*****accepted match is:", subMatch)
	u.Data["json"] = subMatch
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getmatch [post]
func (u *UserController) GetMatch() {
	fmt.Println("***************GET MATCH***************")
	var Ob models.Match
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	match := models.GetMatch(Ob)
	fmt.Println("*****match is:", match)
	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /updatestudent [post]
func (u *UserController) UpdateStudent() {
	fmt.Println("***************UPDATE STUDENT***************")
	var Ob models.Student
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	student := models.UpdateStudent(Ob)
	fmt.Println("*****student is:", student)
	u.Data["json"] = student
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /addsubmatch [post]
func (u *UserController) AddSubMatch() {
	fmt.Println("***************ADD SUBMATCH***************")
	var Ob models.Match
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	subMatch := models.AddSubMatch(Ob)
	fmt.Println("*****added submatch is:", subMatch)
	u.Data["json"] = subMatch
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /updatesubmatch [post]
func (u *UserController) UpdateSubMatch() {
	fmt.Println("***************UPDATE SUBMATCH***************")
	var Ob models.SubMatch
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	subMatch := models.UpdateSubMatch(Ob)
	fmt.Println("*****sub match is:", subMatch)
	u.Data["json"] = subMatch
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /requestmatch [post]
func (u *UserController) RequestMatch() {
	fmt.Println("***************REQUEST MATCH***************")
	var Ob models.Match
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	match := models.RequestMatch(Ob)
	fmt.Println("*****match is:", match)
	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /acceptmatchrequest [post]
func (u *UserController) AcceptMatchRequest() {
	fmt.Println("***************ACCEPT MATCH REQUEST***************")
	var Ob models.Match
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	match := models.AcceptMatchRequest(Ob)
	fmt.Println("*****match is:", match)
	u.Data["json"] = match
	u.ServeJSON()
}

// @Title Get
// @Description add student
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uimei is empty
// @router /getfriend [post]
func (u *UserController) GetFriend() {
	fmt.Println("***************GET FRIEND***************")
	var Ob models.Friend
	json.Unmarshal(u.Ctx.Input.RequestBody, &Ob)
	fmt.Println(Ob)
	friend := models.GetFriend(Ob)
	fmt.Println("*****friend is:", friend)
	u.Data["json"] = friend
	u.ServeJSON()
}
