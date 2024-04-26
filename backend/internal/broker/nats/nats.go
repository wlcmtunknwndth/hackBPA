package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/wlcmtunknwndth/hackBPA/internal/config"
)

type Nats struct {
	*nats.Conn
}

func New(cfg *config.Config) (*Nats, error) {
	const op = "broker.nats.New"

	natsService, err := nats.Connect(cfg.Nats.Address,
		nats.RetryOnFailedConnect(cfg.Nats.Retry),
		nats.MaxReconnects(cfg.Nats.MaxReconnects),
		nats.ReconnectWait(cfg.Nats.ReconnectWait),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Nats{natsService}, nil
}

func (n *Nats) Close() error {
	return n.Close()
}
