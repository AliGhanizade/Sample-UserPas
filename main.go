package main

import (
	"fmt"
	"userpas/routes"

	pt "github.com/yaa110/go-persian-calendar"

	"github.com/carmo-evan/strtotime"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {

	// fmt.Println(pt.Yesterday().Weekday()) // output: شنبه
	// sdfsd := pt.Now().LastWeekday().Format("yyyy-MM-w-dd")
	// test := pt.Now().Format("yyyy-MM-w-dd")
	// var t time.Time = time.Date(2016, time.January, 1, 12, 1, 1, 0, pt.Iran())
	// fmt.Println(sdfsd, "\n", test, "\n", t)

	// now := time.Now()
	// // محاسبه دیروز
	// yesterday := now.Add(-24 * time.Hour)
	// lastweek := now.AddDate(0, 0, -7)
	// yesterdayAtEight := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 8, 0, 0, 0, yesterday.Location())
	// fmt.Println("دیروز ساعت 08:00:", yesterdayAtEight)
	// test1 := pt.New(yesterdayAtEight).Format("yyyy-MM-w-dd")
	// test2 := pt.New(yesterday).Format("yyyy-MM-w-dd")
	// test3 := pt.New(lastweek).Format("yyyy-MM-w-dd")
	// fmt.Println(test1, "\n", test2, "\n", test3)

	u, err := strtotime.Parse("last week 8am", pt.Now().Unix())

	fmt.Println(u, "\n", err, "\n")

	// config.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }

	// config.DB.AutoMigrate(&model.User{})
	r := routes.SetupRouter()

	//running
	r.Run(":8443")
}
