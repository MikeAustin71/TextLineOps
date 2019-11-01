package main

import (
	"fmt"
	"local.com/amarillomike/TextLineOps/textlinebuilder"
)

func main() {

	txtLnAry := make([]textlinebuilder.TextFieldSpec, 0)

	blankLines := textlinebuilder.BlankLinesSpec{}

	blankLines.NumBlankLines = 3

	txtLnAry = append(txtLnAry, blankLines)

	stringField := textlinebuilder.StringFieldSpec{}


	txtLnAry = append(txtLnAry, )

	fmt.Printf("Testing Text Line Builder")


}
