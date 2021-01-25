package services

import (
	"awesomeProject/app/model"
	. "awesomeProject/init"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func Test(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {

		dst := filepath.Join("./upload", strconv.Itoa(int(time.Now().Unix()))+file.Filename)

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

// @Summary 查找推荐关系
// @Description   根据t_userid一直找到公司
// @Param id query int true "id"
// @Success 200 {object} domain.MemberList
// @Router /query [get]
func FindRecommended(c *gin.Context) {

	var ms []model.MemberList
	var m model.MemberList

	if err := c.ShouldBind(&m); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	c.JSON(200, DownToUp(m.MemberListID, &ms))

}

// @Summary 查看团队
// @Description   根据自己的role，查找团队成员
// @Param id query int true "id"
// @Success 200 {object} domain.MemberList
// @Router /team [get]
func Team(c *gin.Context) {
	m := model.MemberList{}

	var ms []model.MemberList

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

func DownToUp(id int, members *[]model.MemberList) []model.MemberList {
	member := model.MemberList{}

	Db.Where("member_list_id = ?", id).First(&member)

	if member.TUserid != 0 {
		*members = append(*members, member)
		DownToUp(member.TUserid, members)
	}

	return *members

}

// @Summary 查找关系链
// @Description   根据提交的字符串，来查找关系链，比如: 15,14,13 查出来的关系就是特约，省级，市级，但是市级的下一级没有指定的话，就是找出所有可能
// @Description  最短可查 2 层关系，最长可查 10 层关系
// @Description  参数 例 : 12,12     以第一个开始
// @Param id formData  string true "id"
// @Success 200 {object} domain.MemberList
// @Router /link [post]
func Link(c *gin.Context) {

	st := c.PostForm("id")
	m := &[]model.MemberList{}

	split := strings.Split(st, ",")

	fmt.Println(len(split))
	var count int64

	switch len(split) {

	case 2:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" ")).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 3:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+""))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 4:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+"")))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 5:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+""))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 6:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+"")))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 7:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+""))))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 8:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+" and t_userid in (?)",
									Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[7]+"")))))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}
	case 9:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+" and t_userid in (?)",
									Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[7]+" and t_userid in (?)",
										Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[8]+""))))))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}

	case 10:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+" and t_userid in (?)",
									Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[7]+" and t_userid in (?)",
										Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[8]+" and t_userid in (?)",
											Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[9]+"")))))))))).Find(&m).Count(&count)

		if m != nil {
			c.JSON(http.StatusOK, gin.H{
				"Count": count,
				"Data":  m,
			})
		}
	case 1:
		c.JSON(http.StatusBadRequest, "自己去查数据库...")

	default:
		c.JSON(http.StatusBadRequest, "所求关系链太长....")

	}

}
