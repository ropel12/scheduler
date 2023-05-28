package pkg

import (
	"errors"

	"github.com/pusher/pusher-http-go"
	"github.com/ropel12/scheduler/config"
)

type Pusher struct {
	Client *pusher.Client
	Env    config.PusherConfig
}

func NewPusher(conf *config.Config) (ps *Pusher) {
	ps = &Pusher{}
	ps.Env = conf.Pusher
	ps.Client = &pusher.Client{
		AppID:   ps.Env.AppId,
		Key:     ps.Env.Key,
		Secret:  ps.Env.Secret,
		Cluster: ps.Env.Cluster,
		Secure:  ps.Env.Secure,
	}
	return ps
}

func (p *Pusher) Publish(data any, event int) error {
	switch event {
	case 1:
		return p.Client.Trigger(p.Env.Channel, p.Env.Event1, data)
	case 2:
		return p.Client.Trigger(p.Env.Channel, p.Env.Event2, data)
	case 3:
		return p.Client.Trigger(p.Env.Channel, p.Env.Event3, data)
	}
	return errors.New("Event Not Available")
}
