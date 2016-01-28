package app

import (
	"fmt"
	r "regexp"
	"strings"
	"time"

	"DBWorker/lib"
	"DBWorker/utils"

	"bitbucket.org/liamstask/goose/lib/goose"
)

const TokenExpr = "^[a-zA-Z0-9]+:[a-zA-Z0-9]+"

type DbSetup interface {
	VerfiyTokens()
}

type Pairs []string

type Tokens map[string]string

func (db *DBConfig) DB_Setup(file utils.File, tokenPairs Pairs) (err *lib.Error) {
	//replace the tokens with the args

	fileContents, err := formatFileContents(file, tokenPairs)
	fmt.Printf("print:%v", fileContents)

	if err != nil {
		return err
	}

	//run it against psql client

	//handle the results
	return nil
}

func verifyPair(pair string) bool {
	matched, _ := r.Match(TokenExpr, pair)
	return matched
}

func makeTokens(pairs Pairs) (tokens Tokens, err *lib.Error) {
	expr, _ := r.Compile("[a-zA-Z0-9]")
	for _, val := range pairs {
		if verifyPair(val) {
			token := strings.Split(val, ":")
			tokens[token[0]] = token[1]
		} else {
			err = lib.NewError(lib.TokenErrorr, "db setup", "token syntax was incorrect")
			tokens = nil
			return
		}
	}
	return
}

func stringFormat(src string, tokens Tokens) string {
	for key, val := range tokens {
		src = strings.Replace(src, ":"+key, val, 1)
	}
	return src
}

func formatFileContents(file utils.File, tokenPairs Pairs) (fileContents string, err *lib.Error) {
	fileContents, err = utils.GetFileContents(file)
	if err != nil {
		return
	}

	tokens, err := makeTokens(tokenPairs)
	return stringFormat(fileContents, tokens)
}
