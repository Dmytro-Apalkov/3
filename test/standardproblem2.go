// +build ignore

package main

// Micromagnetic standard problem 2 according to
// http://www.ctcms.nist.gov/~rdm/mumag.org.html

import (
	"code.google.com/p/mx3/cuda"
	. "code.google.com/p/mx3/engine"
	"fmt"
	"log"
	"math"
)

func main() {
	cuda.Init()
	Init()
	defer Close()

	// Msat and Aex should not matter
	Msat = Const(1000e3)
	Aex = Const(10e-12)

	Alpha = Const(1.0)
	lex := math.Sqrt(Aex() / (0.5 * Mu0 * Msat() * Msat()))
	I := 30. // we test the solution for this d/lex value
	d := I * lex
	Sz, Sy, Sx := 0.1*d, d, 5*d

	nz, ny, nx := 1, 1, 1
	for Sy/float64(ny) > 0.75*lex {
		ny *= 2
	}
	for Sx/float64(nx) > 0.75*lex {
		nx *= 2
	}

	SetMesh(nx, ny, nz, Sx/float64(nx), Sy/float64(ny), Sz/float64(nz))

	M.Set(Uniform(1, 0.1, 0))

	Solver.Maxdt = 1e-12
	Run(2e-9)
	Solver.MaxErr = 1e-5
	Run(2e-9)

	m := M.Average()
	fmt.Println("remanent m for d/lex=", I, ":", m)
	// check against solution from Donahue:
	expect(m[X], 0.9627)
	expect(m[Y], 0.0756)
	expect(m[Z], 0)
	fmt.Println("OK")
}

func expect(have, want float64) {
	if math.Abs(have-want) > 1e-3 {
		log.Fatalln("have:", have, "want:", want)
	}
}
