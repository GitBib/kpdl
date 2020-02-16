package main

import (
	"bufio"
	"flag"
	"fmt"
	"kpdl/api"
	_struct "kpdl/api/struct"
	"os/exec"
	"strconv"
	"strings"
)

func DownloadSeries(e int, episode _struct.Episodes, seasonFlag int, itemId int) {
	fileName := "s" + strconv.Itoa(seasonFlag) + "e" + strconv.Itoa(e)
	var urlFile = ""
	var minW = 4045
	for _, file := range episode.Files {
		if file.W < minW {
			minW = file.W
			urlFile = file.Url.Http
		}
	}
	command := "-o " + strconv.Itoa(itemId) + "/" + fileName + ".mp4 " + urlFile
	cmd := exec.Command("youtube-dl", strings.Split(command, " ")...)
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}

func main() {
	ApiKey := flag.String("key", "foo", "Api Key")
	IdItem := flag.Int("item", 0, "id item")
	seasonFlag := flag.Int("s", 0, "id season")
	episodeFlag := flag.Int("e", 0, "id episode")
	flag.Parse()
	fmt.Println("key: ", *ApiKey, " item: ", *IdItem, " s:", *seasonFlag, " e:", *episodeFlag)
	var api = api.Api{ApiKey: *ApiKey}
	seasons := api.GetInfoItem(*IdItem).Seasons
	if -1 == *episodeFlag {
		fmt.Println(seasons)
		for e, episode := range seasons[*seasonFlag].Episodes {
			DownloadSeries(e, episode, *seasonFlag, *IdItem)
		}
	} else {
		DownloadSeries(*episodeFlag, seasons[*seasonFlag].Episodes[*episodeFlag], *seasonFlag, *IdItem)
	}
}
