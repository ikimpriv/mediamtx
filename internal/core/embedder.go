package core

import (
	"context"
	"fmt"
	"github.com/bluenviron/mediamtx/internal/conf"
	"github.com/bluenviron/mediamtx/internal/logger"
)

func NewEmbed(ctx context.Context, cancelFunc context.CancelFunc) (*Core, bool) {
	p := &Core{
		ctx:            ctx,
		ctxCancel:      cancelFunc,
		chAPIConfigSet: make(chan *conf.Conf),
		done:           make(chan struct{}),
	}

	tempLogger, _ := logger.New(logger.Warn, []logger.Destination{logger.DestinationStdout}, "")

	var err error
	p.conf, p.confPath, err = conf.Load(cli.Confpath, defaultConfPaths, tempLogger)
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		return nil, false
	}

	err = p.createResources(true)
	if err != nil {
		if p.logger != nil {
			p.Log(logger.Error, "%s", err)
		} else {
			fmt.Printf("ERR: %s\n", err)
		}
		p.closeResources(nil, false)
		return nil, false
	}

	go p.run()

	return p, true
}
