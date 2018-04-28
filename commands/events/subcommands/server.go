package subcommands

import (
	log "github.com/sirupsen/logrus"

	"github.com/ronelliott/go-events"
)

type ServerCmd struct {
	Address string `
		long:"address"
		env:"ADDRESS"
		default:"0.0.0.0:9999"`

	Debug bool `
		short:"d"
		long:"debug"
		env:"DEBUG"`
}

func (opts *ServerCmd) Run() error {
	if opts.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	errs := make(chan error)
	defer close(errs)

	go func() {
		for err := range errs {
			log.Error(err)
		}
	}()

	srv := events.NewServer()
	log.Info("Listening on: ", opts.Address)
	srv.Listen(opts.Address, errs)
	return nil
}
