// Copyright (c) 2015 Leonid Kneller

// "PCH2: Parallel convex hull (in 2 dimensions)."
package main

import (
	"flag"
	"fmt"
	"github.com/reconditematter/pq"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Start")
	runtime.GOMAXPROCS(2 * runtime.NumCPU())
	//
	ncpuflag := flag.Int("ncpu", 0, "number of CPUs")
	flag.Parse()
	ncpu := *ncpuflag
	if ncpu <= 0 || ncpu > runtime.NumCPU() {
		ncpu = runtime.NumCPU()
	}
	fmt.Println("NumCPU", ncpu)
	//
	const N = 100000
	fmt.Println("Point#", N)
	ps := make([]pq.Point2q, N)
	for i := 0; i < N; i++ {
		//
		xf, yf := genI2()
		//
		xq, yq := pq.FtoQ(xf), pq.FtoQ(yf)
		ps[i] = pq.XYtoP(xq, yq)
	}
	//
	T := time.Now()
	lower, upper := pq.ParConvHull2q(ncpu, ps)
	TT := time.Since(T)
	//
	fmt.Println("Lower#", len(lower))
	fmt.Println("Upper#", len(upper))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " Ready; T=" + TT.String())
}

// genI2 generates a random point on the unit square [-1/2,+1/2]².
func genI2() (x, y float64) {
	x = rand.Float64() - 0.5
	y = rand.Float64() - 0.5
	return
}

// genB2 generates a random point on the unit disc.
func genB2() (x, y float64) {
	for {
		x = 2*rand.Float64() - 1
		y = 2*rand.Float64() - 1
		if x*x+y*y < 1 {
			return
		}
	}
}
