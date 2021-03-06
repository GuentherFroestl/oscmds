package main


import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"regexp"
)

var numRegEx = regexp.MustCompile("\\d+")

func main() {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("'date' cmd:\n%s\n", out)
	}


	out,err = exec.Command("ls", "-1").Output()
	fmt.Printf("'ls -1' cmd:\n%s\n", out)
	if err != nil {
		log.Fatal(err)
	}else{
		fmt.Printf("'ls -1' cmd:\n%s\n", out)
	}
	out,err = exec.Command("sh", "./countFiles.sh").Output()

	if err != nil {
		log.Fatal(err)
	}else{
		sNmbStr := string(numRegEx.Find(out))
		fmt.Printf("'ls -1 | wc -l' response:\n'%v'\n", sNmbStr)
		nmb,_ := strconv.ParseFloat(sNmbStr, 64)
		fmt.Printf("'number of files in dir (float64)':\n%v\n", nmb)
	}

}
