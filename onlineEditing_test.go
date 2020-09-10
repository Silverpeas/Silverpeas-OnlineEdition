package main

import (
	"strings"
	"testing"
)

func TestCustomToWebdavURL(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		result := customToWebdavURL("")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("no custom protocol", func(t *testing.T) {
		result := customToWebdavURL("spwebtoto://dummy/toto.odp")
		if !strings.EqualFold(result, "spwebtoto://dummy/toto.odp") {
			t.Fail()
		}
	})
	t.Run("custom protocol", func(t *testing.T) {
		result := customToWebdavURL("spwebdav://dummy/toto.odp")
		if !strings.EqualFold(result, "http://dummy/toto.odp") {
			t.Fail()
		}
	})
	t.Run("secure custom protocol", func(t *testing.T) {
		result := customToWebdavURL("spwebdavs://dummy/toto.odp")
		if !strings.EqualFold(result, "https://dummy/toto.odp") {
			t.Fail()
		}
	})
}

func TestGetFileExtension(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		result := getFileExtension("")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("no point", func(t *testing.T) {
		result := getFileExtension("https://dummy/totoodp")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("lower case extension", func(t *testing.T) {
		result := getFileExtension("https://dummy/toto.odp")
		if !strings.EqualFold(result, "odp") {
			t.Fail()
		}
	})
	t.Run("upper case extension", func(t *testing.T) {
		result := getFileExtension("https://dummy/toto.ODP")
		if !strings.EqualFold(result, "ODP") {
			t.Fail()
		}
	})
}

func TestGetApplicationRunnable(t *testing.T) {
	t.Run("odt", func(t *testing.T) {
		result := getApplicationRunnable("odt")
		if !strings.EqualFold(result, "swriter.exe") {
			t.Fail()
		}
	})
	t.Run("odp", func(t *testing.T) {
		result := getApplicationRunnable("odp")
		if !strings.EqualFold(result, "simpress.exe") {
			t.Fail()
		}
	})
	t.Run("ods", func(t *testing.T) {
		result := getApplicationRunnable("ods")
		if !strings.EqualFold(result, "scalc.exe") {
			t.Fail()
		}
	})
	t.Run("unknown", func(t *testing.T) {
		result := getApplicationRunnable("unknown")
		if !strings.EqualFold(result, "soffice.exe") {
			t.Fail()
		}
	})
	t.Run("doc", func(t *testing.T) {
		result := getApplicationRunnable("doc")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("dot", func(t *testing.T) {
		result := getApplicationRunnable("dot")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("docx", func(t *testing.T) {
		result := getApplicationRunnable("docx")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("dotx", func(t *testing.T) {
		result := getApplicationRunnable("dotx")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("xls", func(t *testing.T) {
		result := getApplicationRunnable("xls")
		if !strings.EqualFold(result, "excel.exe") {
			t.Fail()
		}
	})
	t.Run("xlsx", func(t *testing.T) {
		result := getApplicationRunnable("xlsx")
		if !strings.EqualFold(result, "excel.exe") {
			t.Fail()
		}
	})
	t.Run("ppt", func(t *testing.T) {
		result := getApplicationRunnable("ppt")
		if !strings.EqualFold(result, "powerpnt.exe") {
			t.Fail()
		}
	})
	t.Run("pptx", func(t *testing.T) {
		result := getApplicationRunnable("pptx")
		if !strings.EqualFold(result, "powerpnt.exe") {
			t.Fail()
		}
	})
	t.Run("mpp", func(t *testing.T) {
		result := getApplicationRunnable("mpp")
		if !strings.EqualFold(result, "winproj.exe") {
			t.Fail()
		}
	})
	t.Run("mpt", func(t *testing.T) {
		result := getApplicationRunnable("mpt")
		if !strings.EqualFold(result, "winproj.exe") {
			t.Fail()
		}
	})
}

func TestGetContentTypeFromRegistry(t *testing.T) {
	t.Run("odt", func(t *testing.T) {
		result := getContentTypeFromRegistry("odt")
		if !strings.EqualFold(result, "application/vnd.oasis.opendocument.text") {
			t.Fail()
		}
	})
	t.Run("odp", func(t *testing.T) {
		result := getContentTypeFromRegistry("odp")
		if !strings.EqualFold(result, "application/vnd.oasis.opendocument.presentation") {
			t.Fail()
		}
	})
	t.Run("ods", func(t *testing.T) {
		result := getContentTypeFromRegistry("ods")
		if !strings.EqualFold(result, "application/vnd.oasis.opendocument.spreadsheet") {
			t.Fail()
		}
	})
	t.Run("unknown", func(t *testing.T) {
		result := getContentTypeFromRegistry("unknown")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("doc", func(t *testing.T) {
		result := getContentTypeFromRegistry("doc")
		if !strings.EqualFold(result, "application/msword") {
			t.Fail()
		}
	})
	t.Run("dot", func(t *testing.T) {
		result := getContentTypeFromRegistry("dot")
		if !strings.EqualFold(result, "application/msword") {
			t.Fail()
		}
	})
	t.Run("docx", func(t *testing.T) {
		result := getContentTypeFromRegistry("docx")
		if !strings.EqualFold(result, "application/vnd.openxmlformats-officedocument.wordprocessingml.document") {
			t.Fail()
		}
	})
	t.Run("dotx", func(t *testing.T) {
		result := getContentTypeFromRegistry("dotx")
		if !strings.EqualFold(result, "application/vnd.openxmlformats-officedocument.wordprocessingml.template") {
			t.Fail()
		}
	})
	t.Run("xls", func(t *testing.T) {
		result := getContentTypeFromRegistry("xls")
		if !strings.EqualFold(result, "application/vnd.ms-excel") {
			t.Fail()
		}
	})
	t.Run("xlt", func(t *testing.T) {
		result := getContentTypeFromRegistry("xlt")
		if !strings.EqualFold(result, "application/vnd.ms-excel") {
			t.Fail()
		}
	})
	t.Run("xlsx", func(t *testing.T) {
		result := getContentTypeFromRegistry("xlsx")
		if !strings.EqualFold(result, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet") {
			t.Fail()
		}
	})
	t.Run("xltx", func(t *testing.T) {
		result := getContentTypeFromRegistry("xltx")
		if !strings.EqualFold(result, "application/vnd.openxmlformats-officedocument.spreadsheetml.template") {
			t.Fail()
		}
	})
	t.Run("ppt", func(t *testing.T) {
		result := getContentTypeFromRegistry("ppt")
		if !strings.EqualFold(result, "application/vnd.ms-powerpoint") {
			t.Fail()
		}
	})
	t.Run("pptx", func(t *testing.T) {
		result := getContentTypeFromRegistry("pptx")
		if !strings.EqualFold(result, "application/vnd.openxmlformats-officedocument.presentationml.presentation") {
			t.Fail()
		}
	})
	t.Run("mpp", func(t *testing.T) {
		result := getContentTypeFromRegistry("mpp")
		if !strings.EqualFold(result, "application/vnd.ms-project") {
			t.Fail()
		}
	})
	t.Run("mpt", func(t *testing.T) {
		result := getContentTypeFromRegistry("mpt")
		if !strings.EqualFold(result, "application/vnd.ms-project") {
			t.Fail()
		}
	})
}

func TestGetMSOpenApplication(t *testing.T) {
	t.Run("odt", func(t *testing.T) {
		result := getMSOpenApplication("odt")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("unknown", func(t *testing.T) {
		result := getMSOpenApplication("unknown")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("doc", func(t *testing.T) {
		result := getMSOpenApplication("doc")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("dot", func(t *testing.T) {
		result := getMSOpenApplication("dot")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("docx", func(t *testing.T) {
		result := getMSOpenApplication("docx")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("dotx", func(t *testing.T) {
		result := getMSOpenApplication("dotx")
		if !strings.EqualFold(result, "winword.exe") {
			t.Fail()
		}
	})
	t.Run("xls", func(t *testing.T) {
		result := getMSOpenApplication("xls")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("xlsx", func(t *testing.T) {
		result := getMSOpenApplication("xlsx")
		if !strings.EqualFold(result, "excel.exe") {
			t.Fail()
		}
	})
	t.Run("ppt", func(t *testing.T) {
		result := getMSOpenApplication("ppt")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("pptx", func(t *testing.T) {
		result := getMSOpenApplication("pptx")
		if !strings.EqualFold(result, "powerpnt.exe") {
			t.Fail()
		}
	})
	t.Run("mpp", func(t *testing.T) {
		result := getMSOpenApplication("mpp")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
	t.Run("mpt", func(t *testing.T) {
		result := getMSOpenApplication("mpt")
		if !strings.EqualFold(result, "") {
			t.Fail()
		}
	})
}

func TestIsMSProjectInstalled(t *testing.T) {
	result := isMSProjectInstalled()
	if !result {
		t.Fail()
	}
}
