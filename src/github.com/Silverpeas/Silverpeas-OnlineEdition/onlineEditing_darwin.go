package main

import "os"

func main() {
	info("Silverpeas Online Edition")
	customURL := os.Args[1]
	webDavURL := customToWebdavURL(customURL)
	runnable := "libreoffice"
	runDesktopApplicationWith(webDavURL, runnable)
}
