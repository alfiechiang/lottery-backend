package six_lottery

import (
	"lottery/model"
	"lottery/serializer"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type CollyRequest struct {
}

func (request *CollyRequest) Colly(c *gin.Context) serializer.Response {

	collect := colly.NewCollector()
	collect.OnHTML("tr[style]", func(e *colly.HTMLElement) {

		styleStr := e.Attr("style")
		if strings.Contains(styleStr, "center") {
			six := model.SixLotteryList{}
			smash1 := strings.Split(e.Text, ")")

			six.Date = getDate(smash1[0])
			num := smash1[1]
			smash2 := strings.Split(num, " ")

			for _, v := range smash2 {
				if strings.Contains(v, ",") {
					v =strings.Replace(v,"\u00a0","",-1)
					v= strings.TrimSpace(v)
					six.SixNum = v
				}
				if strings.Contains(v, "\n\n") {
					smash3 := strings.Split(v, "\n\n")
					i, _ := strconv.Atoi(smash3[0]) // result: i = -18
					six.SpecilNum = i
				}
			}

			var check model.SixLotteryList
			model.DB.Table("six_lottery_list").Where("date", getDateString(smash1[0])).Find(&check)
			if check == (model.SixLotteryList{}) {
				model.DB.Table("six_lottery_list").Create(&six)
			}

		}

	})

	collect.Visit("https://www.pilio.idv.tw/ltohk/ServerC/list.asp")

	return serializer.Response{}

}

func getDateString(smash string) string {
	s1 := strings.Split(smash, "(")
	d1 := strings.Replace(s1[0], "/", "-", 2)
	d2 := strings.Trim(d1, "\n ")
	return d2
}

func getDate(smash string) model.Date {

	s1 := strings.Split(smash, "(")
	d1 := strings.Replace(s1[0], "/", "-", 2)
	d2 := strings.Trim(d1, "\n ")
	date, _ := time.Parse("2006-01-02", d2)
	return model.Date{Time: date}
}
