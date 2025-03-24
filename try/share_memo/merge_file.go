package sharememo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	content     = make(chan string, 1000)
	readFileCh  = make(chan struct{}, 3) // flag to mark no more content
	writeFileCh = make(chan struct{}, 0)
)

func readFile(infile string) {
	fin, err := os.Open(infile)
	if err != nil {
		fmt.Println(err)
	}
	defer fin.Close()

	reader := bufio.NewReader(fin)
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			content <- line
		} else {
			if err == io.EOF {
				if len(line) > 0 {
					content <- line + "\n"
				}
				break
			} else {
				fmt.Println(err)
			}
		}
	}
	<-readFileCh
	if len(readFileCh) == 0 {
		close(content)
	}
}

func writeFile(mergedFile string) {
	fout, err := os.OpenFile(mergedFile, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer fout.Close()

	writer := bufio.NewWriter(fout)
	// for {
	// 	line := <-content
	// 	writer.WriteString(line)
	// }
	// if content was closed then exist for loop
	for line := range content {
		writer.WriteString(line)
	}
	writer.Flush()

	writeFileCh <- struct{}{}
}

func main() {
	for i := 0; i < 3; i++ {
		readFileCh <- struct{}{}
	}
	go readFile("data/1.txt")
	go readFile("data/2.txt")
	go readFile("data/3.txt")

	go writeFile("data/big.txt")

	<-writeFileCh
}
