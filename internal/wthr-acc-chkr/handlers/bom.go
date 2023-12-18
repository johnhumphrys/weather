package handlers

import (
	"encoding/xml"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
	"johnhumphrys.dev/internal/wthr-acc-chkr/handlers/bommodel"
)

const (
	airTmpMin    = "air_temperature_minimum"
	airTmpMax    = "air_temperature_maximum"
	MelbourneAAC = "VIC_PT042"

	ftpTimeout = 5 * time.Second
)

func CallWthrFtpSvc() []byte {
	c, err := ftp.Dial("ftp.bom.gov.au:21", ftp.DialWithTimeout(ftpTimeout))
	if err != nil {
		log.Fatalf("Error connecting to FTP server: %v", err)
	}

	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatalf("Error logging into FTP server: %v", err)
	}

	err = c.ChangeDir("anon/gen/fwo")
	if err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	r, err := c.Retr("IDV10450.xml")
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
	data := CallWthrFtpSvc()

	var bom bommodel.Product

	err := xml.Unmarshal(data, &bom)
	if err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	var targetArea bommodel.Area

	for _, area := range bom.Forecast.Area {
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
