package services

import (
	. "awesomeProject/lj/domain"
	. "awesomeProject/lj/init"
	"github.com/gin-gonic/gin"
	"strconv"
)

//查询单个用户
//func Query(c *gin.Context) {
//	ms := []MemberList{}
//
//	id := c.Query("id")
//
//	atoi, err := strconv.Atoi(id)
//	if err != nil {
//		panic(err)
//	}
//
//	up := DownToUp(atoi, ms)
//	c.JSON(200, up)
//
//}

func Query(c *gin.Context) {
	ms := []MemberList{}

	id := c.Query("id")

	atoi, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, DownToUp(atoi, &ms))

}

func DownToUp(id int, members *[]MemberList) []MemberList {
	member := MemberList{}

	Db.Where("member_list_id = ?", id).First(&member)

	if member.TUserid != 0 {
		*members = append(*members, member)
		DownToUp(member.TUserid, members)
	}

	return *members

}
