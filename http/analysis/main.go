package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

// 第一个参数是传入的，第二个是默认的
var cpuprofile = flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		// 创建文件
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		// 进行性能分析
		pprof.StartCPUProfile(f)
		// 方法结束退出分析
		defer pprof.StopCPUProfile()
		var result int
		for i := 0; i < 10000000000; i++ {
			result += i
		}
		log.Println("result:", result)
	}
}
