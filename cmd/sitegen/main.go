package main

import (
	"github.com/fulldump/goconfig"
	"github.com/fulldump/holadoc"
)

func main() {

	c := holadoc.Config{}
	goconfig.Read(&c)

	holadoc.HolaDoc(c)
}
