package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "AcceptFriendship",
            Router: `/acceptfriendship`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "AcceptMatch",
            Router: `/acceptmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "AcceptMatchRequest",
            Router: `/acceptmatchrequest`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddStudent",
            Router: `/addstudent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddSubMatch",
            Router: `/addsubmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "ChangeAvatar",
            Router: `/changeavatar`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "FriendRequest",
            Router: `/friendrequest`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetFriend",
            Router: `/getfriend`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetMatch",
            Router: `/getmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetMyFriends",
            Router: `/getmyfriends`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetMySchool",
            Router: `/getmyschool`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSchool",
            Router: `/getschool`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSchools",
            Router: `/getschools`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetStudent",
            Router: `/getstudent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetSubMatch",
            Router: `/getsubmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "IsMatch",
            Router: `/ismatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "RequestMatch",
            Router: `/requestmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "StartMatch",
            Router: `/startmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "StartMatchForFriendReq",
            Router: `/startmatchforfriendreq`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateStudent",
            Router: `/updatestudent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/najidroid/quizOfStudents/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateSubMatch",
            Router: `/updatesubmatch`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
