/**
* Server that runs on http://localhost:8080
**/
package main
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"strconv"
	"math"
	"encoding/json"
)
/**
* Struct to convert the fibonacci sequence found into a JSON object
**/
type FibonacciSequenceResponse struct {
	Sequence []int `json:sequence"`
}

/**
* Determines if an integer overflow will happen if this fibonacci value is found
* @Params: int, int 
* @Returns: bool
**/
func overflowed(number int, number2 int) bool {
	if ((number > 0) && (number2 > math.MaxUint32 - number)) {
		return true
	}
	return false
}

/**
* Determines the fibonacci sequence given the digit from the client
* @Params: int
* @Returns: []int
**/
func FibonacciAlgorithm(digits int) []int {
	var bad []int
	var fibonacciValues = make([]int,  digits)
	fibonacciValues[0] = 0;
	if digits == 1 {
		return fibonacciValues
	}
	fibonacciValues[1] = 1;
	for i:= 2; i < digits; i++{
		if overflowed(fibonacciValues[i-1], fibonacciValues[i-2]) {
			return bad
		}
		fibonacciValues[i] = fibonacciValues[i - 1] + fibonacciValues[i - 2];
	}
	return fibonacciValues
}

/**
* The router that retrieves the client digit passed, determines if it is good and passes
* a result to the user
* @Params: http.ReponseWriter, http.Request, httprouter.Params
**/
func FibonacciSequence(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	digits, err := strconv.Atoi(ps.ByName("digitsToParse"))
	if err != nil {
		http.Error(w, "Value given must be an number greater than 0 and less than "+strconv.Itoa(math.MaxUint32)+"; "+ ps.ByName("digitsToParse")+ " is invalid.", http.StatusUnprocessableEntity)
		return
	}
	if digits <= 0 {
		http.Error(w, "Number given needs to be an integer greater than 0.", http.StatusUnprocessableEntity)
		return
	}

	value := FibonacciAlgorithm(digits)
	if len(value) == 0 {
		http.Error(w, "Fibonacci values were too large. Try a digit smaller than 50.", http.StatusUnprocessableEntity)
		return
	}
	sequenceToReturn, _ := json.Marshal(value)
	w.Header().Set("Content-Type", "application/json")
  	w.Write(sequenceToReturn)
	fmt.Println(string(sequenceToReturn))
}

func main() {
	router := httprouter.New()
	router.GET("/api", Index)
	router.GET("/api/hello/:name", Hello)
	router.GET("/api/fibonacci/:digitsToParse", FibonacciSequence)
	log.Fatal(http.ListenAndServe(":8080", router))
}
