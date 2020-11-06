package main

import (
	. "awesomeProject/domain"
	. "awesomeProject/init"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	SqlInstance()
	m := &[]MemberList{}

	st := "15,15,15,15,15,15,15"

	split := strings.Split(st, ",")

	fmt.Println(len(split))

	switch len(split) {

	case 2:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" ")).Find(&m)
		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 3:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+""))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 4:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+"")))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 5:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+""))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 6:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",

							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+"")))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 7:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+""))))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	case 8:
		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+" and t_userid in (?)",
									Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[7]+"")))))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))
	case 9:

		Db.Where("role = "+split[0]+" and t_userid in (?)",
			Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[1]+" and t_userid in (?)",
				Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[2]+" and t_userid in (?)",
					Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[3]+" and t_userid in (?)",
						Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[4]+" and t_userid in (?)",
							Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[5]+" and t_userid in (?)",
								Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[6]+" and t_userid in (?)",
									Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[7]+" and t_userid in (?)",
										Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[8]+""))))))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

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
											Db.Table("wd_member_list").Select("member_list_id").Where("role = "+split[9]+"")))))))))).Find(&m)

		marshal, _ := json.MarshalIndent(m, " ", " ")

		fmt.Println(string(marshal))

	default:

	}

}

func insertSort(list []int) {
	len := len(list)

	for i := 1; i < len; i++ {
		right := list[i]
		leffindex := i - 1
		if list[leffindex] > right {
			for ; leffindex >= 0 && list[leffindex] > right; leffindex-- {
				list[leffindex+1] = list[leffindex]
			}
			list[leffindex+1] = right
		}

	}

}

func ShellSort(list []int) {

	for step := len(list) / 2; step >= 1; step /= 2 {
		for i := step; i < len(list); i++ {
			for x := i - step; x >= 0 && list[x+step] < list[x]; x -= step {
				list[x], list[x+step] = list[x+step], list[x]

			}

		}

	}

}
