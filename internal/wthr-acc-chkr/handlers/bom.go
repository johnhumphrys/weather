package handlers

import (
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"time"
)

func CallWthrFtpSvc() {
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
	println(string(buf))
}
