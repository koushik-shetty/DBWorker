package utils

import (
	r "regexp"
	"strings"

	"DBWorker/lib"
)

const TokenExpr = "^[a-zA-Z0-9_]+:[a-zA-Z0-9_]+"

type Pair string
type Pairs []Pair
type Tokens map[string]string

func ToPairs(strings []string) (pairs Pairs) {
	for _, str := range strings {
		pair := Pair(str)
		pairs = append(pairs, pair)
	}
	return
}

func (p Pair) String() string {
	return string(p)
}

func (p Pair) isGood() bool {
	matched, _ := r.Match(TokenExpr, []byte(p))
	return matched
}

func (ps *Pairs) Verify() *lib.Error {
	for _, pair := range *ps {
		if !pair.isGood() {
			return lib.NewError(lib.TokenError, "VerifyPairs", "pairs not in the right format")
		}
	}
	return nil
}

func (ps *Pairs) ToTokens() Tokens {
	tokens := make(Tokens)
	for _, pair := range *ps {
		token := strings.Split(pair.String(), ":")
		tokens[token[0]] = token[1]
	}
	return tokens
}

func (t *Tokens) stringInterpolate(src string) string {
	for key, val := range *t {
		src = strings.Replace(src, ":"+key, val, -1)
	}
	return src
}
