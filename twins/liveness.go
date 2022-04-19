package twins

type Liveness struct {
	s       State
	method  string
	maxStep int
	trace   []Trace
	temp    int
	tt      int
	rt      int
}

func (l *Liveness) ExecProgram(s0 State) {
	l.s = s0
	l.trace = make([]Trace, 10)
	l.temp = 0
	n := 0

	for l.Enabled() && n < l.maxStep {

		p := GetNext(l.s)
		ss := TransferState(p, l.s)
		l.trace = append(l.trace, Trace{
			p: p,
		})
		l.s = ss
		n += 1

		if l.method == "Temp" {
			l.checkTemp()
		} else if l.method == "Caching" {
			l.checkLasso()
		}
	}

}

func (l *Liveness) checkTemp() {

	if IsHot() {
		l.temp += 1
		if l.temp == l.tt {
			l.reportLiveness()
		}
	} else {
		l.temp = 0
	}

}

func (l *Liveness) checkLasso() {

}

func ReplayCycle() {

}

func TransferState(p Process, s State) (ss State) {

	return
}

func (l *Liveness) Enabled() bool {
	return false
}

func IsHot() bool {
	return false
}

func (l *Liveness) reportLiveness() {

}

func Hash() {

}
