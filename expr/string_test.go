// SPDX-License-Identifier: Apache-2.0

package expr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/substrait-io/substrait-go/expr"
	"github.com/substrait-io/substrait-go/types"
)

func MustLiteral(l expr.Literal, err error) expr.Literal {
	if err != nil {
		panic(err)
	}
	return l
}

func TestLiteralToString(t *testing.T) {
	tests := []struct {
		t   expr.Literal
		exp string
	}{
		{&expr.PrimitiveLiteral[int16]{Value: 0, Type: &types.Int16Type{}}, "i16(0)"},
		{expr.NewPrimitiveLiteral[int8](0, true), "i8?(0)"},
		{expr.NewNestedLiteral(expr.ListLiteralValue{
			expr.NewNestedLiteral(expr.MapLiteralValue{
				{
					Key:   expr.NewPrimitiveLiteral("foo", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
				{
					Key:   expr.NewPrimitiveLiteral("baz", false),
					Value: expr.NewFixedCharLiteral(types.FixedChar("bar"), false),
				},
			}, true),
		}, true), "list?<map?<string,char<3>>>([map?<string,char<3>>([{string(foo) char<3>(bar)} {string(baz) char<3>(bar)}])])"},
		{MustLiteral(expr.NewLiteral(float32(1.5), false)), "fp32(1.5)"},
		{MustLiteral(expr.NewLiteral(&types.VarChar{Value: "foobar", Length: 7}, true)), "varchar?<7>(foobar)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "precisiontimestamp?<0>(123456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "precisiontimestamp?<3>(123456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "precisiontimestamp?<6>(123456)"},
		{expr.NewPrecisionTimestampLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "precisiontimestamp?<9>(123456)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionSeconds, types.NullabilityNullable), "precisiontimestamptz?<0>(123456)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMilliSeconds, types.NullabilityNullable), "precisiontimestamptz?<3>(123456)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionMicroSeconds, types.NullabilityNullable), "precisiontimestamptz?<6>(123456)"},
		{expr.NewPrecisionTimestampTzLiteral(123456, types.PrecisionNanoSeconds, types.NullabilityNullable), "precisiontimestamptz?<9>(123456)"},
	}

	for _, tt := range tests {
		t.Run(tt.exp, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.t.String())
		})
	}
}
