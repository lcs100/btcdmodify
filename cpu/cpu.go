package cpu

import "sync"

var EnergyPerBlock int = 0
var TotalEnergy int64 = 0

var Mutex sync.Mutex
var ProofNumber int64 = 0

var Mutex1 sync.Mutex
var Flag int64 = 0

var Mutex2 sync.Mutex
var WeakBlocks int64 = 0

var WeakNodes = 7
var StrongNodes = 2
