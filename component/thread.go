package component

import "github.com/topfreegames/pitaya/logger"

type (
	LogicThread struct {
		Base
		Name string
		option options
		DieChan chan struct{}
		Run func()
	}

	Runnable interface {
		Component
		SetName(opts []Option)
		GetName() string
	}
)

func (p *LogicThread) SetName(opts []Option)  {
	for _,opt := range opts {
		opt(&p.option)
	}

	if name := p.option.name;name != "" {
		p.Name = name
	}
}

func (p *LogicThread) GetName() string {
	return p.Name
}

func (p *LogicThread) Init()  {
	go func() {
		for {
			select {
			case <- p.DieChan:
				logger.Log.Infof("线程%s关闭",p.Name)
				close(p.DieChan)
				return
			default:
				p.Run()
			}
		}
	}()
}

func (p *LogicThread) Shutdown()  {
	p.DieChan <- struct{}{}
}







