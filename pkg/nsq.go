package pkg

import (
	"errors"

	"github.com/nsqio/go-nsq"
	"github.com/ropel12/scheduler/config"
)

func NewNSQ(conf *config.Config) (np *NSQProducer, err error) {
	np = &NSQProducer{}
	np.Env = conf.NSQ
	nsqConfig := nsq.NewConfig()
	np.Producer, err = nsq.NewProducer(np.Env.Host+":"+np.Env.Port, nsqConfig)
	if err != nil {
		return nil, err
	}

	return np, nil
}

type NSQProducer struct {
	Producer *nsq.Producer
	Env      config.NSQConfig
}

func (np *NSQProducer) Publish(Topic string, message []byte) error {
	switch Topic {
	case "1":
		return np.Producer.Publish(np.Env.Topic, message)
	case "2":
		return np.Producer.Publish(np.Env.Topic2, message)
	}
	return errors.New("Topic not available")
}

func (np *NSQProducer) Stop() {
	np.Producer.Stop()
}
