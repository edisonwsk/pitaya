package pitaya

import (
	"github.com/topfreegames/pitaya/component"
)

var (
	threadsComp = make([]threadComp,0)
	ThreadPool = &threadPool{map[string]component.Runnable{},}
)

type threadPool struct {
	pool map[string]component.Runnable
}

type threadComp struct {
	comp component.Runnable
	opts []component.Option
}

func RegisterThread(c component.Runnable, opts ...component.Option) {
	threadsComp = append(threadsComp, threadComp{c, opts})
}

func (p *threadPool) RegisterThread(name string,thread component.Runnable)  {
	p.pool[name] = thread
}

func (p *threadPool) GetThread(name string) component.Runnable {
	return p.pool[name]
}

func startThreads()  {
	for _,t := range threadsComp {
		t.comp.SetName(t.opts)
		t.comp.Init()
		ThreadPool.RegisterThread(t.comp.GetName(),t.comp)
	}

	for _,t := range threadsComp {
		t.comp.AfterInit()
	}
}

func shutdownThreads() {
	// reverse call `BeforeShutdown` hooks
	length := len(threadsComp)
	for i := length - 1; i >= 0; i-- {
		threadsComp[i].comp.BeforeShutdown()
	}

	// reverse call `Shutdown` hooks
	for i := length - 1; i >= 0; i-- {
		threadsComp[i].comp.Shutdown()
	}
}


