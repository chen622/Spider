package model

import "time"

type Video struct {
	Author      string `jpath:"author"`
	Mid         int    `jpath:"mid"`
	Aid         int    `jpath:"aid"`
	Title       string `jpath:"title"`
	Created     int64  `jpath:"created"`
	Description string `jpath:"description"`
	Pic         string `jpath:"pic"`
}

func (v Video) GetTime() time.Time {
	unix := time.Unix(v.Created, 0)
	return unix
}
