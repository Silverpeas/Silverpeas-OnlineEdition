package main

import "log"
import "fmt"
import "github.com/skratchdot/open-golang/open"

func main() {
	fmt.Printf("Hello, world.\n")
	
	openWithMac()
}

func openWithMac() {
	log.Print("This is Mac OS")
	
	err2 := open.RunWith("https://intranoo.silverpeas.com/silverpeas/repository/jackrabbit/webdav/attachments/kmelia1456/b97aeafc-0448-4d32-9f9f-8e1810a406f0/fr/Soutenance.odp", "libreoffice")
	if err2 != nil {
		log.Fatal(err2)
	}
}