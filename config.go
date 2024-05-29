package gotracing

type config struct {
	minConsolePrintLevel LevelFilter
	minStoreLevel        LevelFilter
	storage              storage
	maxPC                uint
}

var conf = config{
	minConsolePrintLevel: LevelFilterOff,
	minStoreLevel:        LevelFilterOff,
	storage:              nil,
	maxPC:                10,
}

func SetMinConsolePrintLevel(level LevelFilter) {
	conf.minConsolePrintLevel = level
}

func SetMinStoreLevel(level LevelFilter) {
	conf.minStoreLevel = level
}

func SetStorage(storage storage) {
	conf.storage = storage
}

func SetMaxProgramCounters(count uint) {
	conf.maxPC = count
}
