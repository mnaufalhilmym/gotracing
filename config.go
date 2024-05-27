package gotracing

type config struct {
	maxLevel levelFilter
	maxPC    int
}

var conf = config{
	maxLevel: levelFilterOff,
	maxPC:    10,
}

func SetMaxLevel(level levelFilter) {
	conf.maxLevel = level
}

func SetMaxProgramCounters(count int) {
	conf.maxPC = count
}
