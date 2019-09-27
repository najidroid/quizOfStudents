package models

import (
	"fmt"
	"math/rand"

	"strings"

	"strconv"

	"time"

	"github.com/astaxie/beego/orm"

	"github.com/claudiu/gocron"
)

func init() {
	gocron.Start()
	s := gocron.NewScheduler()
	gocron.Every(2).Minutes().Do(checkForLockedMatches)
	s.Start()
}

func checkForLockedMatches() {
	fmt.Println("***************CHECKING FOR LOCKED MATCHES***************")
	o := orm.NewOrm()
	var matches []*Match
	o.QueryTable("Match").Filter("IsLocked", true).OrderBy("Time").RelatedSel().All(&matches)

	for _, item := range matches {
		if time.Now().Sub(item.LockTime).Seconds() > 120 {
			item.IsLocked = false
			o.Update(item)
			fmt.Println("item updated, id:", item.Id)
		}
	}
}

func SetUsers() []*School {
	var data []*School
	orm.NewOrm().QueryTable(new(School)).All(&data)
	fmt.Println(data)
	push("topic", "hi")
	return data
}

func GetStudent(std Student) *Student {
	student := Student{Id: std.Id}
	orm.NewOrm().Read(&student)
	fmt.Println("student1 is: ", student)
	orm.NewOrm().LoadRelated(&student, "Friends")
	orm.NewOrm().LoadRelated(&student, "Matches")
	if std.Token == student.Token {
		return &student
	} else {
		return nil
	}
}

func GetMyFriends(std Student) []*Friend {
	student := Student{Id: std.Id}
	orm.NewOrm().Read(&student)
	fmt.Println("student is:", student)
	orm.NewOrm().LoadRelated(&student, "Friends")
	return student.Friends
}

func AddStudent(std Student) (*Student, bool) {
	o := orm.NewOrm()
	o.Begin()
	student := Student{Name: std.Name, FamilyName: std.FamilyName,
		School: getSchool(std.School.Id), Grade: std.Grade, Field: std.Field,
		EmailAdress: std.EmailAdress, MobileNumber: std.MobileNumber, AvatarCode: "1",
		Token: TokenGenerator(), SchoolName: std.School.Name}
	o.Insert(&student)
	friend := Friend{Name: std.Name, FamilyName: std.FamilyName, MobileNumber: std.MobileNumber,
		Grade: std.Grade, Field: std.Field, SchoolName: std.School.Name, AvatarCode: "1",
		StudentId: student.Id}
	o.Insert(&friend)
	o.Commit()
	return &student, true
}

func getSchool(id int) *School {
	school := &School{}
	orm.NewOrm().QueryTable("School").Filter("Id", id).RelatedSel().One(school)
	return school
}

func GetMySchool(sch School) *School {
	school := School{Id: sch.Id}
	orm.NewOrm().Read(&school)
	orm.NewOrm().LoadRelated(&school, "Students")
	return &school
}

func GetSchool(sch School) *School {
	school := School{Id: sch.Id}
	orm.NewOrm().Read(&school)
	orm.NewOrm().LoadRelated(&school, "Students")
	return &school
}

func GetSchools() []*School {
	var schools []*School
	orm.NewOrm().QueryTable(new(School)).OrderBy("-Score").All(&schools)
	return schools
}

func RequestFriendShip(friend Friend) *Friend {
	o := orm.NewOrm()
	o.Begin()
	std1 := Student{Id: friend.Students[0].Id}
	std2 := Student{Id: friend.Students[1].Id}
	o.Read(&std2)
	o.Read(&std1)
	var friend1 Friend
	var friend2 Friend
	o.QueryTable("Friend").Filter("MobileNumber", std1.MobileNumber).RelatedSel().One(&friend1)
	o.QueryTable("Friend").Filter("MobileNumber", std2.MobileNumber).RelatedSel().One(&friend2)

	std2.Friends = append(std2.Friends, &friend1)
	std1.Friends = append(std1.Friends, &friend2)
	std2.FriendRequests += "," + strconv.Itoa(friend1.Id) + ","
	std1.RequestedFriends += "," + strconv.Itoa(friend2.Id) + ","
	m2m1 := o.QueryM2M(&std1, "Friends")
	m2m1.Add(&friend2)
	m2m2 := o.QueryM2M(&std2, "Friends")
	m2m2.Add(&friend1)
	o.Update(&std2)
	o.Update(&std1)

	fmt.Println("std1:", std1)
	fmt.Println("std2:", std2)
	fmt.Println("friend1:", friend1)
	fmt.Println("friend2:", friend2)
	o.Commit()
	//	push("update/id-"+strconv.Itoa(std1.Id), "FR-"+strconv.Itoa(friend1.Id))
	push("update/id-"+strconv.Itoa(std2.Id), "FR-"+strconv.Itoa(friend1.Id))
	return &friend2
}

func AcceptFriendShip(friend1 Friend) *Friend {
	o := orm.NewOrm()
	o.Begin()
	std2 := Student{Id: friend1.Students[0].Id}
	o.Read(&std2)
	o.Read(&friend1)

	var friend2 Friend
	var std1 Student
	o.QueryTable("Friend").Filter("MobileNumber", std2.MobileNumber).RelatedSel().One(&friend2)
	o.QueryTable("Student").Filter("MobileNumber", friend1.MobileNumber).RelatedSel().One(&std1)
	std2.FriendRequests = strings.Replace(std2.FriendRequests, ","+strconv.Itoa(friend1.Id)+",", "", -1)
	std1.RequestedFriends = strings.Replace(std1.RequestedFriends, ","+strconv.Itoa(friend2.Id)+",", "", -1)
	o.Update(&std2)
	o.Update(&std1)
	fmt.Println("friend1:", friend1)
	fmt.Println("friend2:", friend2)
	fmt.Println("srudent1:", std1)
	fmt.Println("srudent2:", std2)
	o.Commit()
	//	push("update/id-"+strconv.Itoa(std1.Id), "AF-"+strconv.Itoa(friend2.Id))
	push("update/id-"+strconv.Itoa(std1.Id), "AF-"+strconv.Itoa(friend2.Id))
	return &friend1
}

func IsMatch(std Student) *Match {
	o := orm.NewOrm()
	match := Match{}
	o.QueryTable("Match").Filter("State", 0).Filter("IsLocked", false).OrderBy("Time").RelatedSel().One(&match)
	fmt.Println("isMatch/ match:", match)
	if match.Id != 0 {
		o.LoadRelated(&match, "Students")
		if match.Students[0].Id != std.Id {
			fmt.Println("here1")
			match.IsLocked = true
			o.Update(&match)
			o.LoadRelated(&match, "Students")
			o.LoadRelated(&match, "SubMatches")
		} else {
			match = Match{}
		}
	}
	return &match
}

func StartMatch(match Match) *Match {
	o := orm.NewOrm()
	o.Begin()
	frstStd := match.Students[0]
	fmt.Println("StartMatch/ frstStd:", frstStd)
	sbMch := match.SubMatches[0]
	subject := sbMch.Subject
	subMatchQuestions := getDBQuestions(subject)
	sbMch.Questions = subMatchQuestions
	o.Insert(&match)
	sbMch.Match = &match
	o.Insert(sbMch)
	m2m := o.QueryM2M(&match, "Students")
	m2m.Add(frstStd)
	m2m1 := o.QueryM2M(sbMch, "Questions")
	m2m1.Add(subMatchQuestions)
	o.Read(&match)
	o.LoadRelated(&match, "Students")
	o.LoadRelated(&match, "SubMatches")
	match.SubMatches[0].Questions = subMatchQuestions
	fmt.Println("***StartMatch/ SubMatches***", match.SubMatches)
	o.Commit()
	return &match
}

func StartMatchForFriendReq(match Match) *Match {
	o := orm.NewOrm()
	o.Begin()
	sbMch := match.SubMatches[0]
	subject := sbMch.Subject
	subMatchQuestions := getDBQuestions(subject)
	sbMch.Questions = subMatchQuestions
	sbMch.Match = &match
	o.Insert(sbMch)
	m2m := o.QueryM2M(sbMch, "Questions")
	m2m.Add(subMatchQuestions)
	o.Read(&match)
	o.LoadRelated(&match, "Students")
	o.LoadRelated(&match, "SubMatches")
	match.SubMatches[0].Questions = subMatchQuestions
	//	o.LoadRelated(&match.SubMatches[0], "Questions")
	o.Commit()
	return &match
}

func UpdateSubMatch(sbMch SubMatch) *Match {
	o := orm.NewOrm()
	o.Begin()
	o.Update(&sbMch)
	o.LoadRelated(&sbMch, "Match")
	mch := sbMch.Match

	mch.PlayedId = sbMch.PlayedId
	if mch.FirstId == 0 {
		mch.FirstId = sbMch.FirstId
		mch.SecondId = sbMch.SecondId
	}
	if mch.FirstId == sbMch.PlayedId {
		mch.FirstTotalAnswers += getNumberOfAnswersFromString(sbMch.FirstAnswers)
		push("update/id-"+strconv.Itoa(mch.SecondId), "MCH-"+strconv.Itoa(mch.Id))
	} else {
		mch.SecondTotalAnswers += getNumberOfAnswersFromString(sbMch.SecondAnswers)
		push("update/id-"+strconv.Itoa(mch.FirstId), "MCH-"+strconv.Itoa(mch.Id))
	}
	o.LoadRelated(mch, "SubMatches")
	if len(mch.SubMatches) >= 10 &&
		len(mch.SubMatches[9].FirstAnswers) >= 3 && len(mch.SubMatches[9].SecondAnswers) >= 3 {
		mch.State = 2
	}
	o.Update(mch)
	o.LoadRelated(mch, "SubMatches")
	o.LoadRelated(mch, "Students")
	o.Commit()
	fmt.Println("mch:", mch)
	return mch
}

func getNumberOfAnswersFromString(answer string) int {
	fmt.Println(answer)
	answer = strings.Replace(answer, "null", "", -1)
	num := 0
	for i := 0; i < 3; i++ {
		fmt.Println("num:", i)
		if "0" == answer[i:i+1] {
			num++
		}
	}
	return num
}

func GetSubMatch(sbMch SubMatch) *SubMatch {
	orm.NewOrm().Read(&sbMch)
	_, err := orm.NewOrm().LoadRelated(&sbMch, "Questions")
	fmt.Println("error:", err)
	return &sbMch
}

func GetMatch(mch Match) *Match {
	orm.NewOrm().Read(&mch)
	orm.NewOrm().LoadRelated(&mch, "SubMAtches")
	orm.NewOrm().LoadRelated(&mch, "Students")
	fmt.Println("match is:", mch)
	return &mch
}

func UpdateStudent(std Student) *Student {
	orm.NewOrm().Read(&std)
	orm.NewOrm().LoadRelated(&std, "Matches")
	fmt.Println("student is:", &std)
	return &std
}

func AcceptMatch(match Match) *SubMatch {
	mch := Match{Id: match.Id}
	orm.NewOrm().Read(&mch)
	orm.NewOrm().LoadRelated(&mch, "Students")
	fmt.Println("mch is:", mch)
	sbMch := match.SubMatches[0]
	orm.NewOrm().Update(&match)
	orm.NewOrm().Update(sbMch)
	orm.NewOrm().LoadRelated(sbMch, "Questions")
	if len(mch.Students) == 1 {
		m2m := orm.NewOrm().QueryM2M(&match, "Students")
		m2m.Add(match.Students[1])
		push("update/id-"+strconv.Itoa(match.Students[1].Id), "MCH-"+strconv.Itoa(match.Id))
	}
	return sbMch
}

func AddSubMatch(match Match) *SubMatch {
	o := orm.NewOrm()
	o.Begin()
	sm := match.SubMatches[len(match.SubMatches)-1]
	subject := sm.Subject
	subMatchQuestions := getDBQuestions(subject)
	sbMch := SubMatch{Subject: sm.Subject, FirstId: sm.FirstId, SecondId: sm.SecondId, FirstName: sm.FirstName,
		FirstFamilyName: sm.FirstFamilyName, SecondName: sm.SecondName, SecondFamilyName: sm.SecondFamilyName,
		State: 0, Questions: subMatchQuestions, Match: &match}
	match.PlayedId = sm.PlayedId
	var opId int
	if match.PlayedId == sm.FirstId {
		opId = sm.FirstId
	} else {
		opId = sm.SecondId
	}
	o.Insert(&sbMch)
	o.Update(&match)
	m2m := o.QueryM2M(&sbMch, "Questions")
	m2m.Add(subMatchQuestions)
	o.Commit()
	push("update/id-"+strconv.Itoa(opId), "MCH-"+strconv.Itoa(match.Id))
	return &sbMch
}

func RequestMatch(mch Match) *Match {
	o := orm.NewOrm()
	o.Begin()
	std1 := mch.Students[0]
	std2 := mch.Students[1]
	o.Read(std2)
	o.Read(std1)
	mch.FirstOpName = std1.Name
	mch.SecondOpName = std2.Name
	mch.State = 1
	mch.ShowInClient1 = true
	mch.ShowInClient2 = true
	mch.PlayedId = std1.Id
	o.Insert(&mch)
	std1.RequestedMatch += "," + strconv.Itoa(mch.Id) + ","
	std2.MatchRequests += "," + strconv.Itoa(mch.Id) + ","
	o.Update(std1)
	o.Update(std2)
	m2m := o.QueryM2M(&mch, "Students")
	m2m.Add(std1)
	m2m.Add(std2)
	o.Commit()
	//	push("update/id-"+strconv.Itoa(std1.Id), "MCH")
	push("update/id-"+strconv.Itoa(std2.Id), "MCH-"+strconv.Itoa(mch.Id))
	return &mch
}

func AcceptMatchRequest(mch Match) *Match {
	o := orm.NewOrm()
	o.Begin()
	o.LoadRelated(&mch, "Students")
	std1 := mch.Students[0]
	std2 := mch.Students[1]
	o.Read(std2)
	o.Read(std1)
	fmt.Println("std1.RequestedMatch", std1.RequestedMatch)
	fmt.Println("std2.RequestedMatch", std2.RequestedMatch)
	std1.RequestedMatch = strings.Replace(std1.RequestedMatch, ","+strconv.Itoa(mch.Id)+",", "", -1)
	std2.MatchRequests = strings.Replace(std2.MatchRequests, ","+strconv.Itoa(mch.Id)+",", "", -1)
	std2.RequestedMatch = strings.Replace(std2.RequestedMatch, ","+strconv.Itoa(mch.Id)+",", "", -1)
	std1.MatchRequests = strings.Replace(std1.MatchRequests, ","+strconv.Itoa(mch.Id)+",", "", -1)
	o.Update(std1)
	o.Update(std2)
	fmt.Println("std1.RequestedMatch", std1.RequestedMatch)
	fmt.Println("std2.RequestedMatch", std2.RequestedMatch)
	o.Commit()
	push("update/id-"+strconv.Itoa(std1.Id), "MCH-"+strconv.Itoa(mch.Id))
	return &mch
}

func GetFriend(fr Friend) *Friend {
	orm.NewOrm().Read(&fr)
	return &fr
}

func getDBQuestions(subject string) []*Question {
	fmt.Println("subject:", subject)

	var questions []*Question
	orm.NewOrm().QueryTable("Question").Filter("Subject", subject).All(&questions)
	//	fmt.Println("questions:", questions)
	var subMatchQuestions []*Question
	for i := 0; i < 3; i++ {
		n := rand.Intn(len(questions))
		q := questions[n]
		//		fmt.Println("random number", n)
		subMatchQuestions = append(subMatchQuestions, q)
	}
	return subMatchQuestions
}

func TokenGenerator() string {
	//	b := make([]byte, 15)
	//	rand.Read(b)
	bytes := make([]byte, 15)
	for i := 0; i < 15; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
	//	return fmt.Sprintf("%x", b)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
