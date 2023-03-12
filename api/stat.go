package api

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type StatController struct{}

func NewStatController() *StatController {
	return &StatController{}
}

func (c *StatController) Home(w http.ResponseWriter, r *http.Request) {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Heap alloc: ", ms.HeapAlloc)
	fmt.Println("Heap use: ", ms.HeapInuse)
	fmt.Println("Heap release: ", ms.HeapReleased)
	fmt.Println("Heap object: ", ms.HeapObjects)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
	w.Write([]byte(time.Now().String()))
}
