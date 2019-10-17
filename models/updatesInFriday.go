package models

import (
	"fmt"

	"strconv"

	"github.com/astaxie/beego/orm"
)

func updateRankOrdersInFriday() {
	fmt.Println("***************UPDATING RANK ORDERS For Friday***************")
	updateSchoolRankOrdersInFriday()
	updateStudentRankOrdersInFriday()
	updateLessonRankOrdersInFriday("ادبیات")
	updateLessonRankOrdersInFriday("عربی")
	updateLessonRankOrdersInFriday("دینی")
	updateLessonRankOrdersInFriday("زبان")
	updateLessonRankOrdersInFriday("ریاضی")
	updateLessonRankOrdersInFriday("فیزیک")
	updateLessonRankOrdersInFriday("شیمی")
}

func updateLessonRankOrdersInFriday(subject string) {
	fmt.Println("***************UPDATING LESSON RANK ORDERS For Friday***************")

	o := orm.NewOrm()
	var lessonsForWeek []*LessonRank
	o.QueryTable("LessonRank").Filter("Subject", subject).OrderBy("-WeekScore").RelatedSel().All(&lessonsForWeek)
	if len(lessonsForWeek) == 0 {
		return
	}
	max := 100000000
	rank := 0
	for _, item := range lessonsForWeek {
		fmt.Println("max:", max, " week score:", item.WeekScore)
		if item.WeekScore < max {
			rank++
			max = item.WeekScore
		}
		item.WeekRank = rank
		o.Update(item)
	}
	var lessonsForTotal []*LessonRank
	o.QueryTable("LessonRank").Filter("Subject", subject).OrderBy("-TotalScore").RelatedSel().All(&lessonsForTotal)
	max = 100000000
	rank = 0
	for _, item := range lessonsForTotal {
		if item.TotalScore < max {
			rank++
			max = item.TotalScore
		}
		item.TotalRank = rank
		o.Update(item)
	}
}

func updateSchoolRankOrdersInFriday() {
	fmt.Println("***************UPDATING SCHOOL RANK ORDERS For Friday***************")

	o := orm.NewOrm()
	var schoolsForWeekRank []*School
	o.QueryTable("School").OrderBy("-WeekScore").RelatedSel().All(&schoolsForWeekRank)
	max := 100000000
	rank := 0
	for _, item := range schoolsForWeekRank {
		if item.WeekScore < max {
			rank++
			max = item.WeekScore
		}
		item.WeekRank = rank
		o.Update(item)
	}

	var schoolsForTotalRank []*School
	o.QueryTable("School").OrderBy("-TotalScore").RelatedSel().All(&schoolsForTotalRank)
	max = 100000000
	rank = 0
	for _, item := range schoolsForTotalRank {
		if item.TotalScore < max {
			rank++
			max = item.TotalScore
		}
		item.TotalRank = rank
		o.Update(item)
	}
}

func updateStudentRankOrdersInFriday() {
	fmt.Println("***************UPDATING STUDENT RANK ORDERS For Friday***************")

	o := orm.NewOrm()
	var studentForWeekRank []*StudentRank
	o.QueryTable("StudentRank").OrderBy("-WeekScore").RelatedSel().All(&studentForWeekRank)
	max := 100000000
	rank := 0
	for _, item := range studentForWeekRank {
		if item.WeekScore < max {
			rank++
			max = item.WeekScore
		}
		item.WeekRank = rank
		o.Update(item)
	}

	var studentForTotalRank []*StudentRank
	o.QueryTable("StudentRank").OrderBy("-TotalScore").RelatedSel().All(&studentForTotalRank)
	max = 100000000
	rank = 0
	for _, item := range studentForTotalRank {
		if item.TotalScore < max {
			rank++
			max = item.TotalScore
		}
		item.TotalRank = rank
		o.Update(item)
	}
}

func UpdateRanksInFriday() {
	fmt.Println("***************UPDATING RANKS For Friday***************")
	o := orm.NewOrm()
	var ranks []*StudentRank
	o.QueryTable("StudentRank").All(&ranks)

	for _, item := range ranks {
		o.LoadRelated(item, "LessonRanks")
		lsnRanks := item.LessonRanks
		for _, itm := range lsnRanks {
			updateLessonRankInFriday(itm)
		}
		updateStudentRankInFriday(item)
	}

	var schools []*School
	o.QueryTable("School").All(&schools)
	for _, item := range schools {
		updateSchoolInFriday(item)
	}
}

func updateSchoolInFriday(school *School) {
	school.WeekWonMatches = 0
	school.WeekTotalMatches = 0
	school.WeekEvenMatches = 0
	school.WeekRankArray += "-" + strconv.Itoa(school.WeekRank) + "-"
	orm.NewOrm().Update(school)
}

func updateStudentRankInFriday(stdRank *StudentRank) {
	stdRank.WeekRankArray += "-" + strconv.Itoa(stdRank.WeekRank) + "-"
	stdRank.SchoolWeekRankArray += "-" + strconv.Itoa(stdRank.SchoolWeekRank) + "-"
	stdRank.WeekWonMatches = 0
	stdRank.WeekTotalMatches = 0
	stdRank.WeekScore = 0
	fmt.Println("student rank:", stdRank)
	orm.NewOrm().Update(stdRank)
}

func updateLessonRankInFriday(lessonRank *LessonRank) {
	if lessonRank.WeekTotalQuestions != 0 {
		lessonRank.WeekPercentsArray += "-" + strconv.Itoa(100*lessonRank.WeekRightAnswers/lessonRank.WeekTotalQuestions) + "-"
	} else {
		lessonRank.WeekPercentsArray += "-0-"
	}
	lessonRank.WeekRankArray += "-" + strconv.Itoa(lessonRank.WeekRank) + "-"
	lessonRank.WeekRightAnswers = 0
	lessonRank.WeekTotalQuestions = 0
	lessonRank.WeekScore = 0
	fmt.Println("lesson rank:", lessonRank)
	orm.NewOrm().Update(lessonRank)
}
