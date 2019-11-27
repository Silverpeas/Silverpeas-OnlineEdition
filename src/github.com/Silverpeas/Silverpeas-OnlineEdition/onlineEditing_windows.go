package main

import (
	"golang.org/x/sys/windows/registry"
	"regexp"
	"strings"
)
import "os"

func main() {
	info("Silverpeas Online Edition")
	customURL := os.Args[1]
	webDavURL := customToWebdavURL(customURL)
	openWithWindows(webDavURL)
}

func openWithWindows(webDavURL string) {
	extension := getFileExtension(webDavURL)
	runnable := getApplicationRunnable(extension)
	runDesktopApplicationWith(webDavURL, runnable)
}

func getApplicationRunnable(extension string) string {
	runnable := getApplicationRunnableFromRegistry(extension)
	if strings.EqualFold(runnable, "") {
		runnable = getMSOpenApplication(extension)
	}
	if strings.EqualFold(runnable, "") {
		runnable = getMSApplication(extension)
	}
	if strings.EqualFold(runnable, "") {
		runnable = getOOApplication(extension)
	}
	return runnable
}

func getMSOpenApplication(extension string) string {
	var runnable string = ""
	contentType := getContentTypeFromRegistry(extension)
	if isMSOfficeInstalled() && strings.Contains(contentType, "vnd.openxmlformats-officedocument") {
		if strings.Contains(contentType, ".word") {
			runnable = "winword.exe"
		} else if strings.Contains(contentType, ".spread") {
			runnable = "excel.exe"
		} else if strings.Contains(contentType, ".presentation") {
			runnable = "powerpnt.exe"
		}
	}
	if !strings.EqualFold(runnable, "") {
		info("Detecting MSOffice and guessing " + runnable + " runnable from '" + contentType + "' content type")
	}
	return runnable
}

func getMSApplication(extension string) string {
	var runnable string = ""
	if isMSProject(extension) {
		runnable = "winproj.exe"
		info("Guessing " + runnable + " runnable from '" + extension + "' extension")
	} else if isMSOfficeInstalled() {
		if isPowerPoint(extension) {
			runnable = "powerpnt.exe"
		} else if isExcel(extension) {
			runnable = "excel.exe"
		} else if isWord(extension) {
			runnable = "winword.exe"
		}
		if !strings.EqualFold(runnable, "") {
			info("Detecting MSOffice and guessing " + runnable + " runnable from '" + extension + "' extension")
		}
	}
	return runnable
}

func getOOApplication(extension string) string {
	runnable := "soffice.exe"
	defer info("Guessing " + runnable + " runnable from '" + extension + "' extension")
	return runnable
}

func isWord(extension string) bool {
	return strings.HasPrefix(extension, "doc") || strings.HasPrefix(extension, "dot")
}

func isPowerPoint(extension string) bool {
	return strings.HasPrefix(extension, "ppt") || strings.HasPrefix(extension, "pot")
}

func isExcel(extension string) bool {
	return strings.HasPrefix(extension, "xls") || strings.HasPrefix(extension, "xlt") || strings.HasPrefix(extension, "xlam")
}

func isMSProject(extension string) bool {
	return strings.HasPrefix(extension, "mpp") || strings.HasPrefix(extension, "mpt")
}

func isMSProjectInstalled() bool {
	return isRegistryPathExisting("MSProject.Application")
}

func isMSOfficeInstalled() bool {
	return isRegistryPathExisting("Word.Application")
}

func isRegistryPathExisting(path string) bool {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, path, registry.QUERY_VALUE)
	defer k.Close()
	if err != nil {
		return false
	}
	return true
}

func getApplicationRunnableFromRegistry(extension string) string {
	var runnable string = ""
	fileClass := getValueSafelyFromRegistry("."+extension, "")
	if !strings.EqualFold(fileClass, "") {
		fullRunnable := getValueSafelyFromRegistry(fileClass+"\\Shell\\Open\\Command", "")
		if !strings.EqualFold(fullRunnable, "") {
			runnable = string(regexp.MustCompile("(?i).*[\\\\/]([^\\\\/]+[.]exe).*").ReplaceAll([]byte(fullRunnable), []byte("$1")))
			if strings.EqualFold(runnable, fullRunnable) {
				runnable = ""
			}
		}
	}
	runnable = strings.ToLower(runnable)
	if !strings.EqualFold(runnable, "") {
		info("Finding " + runnable + " runnable from registry base")
	}
	return runnable
}

func getContentTypeFromRegistry(extension string) string {
	contentType := getValueSafelyFromRegistry("."+extension, "Content Type")
	return strings.ToLower(contentType)
}

func getValueSafelyFromRegistry(path string, key string) string {
	k, err := registry.OpenKey(registry.CLASSES_ROOT, path, registry.QUERY_VALUE)
	defer k.Close()
	if err != nil {
		return ""
	}
	value, _, err := k.GetStringValue(key)
	if err != nil {
		return ""
	}
	return value
}
