package main

import (
	"github.com/xiaoyuanzhu-com/onwards/server"
)

func main() {
	srv := server.NewServer()
	srv.Start()
}
