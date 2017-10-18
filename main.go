package main

import (
	"github.com/linpinger/gobilibili/bilibili"
	"os"
	"strconv"
	"fmt"
)

func main() {
	roomID, _ := strconv.Atoi(os.Args[1])
	fmt.Println("roomid: ", os.Args[1])

	bili := bilibili.NewBiliBiliClient()
//	bili.ConnectServer(1016)
	bili.ConnectServer(roomID)
}
