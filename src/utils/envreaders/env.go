package envreaders

import (
	"bufio"
	log "golang-heroku/src/business/usecases/log"
	"os"
	"strings"
)

type Interface interface {
	Load(log log.Logger)
}

type Options struct {
	Name string
}

type envReader struct {
	opt Options
}

func Init(opt Options) Interface {
	e := &envReader{
		opt: opt,
	}

	return e
}

// Load is load configs from a env file.
func (e *envReader) Load(logger log.Logger) {
	f, err := os.Open(e.opt.Name)
	if err != nil {
		logger.LogError("%s", err)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		logger.LogError("%s", err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
}
