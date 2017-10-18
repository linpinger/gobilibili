package main

import (
	"github.com/lyyyuna/gobilibili"
	"os"
	"strconv"
	"fmt"
)

func main() {
	bili := gobilibili.NewBiliBiliClient()
//	bili.ConnectServer(1016)
	roomID, _ := strconv.Atoi(os.Args[1])
	fmt.Println("roomid: ", os.Args[1])
	bili.ConnectServer(roomID)
}
