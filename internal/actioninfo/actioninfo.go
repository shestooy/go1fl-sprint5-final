package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("parsing error: %s", err.Error())
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("get info about action error: %s", err.Error())
			continue
		}

		fmt.Println(info)
	}
}
