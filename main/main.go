package main

import (
	"flag"
)

var a = flag.Bool("all", false, "获取全部投资历史记录")
var o = flag.Bool("options", false, "获取全部options历史记录")
var thisweek = flag.Bool("this week", false, "获取本周操作记录")
var update_thisweek= flag.Bool("update this week", false, "获取本周操作记录")
var nextweek =flag.Bool("next week", false, "获取下周操作记录")
var update_nextweek= flag.Bool("update next week", false, "获取本周操作记录")

func main() {
	flag.Parse()
	if *a {
		return
	}
	if *o {
		return
	}
	if *thisweek {
		return
	}
	if *nextweek {
		return
	}
	if *update_thisweek {
		return
	}
	if *update_nextweek {
		return
	}

}
