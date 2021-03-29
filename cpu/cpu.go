package cpu

import "sync"

var EnergyPerBlock int = 0
var TotalEnergy int64 = 0

var Mutex sync.Mutex
var ProofNumber int64 = 0
