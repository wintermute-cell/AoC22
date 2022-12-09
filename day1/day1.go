package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func addTop(toplist [3]int, newEntry int) [3]int {
    for i, entry := range toplist {
        if newEntry > entry {
            for j := 2; j > i; j-- {
                toplist[j] = toplist[j-1]
            }
            toplist[i] = newEntry
            fmt.Println(toplist[i])
            break
        }
    }
    return toplist
}

func main() {
    readFile, err := os.Open("input")
    if err != nil {
        fmt.Println(err)
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    //largestCalories := -1
    var lCalTopThree = [3]int{-1,-1,-1}
    currSum := 0
    for _, line := range fileLines {
        if line != "" {
            number, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println(err)
            }
            currSum += number
        } else {
            if currSum > lCalTopThree[2] {
                lCalTopThree = addTop(lCalTopThree, currSum)
            }
            currSum = 0
        }
    }

    fmt.Println(lCalTopThree)

    total := 0
    for _, num := range lCalTopThree { 
        total += num
    }
    fmt.Println(total)
}

