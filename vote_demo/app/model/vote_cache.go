package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func GetVoteCachae(ctx context.Context, id int64) VoteWithOpt {
	var ret VoteWithOpt
	key := fmt.Sprintf("vote_cache_%d", id)
	voteStr, err := Rdb.Get(ctx, key).Result()
	if err == nil || len(voteStr) > 0 {
		_ = json.Unmarshal([]byte(voteStr), &ret)
		return ret
	}
	fmt.Printf("err:%s", err.Error())
	ret = GetVote(id)
	if ret.Vote.Id > 0 {
		s, _ := json.Marshal(ret)
		err1 := Rdb.Set(ctx, key, s, 600*time.Second).Err()
		if err1 != nil {
			fmt.Printf("set cache err:%s", err1.Error())
		}
	}
	return ret

}
