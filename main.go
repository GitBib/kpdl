package main

import (
	"flag"
	"fmt"
	"io"
	"kpdl/api"
	_struct "kpdl/api/struct"
	"net/http"
	"os"
	"strconv"
	"strings"
)

import humanize "github.com/dustin/go-humanize"

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
	dir := strconv.Itoa(itemId) + "/"
	os.MkdirAll(dir, os.ModePerm)
	dirFull := dir + fileName + ".mp4"
	fmt.Printf("\rStarting download %s\r", fileName)
	err := DownloadFile(dirFull, urlFile, fileName)
	if err != nil {
		panic(err)
	}
}

type WriteCounter struct {
	Total uint64
	Title string
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading %s... %s complete", wc.Title, humanize.Bytes(wc.Total))
}

func main() {
	ApiKey := flag.String("key", "foo", "Api Key")
	IdItem := flag.Int("item", 0, "id item")
	seasonFlag := flag.Int("s", 0, "id season")
	episodeFlag := flag.Int("e", -1, "id episode")
	flag.Parse()
	fmt.Println("key: ", *ApiKey, " item: ", *IdItem, " s:", *seasonFlag, " e:", *episodeFlag)
	var api = api.Api{ApiKey: *ApiKey}
	seasons := api.GetInfoItem(*IdItem).Seasons

	if seasons == nil {
		fmt.Println("Couldn't get the contents!")
		os.Exit(1)
	}

	if -1 == *episodeFlag {
		for e, episode := range seasons[*seasonFlag].Episodes {
			DownloadSeries(e, episode, *seasonFlag, *IdItem)
		}
	} else {
		DownloadSeries(*episodeFlag, seasons[*seasonFlag].Episodes[*episodeFlag], *seasonFlag, *IdItem)
	}
}

func DownloadFile(filepath string, url string, nameFile string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{Title: nameFile}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
