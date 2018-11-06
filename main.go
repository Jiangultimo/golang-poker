package main
import (
	"log"
	"os"
	"matchers"
	"search"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}