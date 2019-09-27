package models

//***************************** HELP FOR QUIERIES**********************************
//https://beego.me/docs/mvc/model/query.md#load-related-field
//*********************************************************************************

//*********************************ADDING USER*************************************
//func AddUser() {
//	school := School{Name: "salam", City: "tehran", Region: 10, Rate: 1}
//	//	fmt.Println(school)
//	_, err := orm.NewOrm().Insert(&school)
//	if err != nil {
//		fmt.Printf("save err... %s", err)
//	}

//	std := Student{Name: "vahid", FamilyName: "johney", School: getSchool("salam"), Grade: 10,
//		Field: "math", EmailAdress: "email", MobileNumber: "0546", Token: TokenGenerator()}
//	_, er := orm.NewOrm().Insert(&std)
//	if er != nil {
//		fmt.Printf("save err... %s", er)
//	}
//	//	fmt.Println(std)
//}
//*********************************************************************************

//************** CALLING SCHOOL OF ONE STUDENT AND VICE VERSA**********************
//	std := Student{Id: 1}
//	orm.NewOrm().Read(&std)
//	orm.NewOrm().LoadRelated(&std, "School")
//	fmt.Println("school is", std.School)

//	school := School{Id: 1}
//	orm.NewOrm().Read(&school)
//	orm.NewOrm().LoadRelated(&school, "Students")
//	fmt.Println("students is", school.Students[0])
//*********************************************************************************

//*******************************USAGE OF QUERIES**********************************
//o := orm.NewOrm()
//var questions []*Question
//qs := o.QueryTable("question").Filter("match_id__in", matchId).RelatedSel()
//qs.All(&questions)
//this will give the questions of specific match
//*********************************************************************************
