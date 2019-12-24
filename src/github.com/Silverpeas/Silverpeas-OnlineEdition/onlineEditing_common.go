package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"log"
	"regexp"
	"strings"
)

func info(message string) {
	fmt.Println(message)
}

func customToWebdavURL(customURL string) string {
	info("Silverpeas Protocol URL: " + customURL)
	webDavURL := string(regexp.MustCompile("^spwebdav([s]?)://").ReplaceAll([]byte(customURL), []byte("http$1://")))
	info("Webdav URL: " + webDavURL)
	return webDavURL
}

func getFileExtension(url string) string {
	ext := string(regexp.MustCompile(".*[.][ ]*([a-zA-Z0-9]+)[ ]*$").ReplaceAll([]byte(url), []byte("$1")))
	if strings.EqualFold(url, ext) {
		return ""
	}
	return strings.ToLower(ext)
}

func runDesktopApplicationWith(webDavURL string, executable string) {
	err := open.RunWith(webDavURL, executable)
	if err != nil {
		log.Fatal(err)
	}
}
