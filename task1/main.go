package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numOfDataCenters := 0
	numOfServers := 0
	numOfEvents := 0

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	_, err := fmt.Fscanln(in, &numOfDataCenters, &numOfServers, &numOfEvents)
	if err != nil {
		return
	}

	if numOfDataCenters <= 0 || numOfServers <= 0 || numOfEvents <= 0 {
		return
	}
	if numOfDataCenters*numOfServers > 1000000 {
		return
	}
	if numOfEvents > 100000 {
		return
	}

	dataCenterServerDownCount := make(map[int]int, numOfDataCenters)
	dataCenterResetCount := make(map[int]int, numOfDataCenters)
	dataCenterServerDisableIndicator := make(map[int][]bool, numOfDataCenters)
	for i := 0; i < numOfDataCenters; i++ {
		dataCenterServerDisableIndicator[i] = make([]bool, numOfServers, numOfServers)
	}

	var message string
	for i := 0; i < numOfEvents; i++ {
		message, _ = in.ReadString('\n')

		splitMessage := strings.Fields(message)

		switch splitMessage[0] {
		case "RESET":
			dataCenterId, _ := strconv.Atoi(splitMessage[1])
			dataCenterResetCount[dataCenterId-1]++
			dataCenterServerDownCount[dataCenterId-1] = 0
			dataCenterServerDisableIndicator[dataCenterId-1] = make([]bool, numOfServers, numOfServers)
		case "DISABLE":
			dataCenterId, _ := strconv.Atoi(splitMessage[1])
			dataServerId, _ := strconv.Atoi(splitMessage[2])
			if dataCenterServerDisableIndicator[dataCenterId-1][dataServerId-1] == false {
				dataCenterServerDisableIndicator[dataCenterId-1][dataServerId-1] = true
				dataCenterServerDownCount[dataCenterId-1]++
			}
		case "GETMAX":
			max := 0
			dataCenterMaxNum := 0
			for i := 0; i < numOfDataCenters; i++ {
				ra := dataCenterResetCount[i] * (numOfServers - dataCenterServerDownCount[i])
				if ra > max {
					max = ra
					dataCenterMaxNum = i
				}
			}
			fmt.Fprintln(out, dataCenterMaxNum+1)
		case "GETMIN":
			min := dataCenterResetCount[0] * (numOfServers - dataCenterServerDownCount[0])
			dataCenterMinNum := 0
			for i := 0; i < numOfDataCenters; i++ {
				ra := dataCenterResetCount[i] * (numOfServers - dataCenterServerDownCount[i])
				if ra < min {
					min = ra
					dataCenterMinNum = i
				}
			}
			fmt.Fprintln(out, dataCenterMinNum+1)
		}
	}
}

//3 3 12
//DISABLE 1 2
//DISABLE 2 1
//DISABLE 3 3
//GETMAX
//RESET 1
//RESET 2
//DISABLE 1 2
//DISABLE 1 3
//DISABLE 2 2
//GETMAX
//RESET 3
//GETMIN
//
//1
//2
//1

//2 3 9
//DISABLE 1 1
//DISABLE 2 2
//RESET 2
//DISABLE 2 1
//DISABLE 2 3
//RESET 1
//GETMAX
//DISABLE 2 1
//GETMIN
