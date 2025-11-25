package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetVotes() []Vote {
	ret := make([]Vote, 0)
	if err := Conn.Table("vote").Find(&ret).Error; err != nil {
		fmt.Printf("查询失败, err:%s\n", err.Error())
	}
	return ret
}

func GetVote(id int64) VoteWithOpts {
	var ret Vote
	if err := Conn.Table("vote").Where("id = ?", id).First(&ret).Error; err != nil {
		fmt.Printf("查询失败, err:%s\n", err.Error())
	}

	opts := make([]VoteOpt, 0)
	if err := Conn.Table("vote_opt").Where("vote_id = ?", id).Find(&opts).Error; err != nil { //Mysql本身比较脆弱，所以不加外键，利用代码关联
		fmt.Printf("查询选项失败, err:%s\n", err.Error())
	}
	return VoteWithOpts{
		Vote: ret,
		Opts: opts,
	}
}

func DoVote(userId int64, voteId int64, optIds []int64) bool {
	// 记录用户投票行为
	var ret Vote
	if err := Conn.Table("vote").Where("id = ?", userId).First(&ret).Error; err != nil {
		fmt.Printf("查询失败, err:%s\n", err.Error())
	}

	// 更新选项投票数
	for _, value := range optIds {
		if err := Conn.Table("vote_opt").Where("id = ?", value).UpdateColumn("count", gorm.Expr("count + ?", 1)).Error; err != nil {
			fmt.Printf("更新选项投票数失败, err:%s\n", err.Error())
			return false
		}
		user := VoteOptUser{
			UserId:     userId,
			VoteId:     voteId,
			VoteOptId:  value,
			CreateTime: time.Now(),
		}
		_ = Conn.Create(&user).Error // 记录用户投票选项
	}
	return true
}
