package pitaya

import "github.com/topfreegames/pitaya/component"

var daoComps = make([]daoComp,0)

type daoComp struct {
	comp component.Component
	ops []component.Option
}

func RegisterDao(c component.Component,options ...component.Option)  {
	daoComps = append(daoComps,daoComp{c,options})
}

func startDaos()  {
	for _,c := range daoComps{
		c.comp.Init()
	}
	for _,c := range daoComps{
		c.comp.AfterInit()
	}
}

func shutdownDaos()  {
	length := len(daoComps)
	for i := length - 1; i >= 0; i-- {
		daoComps[i].comp.BeforeShutdown()
	}

	// reverse call `Shutdown` hooks
	for i := length - 1; i >= 0; i-- {
		daoComps[i].comp.Shutdown()
	}
}