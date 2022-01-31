//Patrick Jaime Simba

//package main produces a webapp that can be used to solve a system of 3 equations using Cramer's rule
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//main calls on the handler and solveSystem functions and sets the url
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/solve", solveSystem)
	http.ListenAndServe(":8080", nil)
}

//handler prints a welcome statement with instructions and an example URL to pass
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Patrick's WebApp! Input 12 comma separated coefficients into the URL following solve?coef= to solve a system of 3 equations.\n")
	fmt.Fprintf(w, "Example URL: http://localhost:8080/solve?coef=4,5,6,7,2,3,1,2,1,2,3,2")
}

//solveSystem finds the determinants for the system of 3 equations and solves for and prints x, y and z using Cramer's rule if there is only one solution; otherwise, it will print inconsistent - no solutions or dependent - multiple solutions.
func solveSystem(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	coeff := r.URL.Query().Get("coef")
	var stringSlice []string
	var coeffSlice []int

	stringSlice = strings.Split(coeff, ",")

	if len(stringSlice) != 12 {
		fmt.Fprintln(w, "Incorrect number of coefficients passed.")
		return
	}

	for key := range stringSlice {
		if val, err := strconv.Atoi(stringSlice[key]); err == nil {
			coeffSlice = append(coeffSlice, val)
		}
	}

	M := map[string]float64{
		"a1": float64(coeffSlice[0]),
		"b1": float64(coeffSlice[1]),
		"c1": float64(coeffSlice[2]),
		"d1": float64(coeffSlice[3]),
		"a2": float64(coeffSlice[4]),
		"b2": float64(coeffSlice[5]),
		"c2": float64(coeffSlice[6]),
		"d2": float64(coeffSlice[7]),
		"a3": float64(coeffSlice[8]),
		"b3": float64(coeffSlice[9]),
		"c3": float64(coeffSlice[10]),
		"d3": float64(coeffSlice[11]),
	}

	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", M["a1"], M["b1"], M["c1"], M["d1"])
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", M["a2"], M["b2"], M["c2"], M["d2"])
	fmt.Fprintf(w, "%vx + %vy + %vz = %v\n", M["a3"], M["b3"], M["c3"], M["d3"])

	var D, Dx, Dy, Dz float64

	//Cramer's rule for solving a system of 3 equations: https://courses.lumenlearning.com/ivytech-collegealgebra/chapter/using-cramers-rule-to-solve-a-system-of-three-equations-in-three-variables/
	D = (M["a1"] * M["b2"] * M["c3"]) + (M["b1"] * M["c2"] * M["a3"]) + (M["c1"] * M["a2"] * M["b3"]) - (M["a3"] * M["b2"] * M["c1"]) - (M["b3"] * M["c2"] * M["a1"]) - (M["c3"] * M["a2"] * M["b1"])
	Dx = (M["d1"] * M["b2"] * M["c3"]) + (M["b1"] * M["c2"] * M["d3"]) + (M["c1"] * M["d2"] * M["b3"]) - (M["d3"] * M["b2"] * M["c1"]) - (M["b3"] * M["c2"] * M["d1"]) - (M["c3"] * M["d2"] * M["b1"])
	Dy = (M["a1"] * M["d2"] * M["c3"]) + (M["d1"] * M["c2"] * M["a3"]) + (M["c1"] * M["a2"] * M["d3"]) - (M["a3"] * M["d2"] * M["c1"]) - (M["d3"] * M["c2"] * M["a1"]) - (M["c3"] * M["a2"] * M["d1"])
	Dz = (M["a1"] * M["b2"] * M["d3"]) + (M["b1"] * M["d2"] * M["a3"]) + (M["d1"] * M["a2"] * M["b3"]) - (M["a3"] * M["b2"] * M["d1"]) - (M["b3"] * M["d2"] * M["a1"]) - (M["d3"] * M["a2"] * M["b1"])

	if D == 0 && (Dx != 0 || Dy != 0 || Dz != 0) {
		fmt.Fprintln(w, "inconsistent - no solution")
	} else if D == 0 && (Dx == 0 && Dy == 0 && Dz == 0) {
		fmt.Fprintln(w, "dependent - with multiple solutions")
	} else {
		x := Dx / D
		y := Dy / D
		z := Dz / D
		fmt.Fprintln(w, "solution:")
		fmt.Fprintf(w, "x = %.2f, y = %.2f, z = %.2f\n", x, y, z)
	}
}
