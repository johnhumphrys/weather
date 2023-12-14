package handlers

import (
	"encoding/xml"
	"github.com/jlaffaye/ftp"
	"io"
	"johnhumphrys.dev/internal/wthr-acc-chkr/handlers/bommodel"
	"log"
	"time"
)

func CallWthrFtpSvc() []byte {
	c, err := ftp.Dial("ftp.bom.gov.au:21", ftp.DialWithTimeout(5*time.Second))
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

	defer r.Close()

	buf, err := io.ReadAll(r)
	return buf
}

func DoSomething() {
	data := CallWthrFtpSvc()

	const airTmpMin = "air_temperature_minimum"
	const airTmpMax = "air_temperature_maximum"

	var bom bommodel.Product

	err := xml.Unmarshal(data, &bom)
	if err != nil {
		log.Fatalf("Error unmarshalling XML: %v", err)
	}

	println(bom.AMOC.Source.Sender)
	println(bom.AMOC.IssueTimeLocal.Time.Format(time.RFC3339))

	targetAAC := "VIC_PT042"
	var targetArea bommodel.Area

	for _, area := range bom.Forecast.Area {
		if area.AAC == targetAAC {
			targetArea = area
			break
		}
	}

	for _, area := range targetArea.ForecastPeriods {
		println(area.StartTimeLocal.Time.Format(time.RFC3339))

		for _, element := range area.Elements {

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
