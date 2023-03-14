package number

import (
	"encoding/json"
	"errors"
	"math"
	"math/big"
)

func FindPrimeNumbers(inputStr []byte) ([]bool, error) {
	inputs := make([]int, 0)
	unmarshalErr := json.Unmarshal(inputStr, &inputs)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	results := make([]bool, 0, len(inputs))
	for _, i := range inputs {
		if i <= 0 {
			return nil, errors.New("wrong input values")
		}
		b := big.NewInt(int64(i))
		results = append(results, b.ProbablyPrime(int(math.Log2(float64(i)))))
	}

	return results, nil
}
