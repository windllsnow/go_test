package main 
import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)
var hellolist = []string{
	"hellp world",
	"asdfasfqas",
	"333333333333333",
	"44444444444444",
	"55555555555s",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(hellolist))
	msg,err := hello(index)
	if err != nil{ 
		log.Fatal(err)
	}
	fmt.Println(msg)
}
func hello (index int ) ( string, error ) {
	if index < 0 || index > len(hellolist)-1 {
		return "",errors.New(" out of range:"+ strconv.Itoa(index))
	}
	return hellolist[index],nil
}

