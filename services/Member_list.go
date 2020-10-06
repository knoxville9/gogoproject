package services

import (
	. "awesomeProject/domain"
	. "awesomeProject/init"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 查找推荐关系
// @Description   根据t_userid一直找到公司
// @Param id query int true "ID"
// @Success 200 {object} domain.MemberList
// @Failure 400 {object} domain.APIError "We need ID!!"
// @Failure 404 {object} domain.APIError "Can not find ID"
// @Router /tid [get]
func FindRecommended(c *gin.Context) {
	ms := []MemberList{}

	id := c.Query("id")

	atoi, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, DownToUp(atoi, &ms))

}

// @Summary 查看团队
// @Description   根据自己的role，查找团队成员
// @Param id query int true "ID"
// @Success 200 {object} domain.MemberList
// @Failure 400 {object} domain.APIError "We need ID!!"
// @Failure 404 {object} domain.APIError "Can not find ID"
// @Router /team [get]
func Team(c *gin.Context) {
	m := MemberList{}

	ms := []MemberList{}

	id := c.Query("id")

	Db.Where("member_list_id = ?", id).First(&m)

	role := m.Role

	switch int(role) {
	case 10:
		Db.Where("role10 = ?", m.MemberListID).Find(&ms)
	case 11:
		Db.Where("role11 = ?", m.MemberListID).Find(&ms)
	case 12:
		Db.Where("role12 = ?", m.MemberListID).Find(&ms)
	case 13:
		Db.Where("role13 = ?", m.MemberListID).Find(&ms)
	case 14:
		Db.Where("role14 = ?", m.MemberListID).Find(&ms)
	case 15:
		Db.Where("role15 = ?", m.MemberListID).Find(&ms)
	case 16:
		Db.Where("role16 = ?", m.MemberListID).Find(&ms)
	}

	c.JSON(200, ms)

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
