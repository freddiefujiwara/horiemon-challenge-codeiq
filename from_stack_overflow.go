package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type triple struct {
	x, y, z int
}

func triples(l int, n []int, list bool) (nt int, t []triple) {
	num_of_list := len(n)
	for i := 0; i < num_of_list-2; i++ {
		x := n[i]
		if x > l {
			break
		}
		for j := i + 1; j < num_of_list-1; j++ {
			y := x + n[j]
			if y > l {
				break
			}
			pos := sort.SearchInts(n[j:], l-y)
			if j < pos+j && pos+j < num_of_list && n[pos+j] == l-y {
				nt++
			}
		}
	}
	return nt, t
}

func readInput() (l int, n []int, err error) {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == nil && len(line) == 0 {
			err = io.EOF
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, nil, err
		}
		i, err := strconv.Atoi(string(line))
		if err == nil && i < 0 {
			err = errors.New("Nonpositive number: " + line)
		}
		if err != nil {
			return 0, nil, err
		}
		n = append(n, i)
	}

	if len(n) > 0 {
		l = n[0]
		n = n[1:]
	}
	sort.Ints(n)
	for i := 1; i < len(n); i++ {
		if n[i] == n[i-1] {
			copy(n[i:], n[i+1:])
			n = n[:len(n)-1]
		}
	}
	return l, n, nil
}

func main() {
	l, n, err := readInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	list := false
	nt, t := triples(l, n, list)
	fmt.Println(nt)
	if list {
		fmt.Println(t)
	}
}
