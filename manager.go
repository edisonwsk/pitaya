package pitaya

import "github.com/topfreegames/pitaya/component"

var mgrComps = make([]mgrComp,0)

type mgrComp struct {
	comp component.Component
	ops []component.Option
}

func RegisterManager(c component.Component,options ...component.Option)  {
	mgrComps = append(mgrComps,mgrComp{c,options})
}

func startManagers()  {
	for _,c := range mgrComps{
		c.comp.Init()
	}
	for _,c := range mgrComps{
		c.comp.AfterInit()
	}
}

func shutdownManagers()  {
	length := len(mgrComps)
	for i := length - 1; i >= 0; i-- {
		mgrComps[i].comp.BeforeShutdown()
	}

	// reverse call `Shutdown` hooks
	for i := length - 1; i >= 0; i-- {
		mgrComps[i].comp.Shutdown()
	}
}