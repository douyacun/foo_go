package commands

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"os"
)

var UniqCommand = cli.Command{
	Name: "uniq",
	Usage: "统计文件重复的行数",
	Action: UniqAction,
}

func UniqAction(cli *cli.Context)  {
	counts := make(map[string]int)
	if len(cli.Args()) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, f := range cli.Args() {
			fp, err := os.Open(f)
			if err != nil {
				fmt.Println(err)
				continue
			}
			countLines(fp, counts)
		}
	}

	for line, n := range counts {
		if n > 0{
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(fp *os.File, counts map[string]int)  {
	defer fp.Close()
	input := bufio.NewScanner(fp)
	for input.Scan() {
		counts[input.Text()]++
	}
}