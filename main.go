package main

import (
	"github.com/linpinger/gobilibili/bilibili"
	"os"
	"strconv"
	"fmt"
)

func main() {
	if 2 == len(os.Args) {
		roomID, _ := strconv.Atoi(os.Args[1])
		fmt.Println("roomid: ", os.Args[1])

		bili := bilibili.NewBiliBiliClient()
		bili.ConnectServer(roomID)
//		bili.ConnectServer(1016)
	} else {
		fmt.Printf("Usage:\n\t%s [ roomID ]\n\n", os.Args[0])
		os.Exit(0)
	}
}
