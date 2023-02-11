package handlers

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
)

type Info struct {
	Function string
}

//MakeInfoHandler creates handler for /watchdog-info endpoint
func MakeWatchDogInfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			defer r.Body.Close()
		}

		info, err := parseInfo(r)
		log.Println("=================== IP ==========", readUserIP(r))
		log.Println("=================== INFO ==========", info)

		if err != nil {
			log.Printf("[Watchdog Info] error %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func parseInfo(r *http.Request) (Info, error) {
	info := Info{}
	bytesOut, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(bytesOut, &info)
	if err != nil {
		return info, err
	}

	return info, err
}

// Ref: https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
func readUserIP(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}
