package subcommands

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/ronelliott/go-events"
)

type ClientCmd struct {
	Address string `
		long:"address"
		env:"ADDRESS"
		default:"0.0.0.0:9999"`

	Debug bool `
		long:"debug"
		short:"d"
		env:"DEBUG"`
}

func (opts *ClientCmd) Run() error {
	if opts.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug enabled")
	} else {
		log.SetLevel(log.InfoLevel)
	}

	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		return err
	}

	log.Info("Connecting to: ", opts.Address)
	client, err := events.NewClient(opts.Address)

	if err != nil {
		return err
	}

	defer client.Close()

	fmt.Print("> ")
	for scanner.Scan() {
		data := scanner.Bytes()
		err = client.Send(data)

		if err != nil {
			return err
		}

		fmt.Println(string(data))
		fmt.Print("> ")
	}

	return nil
}
