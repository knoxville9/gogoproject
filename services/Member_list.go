package services

import (
	. "awesomeProject/domain"
	. "awesomeProject/init"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	var m MemberList

	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

}

// @Summary 查找推荐关系
// @Description   根据t_userid一直找到公司
// @Param id query int true "ID"
// @Success 200 {object} domain.MemberList
// @Failure 400 {object} domain.APIError "We need ID!!"
// @Failure 404 {object} domain.APIError "Can not find ID"
// @Router /tid [get]
func FindRecommended(c *gin.Context) {
	var ms []MemberList
	var m MemberList

	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	c.JSON(200, DownToUp(m.MemberListID, &ms))

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

	var ms []MemberList

	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	Db.Where("member_list_id = ?", m.MemberListID).First(&m)

	switch m.Role {
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

	if ms == nil {
		c.JSON(http.StatusBadRequest, "没有任何下级")
		return
	}

	c.JSON(http.StatusOK, ms)

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
