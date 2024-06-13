package main

import (
	"atkit/model"
	"atkit/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//初期化
	Init()

	model.Init()

	err := model.CreateUser("admin", "admin", "admin")

	if err != nil {
		log.Fatalln(err)
	}

	err = model.CreateUser("user", "user", "user")

	if err != nil {
		log.Fatalln(err)
	}

	//グループ作成
	groupid,err := model.CreateGroup("test", "admin")

	if err != nil {
		log.Fatalln(err)
	}

	//グループ取得
	group, err := model.Get_Group(groupid)

	if err != nil {
		log.Fatalln(err)
	}

	//メンバーを追加
	err = group.Add_Member(model.GroupMember{
		MemberID: util.GenID(),
		UserID:   "user",
		Role:     "admin",
	})

	if err != nil {
		log.Fatalln(err)
	}

	//メンバーを取得
	members, err := group.Get_Members()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(members)

	//グループ2を作成
	groupid2,err := model.CreateGroup("test2", "admin")

	if err != nil {
		log.Fatalln(err)
	}

	//メンバーを追加
	err = group.Add_Member(model.GroupMember{
		MemberID: util.GenID(),
		UserID:   "user",
		Role:     "admin",
	})

	if err != nil {
		log.Fatalln(err)
	}

	//グループ2を取得
	group2, err := model.Get_Group(groupid2)

	if err != nil {
		log.Fatalln(err)
	}

	//メンバーを取得
	members2, err := group2.Get_Members()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(members2)

	log.Println("finish")
}

func ServerMain() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
