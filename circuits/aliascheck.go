package circuits

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
)

func AliasCheck(api frontend.API, in []frontend.Variable) {
	q := api.Compiler().Field()
	q.Sub(q, big.NewInt(1))
	api.AssertIsEqual(CompConstant(api, in, q), 0)
}
