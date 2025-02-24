package visistservice

import (
	"log"
	"net/http"
	"service/logservice"
	"strconv"
	"sync/atomic"
)

type visistService struct {
	visistCount atomic.Int32
}

func Init() {
	s := &visistService{
		visistCount: atomic.Int32{},
	}
	s.register()
}

func (s *visistService) register() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.visistCount.Add(1)
		count := strconv.Itoa(int(s.visistCount.Load()))
		err := logservice.Println(count)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("Log service println error: %s\n", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(count))
	})
}
