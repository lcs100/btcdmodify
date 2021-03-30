package cpu

import "sync"

var EnergyPerBlock float64 = 0.0
var TotalEnergy float64 = 0.0

var Mutex sync.Mutex
var ProofNumber int64 = 0

var Mutex1 sync.Mutex
var Flag int64 = 0

var Mutex2 sync.Mutex
var WeakBlocks int64 = 0

var WeakNodes = 3
var StrongNodes = 1

var Mutex3 sync.Mutex
var StrongBlocks int64 = 0

var Mutex4 sync.Mutex
var WeakBlocks1 int64 = 0

var Type = 1

var Duration float64 = 0.0
