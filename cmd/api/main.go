package main

import (
	"o_clinic/cmd/server"

	log "github.com/sirupsen/logrus"
)

//	@title			O-Clinic API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	omarsabra509@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// 	@host			localhost:8080
// 	@BasePath		/api
func init() {
	log.SetLevel(log.DebugLevel)
}


func main() {
	s := server.NewServer() 

	if err := s.ListenAndServe(); err != nil {
		log.Errorln(err)
	}
}   