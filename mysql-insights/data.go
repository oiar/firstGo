package main

import (
	"fmt"
	"os"
	"sync"
)

const (
	nasdaqHistoryFormat = ""
)

type companies struct {
	total int64
}

func (c *companies) loadHistoryForCompany(code string) error {
	f, err := os.Open(fmt.Sprintf(nasdaqHistoryFormat, code))
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
}
