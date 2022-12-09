package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() []string {
    // read the input file
    readFile, _ := os.Open("input")
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
    readFile.Close()
    return fileLines
}

func determineOutcome(enemyPick int, myPick int) int {
    if (myPick-23) == enemyPick { //draw
        return 3
    }
    diff := myPick - enemyPick
    if diff == 24 || diff == 21 { //win
        return 6
    }
    return 0 //loose
}
func partOne(fileLines []string) {
    totalPoints := 0
    for _, line := range fileLines {
        elems := strings.Split(line, " ")
        enemyPick := int([]rune(elems[0])[0])
        myPick := int([]rune(elems[1])[0])

        pickBonus := myPick - 87
        outcomePoints := determineOutcome(enemyPick, myPick)
        totalPoints += (pickBonus+outcomePoints)
    }
    fmt.Println("Part 1: ", totalPoints)
}

func partTwo(fileLines []string) {
    totalPoints := 0
    for _, line := range fileLines {
        elems := strings.Split(line, " ")
        enemyPick := int([]rune(elems[0])[0]) -65
        needOutcome := int([]rune(elems[1])[0]) -88
        var myPick int
        switch needOutcome {
            case 0:
                myPick = ((enemyPick-1)+3)%3
                totalPoints+=0
            case 1:
                myPick = enemyPick
                totalPoints+=3
            case 2:
                myPick = (enemyPick+1)%3
                totalPoints+=6
        }
        fmt.Println(myPick, enemyPick, elems[1], needOutcome)
        pickBonus := myPick + 1
        totalPoints += pickBonus
    }
    fmt.Println("Part 2: ", totalPoints)
}

func main() {

    fileLines := readInput()
    partOne(fileLines)
    partTwo(fileLines)
}
