package handlers

import (
	"encoding/xml"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
	"johnhumphrys.dev/internal/weather/handlers/bommodel"
)

const (
	airTmpMin    = "air_temperature_minimum"
	airTmpMax    = "air_temperature_maximum"
	MelbourneAAC = "VIC_PT042"

	ftpTimeout = 5 * time.Second
)

func getXml() []byte {
	ftpAddress := "ftp.bom.gov.au:21"
	ftpUser := "anonymous"
	ftpPw := "anonymous"
	ftpDir := "anon/gen/fwo"
	ftpFile := "IDV10450.xml"

	c, err := ftp.Dial(ftpAddress, ftp.DialWithTimeout(ftpTimeout))
	if err != nil {
		log.Fatalf("Error connecting to FTP server: %v", err)
	}

	err = c.Login(ftpUser, ftpPw)
	if err != nil {
		log.Fatalf("Error logging into FTP server: %v", err)
	}

	err = c.ChangeDir(ftpDir)
	if err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	r, err := c.Retr(ftpFile)
	if err != nil {
		log.Fatalf("Error retrieving file: %v", err)
	}

	buf, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	defer r.Close()

	return buf
}

func DoSomething() {
	data := getXml()

	var wthrData bommodel.Product

	err := xml.Unmarshal(data, &wthrData)
	if err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	var targetArea bommodel.Area

	for _, area := range wthrData.Forecast.Area {
		if area.AAC == MelbourneAAC {
			targetArea = area
			break
		}
	}

	for i := range targetArea.ForecastPeriods {
		area := &targetArea.ForecastPeriods[i]
		println(area.StartTimeLocal.Time.Format(time.RFC3339))

		for j := range area.Elements {
			element := &area.Elements[j]

			if element.Type == airTmpMin {
				println(element.Type)
				println(element.Text)
			}

			if element.Type == airTmpMax {
				println(element.Type)
				println(element.Text)
			}
		}
	}
}
