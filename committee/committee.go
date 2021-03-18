package committee

import (
	"container/list"
	"sync"
)

var Mutex sync.Mutex
var CommitteeList list.List
