package schedules

import (
	"fmt"
	"time"

	"github.com/vote_demo/app/model"
)

func Strat() {
	go EndVote()

}

func EndVote() {
	t := time.NewTicker(5 * time.Second)
	defer func() {
		t.Stop()
	}()

	for {
		select {
		case <-t.C:
			fmt.Println("定时器启动...")
			// 执行定时任务逻辑
			model.EndVote()
			fmt.Println("定时器结束...")
		}
	}
}
