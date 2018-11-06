package main
import (
	"log"
	"os"
	"matchers"
	"search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}