package twins

type State struct {
}

type Process struct {
}

type Trace struct {
	p         Process
	stateInfo map[string]*StateInfo
}

type StateInfo struct {
	enabled []State
	isHot   bool
	hash    string
}

func GetNext(s State) (p Process) {

	return
}

func Scheduled() {

}
