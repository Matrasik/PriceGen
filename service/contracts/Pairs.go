package contracts

import (
	"sync"
)

type Pairs interface {
	GenVal(wg *sync.WaitGroup)
}
