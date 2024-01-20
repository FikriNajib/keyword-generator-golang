package shared

import (
	"keyword-generator/src/infrastructure/shared/logger"
	"net"
	"strings"
)

func GetErrorStatusMessageHelper(statusMessage StatusMessage, err error) StatusMessage {
	if err != nil {
		logger.WithFields(logger.Fields{"error": err}).Errorf("error : %v", err)
		return statusMessage
	}

	return statusMessage
}

func LookupLocalIPv4() (string, error) {
	var (
		ip  string
		err error
	)

	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ip, err
	}

	for _, a := range addr {
		if strings.ContainsAny(a.String(), ".") { //&& !strings.ContainsAny(a.String(), "127") {
			ip = a.String()
		}
	}

	return ip, err
}
