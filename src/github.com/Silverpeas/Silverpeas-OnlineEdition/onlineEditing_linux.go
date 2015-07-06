package main

import "log"
import "os"
import "strings"
import "github.com/skratchdot/open-golang/open"

func main() {
	customURL := os.Args[1]
	
	customURL2 := strings.Replace(customURL, "spwebdav://", "http://", 1)
	webDavURL := strings.Replace(customURL2, "spwebdavs://", "https://", 1)
	
	err2 := open.RunWith(webDavURL, "soffice")
	if err2 != nil {
		log.Fatal(err2)
	}
}