package arbitrator

import (
	"encoding/json"
	"log"
	"io/ioutil"
)

type InstanceInfo struct {
	Count uint32
	IsUseTimePriority bool
	Images []string
	Addresses []string
}

type ExceptionRuleInfo struct {
	Threshold uint32
	MaxFailures uint32
	RestartImage string
}

type ConfigInfo struct {
	Instance *InstanceInfo
	ExceptionRule *ExceptionRuleInfo
}

type Parser struct {}

func (p *Parser) Parse(file string, v interface{}) {
	var data []byte
	if file != "" {
		var err error
		if data, err = ioutil.ReadFile(file); err != nil {
			log.Fatalf("Read config file: %s failed, reson: %v \n", file, err)
		}
	} else {
		log.Fatal("Config file path is empty")
	}

	if err := json.Unmarshal(data, v); err != nil {
		log.Fatalf("Parse config file: %s failed, reason: %v \n", file, err)
	}
}

func NewParser() *Parser {
	return &Parser{}
}

