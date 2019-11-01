package textlinebuilder

import (
	"errors"
	"fmt"
	"strings"
)


type TextLineBuilder struct {
	Input    string
	Output   string
}

func (txtBuilder TextLineBuilder) CenterInField(
	strValue string,
	strFieldLen int,
	padChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CenterInField() "


	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n" +
			"strValue='%v'\n" +
			"strValue length='%v'\n" +
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteString(strValue)\n" +
				"strValue string length == strFieldLen\n" +
				"strValue='%v'\n" +
				"Error='%v'\n", strValue, err.Error())
		}
	}

	if padChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'padChar' is ZERO!\n")
	}

	grossPad := strFieldLen - lenStr
	var leftPad, rightPad int

	leftPad = grossPad / 2
	rightPad = grossPad - leftPad

	for i:= 0; i < leftPad; i++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(left padChar)\n" +
				"padChar='%v'   i='%v'\n" +
				"Error='%v'\n", padChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned after left pad by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	for j:=0; j < rightPad; j++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(right padChar)\n" +
				"padChar='%v'   j='%v'\n" +
				"Error='%v'\n", padChar, j, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateBlankLines(
	blankLines BlankLinesSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateBlankLines() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if blankLines.NumBlankLines < 1 {
		return nil
	}

	var err error

	for i:=0; i < blankLines.NumBlankLines; i++ {

		_, err = b.WriteRune('\n')

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(newLine).\n" +
				"i='%v'\n" +
				"Error='%v'\n", i, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateIntegerText(
	intTxtSpec IntegerTextSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateIntegerText() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	numFmt := ""

	if len(intTxtSpec.NumericFieldSpec) == 0 {
		numFmt = "%d"
	} else {
		numFmt = intTxtSpec.NumericFieldSpec
	}

	numStr := fmt.Sprintf(numFmt, intTxtSpec.NumericValue)

	lenNumStr := len(numStr)

	if lenNumStr > intTxtSpec.NumericFieldLength {
		return fmt.Errorf(ePrefix +
			"\nError: Length of number string is greater than NumericFieldLength!\n" +
			"Number String='%v'\n" +
			"Number String Length='%v'\n" +
			"Numeric Field Length='%v'\n",
			numStr,
			lenNumStr,
			intTxtSpec.NumericFieldLength)
	}

	var err error

	if lenNumStr == intTxtSpec.NumericFieldLength {

		_, err = b.WriteString(numStr)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteString(numStr).\n" +
				"numStr='%v'\n" +
				"Error='%v'\n", numStr, err.Error())
		}
	}

	switch intTxtSpec.NumericPosition {
	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	default:

		err = errors.New(ePrefix +
			"\nError: Invalid parameter intTxtSpec.NumericPosition\n")
	}

	if err != nil {
		return err
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateLineBreak(
	lineBreak LineBreakSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateLineBreak() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if !lineBreak.CreateLineBreak {
		// Nothing to do
		return nil
	}

	var err error

	err = txtBuilder.CreateBlankLines(
		lineBreak.LeadingBlankLines,
		b,
		ePrefix)

	if err != nil {
		return err
	}


	err = txtBuilder.CreateMarginField(
		lineBreak.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginField(
		lineBreak.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLine(lineBreak.LineSpec, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginField(
		lineBreak.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLine(
		lineBreak.TerminateWithNewLine,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateBlankLines(
		lineBreak.TrailingBlankLines,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateLine(
	lineSpec LineSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateLine() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!\n")
	}

	if lineSpec.LineLength < 1 {
		return nil
	}


	if lineSpec.LineChar == 0 {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter lineSpec.LineChar is ZERO!\n")
	}

	var err error

	for i:=0; i < lineSpec.LineLength; i++ {

		_, err = b.WriteRune(lineSpec.LineChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(lineSpec.LineChar).\n" +
				"lineSpec.LineChar='%v'\n" +
				"Error='%v'\n", lineSpec.LineChar, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateMarginField(
	margin MarginFieldSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateMarginField() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	if len(margin.MarginStr) > 0 {
		_, err = b.WriteString(margin.MarginStr)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteString(margin.MarginStr)\n" +
				"margin.MarginStr='%v'\n" +
				"Error='%v'\n", margin.MarginStr, err.Error() )
		}
	}

	if margin.MarginLength == 0 {
		// Nothing to do
		return nil
	}

	if margin.MarginChar == 0 {
		return errors.New(ePrefix +
			"\nError: margin.MarginChar = 0\n")
	}

	xb := strings.Builder{}
	xb.Grow(margin.MarginLength + 2)

	for i:=0; i < margin.MarginLength; i++ {
		_, err = xb.WriteRune(margin.MarginChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by xb.WriteRune(margin.MarginChar)\n" +
				"margin.MarginChar='%v'\n" +
				"Error='%v'\n", margin.MarginChar, err.Error())
		}
	}

	_, err = b.WriteString(xb.String())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(xb.String())\n" +
			"xb.String()='%v'\n" +
			"Error='%v'\n", xb.String(), err.Error())
	}

		return nil
}

func (txtBuilder TextLineBuilder) CreateNewLine(
	newLineSpec NewLineSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateNewLine() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if !newLineSpec.AddNewLine {
		// Nothing to do
		return nil
	}

	var err error

		_, err = b.WriteRune('\n')

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(newline)\n" +
				"Error='%v'\n", err.Error())
		}

	return nil
}

func (txtBuilder TextLineBuilder) CreateNumericIntField(
	numSpec NumericIntFieldSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateNumericIntField() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateMarginField(
		numSpec.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}


	err = txtBuilder.CreateMarginField(
		numSpec.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	numFmt := ""

	if len(numSpec.NumericFieldSpec) == 0 {
		numFmt = "%d"
	} else {
		numFmt = numSpec.NumericFieldSpec
	}

	numStr := fmt.Sprintf(numFmt, numSpec.NumericValue)

	switch numSpec.NumericPosition {

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			numStr,numSpec.NumericFieldLength, numSpec.NumericPadChar, b, ePrefix)

	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			numStr,numSpec.NumericFieldLength, numSpec.NumericPadChar, b, ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			numStr,numSpec.NumericFieldLength, numSpec.NumericPadChar, b, ePrefix)

	default:

		err = errors.New(ePrefix +
			"\nError: Number Specification Field Postion is Invalid!\n")
	}

	err = txtBuilder.CreateMarginField(
		numSpec.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	if numSpec.TerminateWithNewLine {
		err = txtBuilder.CreateNewLine(1, b, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateOneLabelOneIntLine(
	oneLabelOneInt OneLabelOneIntLine,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzLogOps.CreateOneLabelOneIntLine() "


	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateNewLine(oneLabelOneInt.LeadingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreak(oneLabelOneInt.TopLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(oneLabelOneInt.Label1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNumericIntField(oneLabelOneInt.Number1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreak(oneLabelOneInt.BottomLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLine(oneLabelOneInt.TrailingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}


	return nil
}

func (txtBuilder TextLineBuilder) CreateTwoLabelStrLine(
	twoLabelStrLine TwoLabelStrLine,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateTwoLabelStrLine() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateNewLine(twoLabelStrLine.LeadingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreak(twoLabelStrLine.TopLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(twoLabelStrLine.Label1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(twoLabelStrLine.Label2, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreak(twoLabelStrLine.BottomLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLine(twoLabelStrLine.TrailingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// CreateStringField - Designed to handle StringFieldSpec specifications.
func (txtBuilder TextLineBuilder) CreateStringField(
	strSpec StringFieldSpec,
	b *strings.Builder,
	ePrefix string) error{

	ePrefix += "TextLineBuilder.CreateStringField() "

	if b == nil{
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateMarginField(
		strSpec.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}


	err = txtBuilder.CreateMarginField(
		strSpec.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	switch strSpec.StrPosition {

	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix +
			"\nError: strSpec.StrPosition is invalid!\n" +
			"StrPosition Value='%v'\n", strSpec.StrPosition.UtilityValue())
	}

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginField(
		strSpec.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	if strSpec.TerminateWithNewLine {
		err = txtBuilder.CreateNewLine(1,b, ePrefix)

		if err != nil {
			return err
		}
	}

	return nil
}

// LeftJustifyField - Left Justifies parameter 'str' in a field of length 'fieldLen'.
func (txtBuilder TextLineBuilder) LeftJustifyField(
	strValue string,
	strFieldLen int,
	trailingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrFmt.LeftJustifyField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n" +
			"strValue='%v'\n" +
			"strValue length='%v'\n" +
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"\nError returned by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	if lenStr == strFieldLen {
		return nil
	}

	if trailingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'trailingPadChar' is Zero!")
	}

	padLen := strFieldLen - lenStr

	for i:=0; i < padLen; i++ {

		_, err = b.WriteRune(trailingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(trailingPadChar)\n" +
				"trailingPadChar='%v'   i='%v'\n" +
				"Error='%v'\n", trailingPadChar, i, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) RightJustifyField(
	strValue string,
	strFieldLen int,
	leadingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {


	ePrefix += "TextLineBuilder.RightJustifyField() "

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix +
			"\nError: Input Parameter 'strValue' length exceeds 'fieldLen'.\n" +
			"strValue='%v'" +
			"strValue length='%v'\n" +
			"fieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error returned by b.WriteString(strValue)\n" +
				"strValue='%v'\n" +
				"Error='%v'\n", strValue, err.Error())
		}

		return nil
	}

	if leadingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'leadingPadChar' is ZERO!\n")
	}

	padLen := strFieldLen - lenStr

	for i:=0; i < padLen; i++ {

		_, err = b.WriteRune(leadingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix +
				"\nError returned by b.WriteRune(leadingPadChar)\n" +
				"leadingPadChar='%v'   i='%v'\n" +
				"Error='%v'\n", leadingPadChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned after padChar by b.WriteString(strValue)\n" +
			"strValue='%v'\n" +
			"Error='%v'\n", strValue, err.Error())
	}

	return nil
}