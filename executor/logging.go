package executor

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"
)

func bindLoggingPipe(name string, output io.Reader) {
	log.WithFields(log.Fields{
		"pipe": name,
	}).Printf("Started logging %s from function.", name)

	scanner := bufio.NewScanner(output)

	logger := log.WithFields(log.Fields{
		"pipe": name,
	})

	go func() {
		for scanner.Scan() {
			logger.Print(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			logger.WithFields(log.Fields{
				"err": err,
			}).Printf("Error scanning %s: %s", name, err.Error())
		}
	}()
}
