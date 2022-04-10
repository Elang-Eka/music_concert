package envreaders

import (
	log "azura-test/src/business/usecases/log"
	"bufio"
	"fmt"
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
	// Cek dir saat ini
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(pwd)

	f, err := os.Open(e.opt.Name)
	if err != nil {
		fmt.Println(err.Error())
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
