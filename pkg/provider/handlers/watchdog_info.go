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
