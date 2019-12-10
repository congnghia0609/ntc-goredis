/**
 *
 * @author nghiatc
 * @since Feb 8, 2018
 */

package main

import (
	"fmt"
	"github.com/congnghia0609/ntc-gconf/nconf"
	"ntc-goredis/nredis"
	"path/filepath"
	"runtime"
)

func GetWDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}

func InitNConf() {
	wdir := GetWDir()
	fmt.Println("wdir:", wdir)
	nconf.Init(wdir)
}

func main() {
	// Init NConf
	InitNConf()

	// Init NRedis
	nredis.InitPoolConf("ruser")
	nredis.ExampleNRedis()

	//nredis.ExampleNJson()

}
