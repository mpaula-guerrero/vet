package api

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"test_ecapture_backend/internal/env"
	"test_ecapture_backend/internal/logger"

	"github.com/gofiber/fiber/v2"

	"github.com/fatih/color"
)

const (
	version = "0.0.0"
	website = ""
	banner  = `test`

	description = `test - %s - Port: %s
by Paula
Version: %s
%s`
)

type server struct {
	listening string
	app       string
	fb        *fiber.App
}

func newServer(listening int, app string, fb *fiber.App) *server {
	return &server{fmt.Sprintf(":%d", listening), app, fb}
}

func (srv *server) Start() {
	e := env.NewConfiguration()
	color.Blue(banner)
	color.Cyan(fmt.Sprintf(description, srv.app, srv.listening, version, website))

	if e.App.TLS {
		ln, _ := net.Listen("tcp", srv.listening)

		cer, err := tls.LoadX509KeyPair(e.App.Cert, e.App.Key)
		if err != nil {
			logger.Error.Printf("error al leer los certificados, error: " + err.Error())
			log.Fatal(err)
		}
		ln = tls.NewListener(ln, &tls.Config{
			Certificates:     []tls.Certificate{cer},
			MinVersion:       tls.VersionTLS13,
			CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		})
		log.Fatal(srv.fb.Listener(ln))
	} else {
		log.Fatal(srv.fb.Listen(srv.listening))
	}
}
