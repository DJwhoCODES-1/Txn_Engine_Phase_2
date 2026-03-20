package utils

import (
	"fmt"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

var counter uint32

func pad(num int, size int) string {
	return fmt.Sprintf("%0*d", size, num)
}

func GenerateReqID() string {
	timestamp := time.Now().UnixMilli()
	pid := os.Getpid() % 1000
	random := rand.Intn(1000)

	c := atomic.AddUint32(&counter, 1) % 1000

	return fmt.Sprintf("UTR%d%s%s%s",
		timestamp,
		pad(pid, 3),
		pad(random, 3),
		pad(int(c), 3),
	)
}
