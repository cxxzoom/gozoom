package main

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

type payload struct {
	UserId    int      `json:"user_id"`
	MeasureId int      `json:"measureId"`
	IsSave    bool     `json:"isSave"`
	Access    access   `json:"systemAccess"`
	Inputs    struct{} `json:"inputs"`
	Answer    []answer `json:"answer"`
}

type access struct {
	accessToken string
}

type answer struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

func main() {
	gauges, err := parseGaugeJson()
	if err != nil {
		log.Panicln(err)
	}

	gaugeInfo, err := getGaugeInfo(*gauges)

	if err != nil {
		log.Panicln(err)
	}

	users, err := parseUserJson()
	if err != nil {
		log.Panicln(err)
	}
	userInfo, err := getUserInfo(users)
	if err != nil {
		log.Panicln(err)
	}

	// TODO 还需要获取用户的数据
	//res := payload{}

	for _, res := range *gaugeInfo {
		for _, opts := range *userInfo {
			tmpAnswer := genAnswer(res.Options)
			_, err := commit(opts.Account, &payload{
				UserId:    opts.Id,
				MeasureId: res.Id,
				IsSave:    true,
				Access:    access{accessToken: "md5"},
				Inputs:    struct{}{},
				Answer:    *tmpAnswer,
			})
			if err != nil {
				log.Panicln(err)
			}
		}

	}
	time.Sleep(20 * time.Second)
}

// genAnswer 构造数据
func genAnswer(options []gaugeOpts) *[]answer {
	var res []answer
	for _, opts := range options {
		tmp := strings.Split(opts.answerStr, " ")

		n := strings.Split(tmp[rand.Intn(len(tmp)-2)], "、")
		res = append(res, answer{
			Id:     opts.questionId,
			Answer: n[0],
		})
	}

	return &res
}
