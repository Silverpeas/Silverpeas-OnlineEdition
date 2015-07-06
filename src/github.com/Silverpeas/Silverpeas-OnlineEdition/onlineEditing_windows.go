package main

import "fmt"
import "log"
import "os"
import "strings"
import "github.com/skratchdot/open-golang/open"
import "golang.org/x/sys/windows/registry"

func main() {
	fmt.Printf("Hello, world.\n")
	
	customURL := os.Args[1]
	log.Printf(customURL)
	
	customURL2 := strings.Replace(customURL, "spwebdav://", "http://", 1)
	webDavURL := strings.Replace(customURL2, "spwebdavs://", "https://", 1)
	
	openWithWindows(webDavURL)
}

func openWithWindows(url string) {
	log.Print("This is Windows !")
	
	if isMSOfficeInstalled() {
		openWithMSOffice(url)
	} else {
		openWithOpenOffice(url)
	}
}

func isMSOfficeInstalled() bool {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, "Word.Application", registry.QUERY_VALUE)
	if err == registry.ErrNotExist {
		log.Print(err)	
		return false
	} else {
		return true
	}
	defer k.Close()
	return false
}

func openWithOpenOffice(url string) {
	getMSApplication(url)
	err2 := open.RunWith(url, "soffice.exe")
	if err2 != nil {
		log.Fatal(err2)
	}
}

func openWithMSOffice(url string) {
	err2 := open.RunWith(url, getMSApplication(url))
	if err2 != nil {
		log.Fatal(err2)
	}
}

func getMSApplication(url string) string {
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	extension := strings.Split(fileName, ".")[1]
	if strings.HasPrefix(extension, "ppt") || strings.HasPrefix(extension, "pot") || strings.HasPrefix(extension, "pps") {
		log.Print("Extension = "+extension+" -> powerpnt.exe")
		return "powerpnt.exe"
	} else if strings.HasPrefix(extension, "xls") || strings.HasPrefix(extension, "xlt") {
		log.Print("Extension = "+extension+" -> excel.exe")
		return "excel.exe"
	} else {
		log.Print("Extension = "+extension+" -> winword.exe")
		return "winword.exe"
	}
}