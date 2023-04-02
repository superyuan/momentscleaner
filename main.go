package main

import (
	"github.com/superyuan/momentscleaner/cleaner"
	"github.com/superyuan/momentscleaner/logconfig"
	_ "net/http/pprof"
)

func main() {
	//go func() {
	//	fmt.Println(http.ListenAndServe("localhost:8888", nil))
	//}()

	logconfig.InitLogrus("cleaner", 10)
	cleaner.DoClean()
}
