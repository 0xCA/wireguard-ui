package util

import "sync"

var TgUseridToClientID = map[int64]([]string){}
var TgUseridToClientIDMutex sync.RWMutex
