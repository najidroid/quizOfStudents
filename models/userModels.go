package models

import (
	"time"

	"fmt"
	"os"

	"github.com/astaxie/beego/orm"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func push(subscription string, message string) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://soldier.cloudmqtt.com:12466")
	opts.SetClientID("quiz of students")
	opts.SetPassword("1ZG2oHDsATYa")
	opts.SetUsername("hnlbufmc")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		//		panic(token.Error())
		fmt.Println("error of connecting:", token.Error())
	}
	fmt.Println("subscription:", subscription, "\n message:", message)

	c.Publish(subscription, 0, false, message)
	//	c.Publish("update/id-2", 0, false, "message: johney")

	//time.Sleep(3 * time.Second)

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe(subscription); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}

type School struct {
	Id               int
	Students         []*Student `orm:"reverse(many)"`
	Name             string
	City             string
	Region           int
	TotalScore       int
	WeekScore        int
	WonMatches       int
	TotalMatches     int
	EvenMatches      int
	WeekWonMatches   int
	WeekTotalMatches int
	WeekEvenMatches  int
	TotalRank        int
	WeekRank         int
	WeekRankArray    string `orm:"size(500)"`
}

type Student struct {
	Id          int
	Name        string
	FamilyName  string
	SchoolName  string
	School      *School      `orm:"rel(fk)"`
	Friends     []*Friend    `orm:"rel(m2m)"`
	Matches     []*Match     `orm:"rel(m2m)"`
	StudentRank *StudentRank `orm:"null;rel(one);on_delete(set_null)"`
	//match requests that send by this student
	RequestedMatch string
	//match request that send to this student
	MatchRequests string
	//friend request that send to this student
	FriendRequests string `orm:"size(500)"`
	//friend requests that send by this student
	RequestedFriends        string `orm:"size(500)"`
	Grade                   string
	Field                   string
	EmailAdress             string
	MobileNumber            string `orm:"unique"`
	AvatarCode              string
	Coin                    int
	LastDailyCoinTimeString string
	TotalScore              int
	WeekScore               int
	Token                   string `orm:"size(500)"`
}

type StudentRank struct {
	Id                  int
	Student             *Student `orm:"reverse(one)"`
	WonMatches          int
	TotalMatches        int
	WeekWonMatches      int
	WeekTotalMatches    int
	TotalScore          int
	WeekScore           int
	TotalRank           int
	WeekRank            int
	SchoolTotalRank     int
	SchoolWeekRank      int
	WeekRankArray       string        `orm:"size(500)"`
	SchoolWeekRankArray string        `orm:"size(500)"`
	LessonRanks         []*LessonRank `orm:"reverse(many)"`
}

type LessonRank struct {
	Id                 int
	StudentId          int
	Subject            string
	TotalRightAnswers  int
	TotalQuestions     int
	WeekRightAnswers   int
	WeekTotalQuestions int
	TotalScore         int
	WeekScore          int
	TotalRank          int
	WeekRank           int
	WeekRankArray      string       `orm:"size(500)"`
	WeekPercentsArray  string       `orm:"size(500)"`
	StudentRank        *StudentRank `orm:"rel(fk)"`
}

type Friend struct {
	Id           int
	Students     []*Student `orm:"reverse(many)"`
	Name         string
	FamilyName   string
	AvatarCode   string
	Grade        string
	Field        string
	SchoolName   string
	MobileNumber string `orm:"unique"`
	StudentId    int
	//0 means first send friend request to second, 1 means second sent friend request to first
	//2 means friend
	//	FriendshipState int
}

type Match struct {
	Id                 int
	Students           []*Student `orm:"reverse(many)"`
	FirstOpName        string
	SecondOpName       string
	FirstId            int
	SecondId           int
	FirstTotalAnswers  int
	SecondTotalAnswers int
	//0 means need oponent, 1 means started but not finished yet, 2 means finished
	State int
	//PlayedId           int
	TurnId             int
	TurnQuestionNumber int
	ShowInClient1      bool
	ShowInClient2      bool
	FirstAvatarCode    string
	SecondAvatarCode   string
	SubMatches         []*SubMatch `orm:"reverse(many)"`
	Time               time.Time   `orm:"auto_now_add;type(datetime)"`
	IsLocked           bool
	LockTime           time.Time `orm:"auto_now_add;type(datetime)"`
}

type SubMatch struct {
	Id       int
	Date     time.Time `orm:"auto_now_add;type(datetime)"`
	Subject  string
	PlayedId int
	FirstId  int
	SecondId int
	//0 means true, 1 means false
	FirstAnswers     string
	SecondAnswers    string
	FirstName        string
	FirstFamilyName  string
	SecondName       string
	SecondFamilyName string
	//0 means not started, 1 waiting for oponent to answer the questions, 2 means finished
	//1 means running , 2 means finished
	State int
	//	Rank      *Rank       `orm:"rel(fk)"`
	Match     *Match      `orm:"rel(fk)"`  //many to one
	Questions []*Question `orm:"rel(m2m)"` //it has to be many to one
}

type Question struct {
	Id       int
	Question string `orm:"size(500)"`
	Answer1  string
	Answer2  string
	Answer3  string
	Answer4  string
	//an integer between 0 to 3
	RightAnswer  int
	Description  string `orm:"size(500)"`
	Grade        int
	Field        string
	Subject      string
	Chapter      int
	SubMatces    []*SubMatch `orm:"reverse(many)"`
	AnswerSource string
	Source       string
	Page         int
}

func init() {
	orm.RegisterModel(new(Student), new(School), new(Match), new(Question),
		new(Friend), new(SubMatch), new(LessonRank), new(StudentRank))
}
