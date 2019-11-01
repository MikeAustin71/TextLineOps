package textlinebuilder

type TextFieldSpec interface {
	FieldType() string
}

type IntegerTextSpec struct {
	NumericValue          int
	NumericFieldSpec      string
	NumericFieldLength    int
	NumericPadChar        rune
	NumericPosition       FieldPositionSpec
}

func (intSpec *IntegerTextSpec) CopyOut() IntegerTextSpec {

	newInt := IntegerTextSpec{}
	newInt.NumericValue = intSpec.NumericValue
	newInt.NumericFieldSpec = intSpec.NumericFieldSpec
	newInt.NumericFieldLength = intSpec.NumericFieldLength
	newInt.NumericPadChar = intSpec.NumericPadChar

	return newInt
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (intSpec IntegerTextSpec) FieldType() string {
	return "IntegerTextSpec"
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (intSpec IntegerTextSpec) NewPtr() *IntegerTextSpec {

	anotherIntTxtSpec := IntegerTextSpec{}

	return &anotherIntTxtSpec
}


type LineSpec struct {
	LineChar         rune
	LineLength       int
}

// CopyOut - Returns a deep copy of LineSpec
// data fields.
func (lineSpec *LineSpec) CopyOut() LineSpec {

	newLineSpec := LineSpec{}
	newLineSpec.LineChar = lineSpec.LineChar
	newLineSpec.LineLength = lineSpec.LineLength

	return newLineSpec
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (lineSpec LineSpec) FieldType() string {
	return "LineSpec"
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (lineSpec LineSpec) NewPtr() *LineSpec {

	anotherLineSpec := LineSpec{}

	return &anotherLineSpec
}


type MarginFieldSpec struct {
	MarginStr           string
	MarginLength        int
	MarginChar          rune
}

func (margin *MarginFieldSpec) CopyOut() MarginFieldSpec {

	newMargin := MarginFieldSpec{}

	newMargin.MarginStr = margin.MarginStr
	newMargin.MarginChar = margin.MarginChar
	newMargin.MarginLength = margin.MarginLength

	return newMargin
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (margin MarginFieldSpec) FieldType() string {
	return "MarginFieldSpec"
}

// NewPtr - Returns a pointer to a new MarginFieldSpec
// instance.
func (margin MarginFieldSpec) NewPtr() *MarginFieldSpec {

	newMarginField := MarginFieldSpec{}

	return &newMarginField
}


type NewLineSpec struct {
	AddNewLine        bool
}

func (newLine *NewLineSpec) CopyOut() NewLineSpec {
	newNewLine := NewLineSpec{}

	newNewLine.AddNewLine = newLine.AddNewLine

	return newNewLine
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (newLine NewLineSpec) FieldType() string {
	return "NewLineSpec"
}


// NewPtr - Returns a pointer to a new NewLineSpec
// instance.
func (newLine NewLineSpec) NewPtr() *NewLineSpec {

	newLineSpec := NewLineSpec{}

	return &newLineSpec
}

type NumericIntFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	NumberSpec            IntegerTextSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  NewLineSpec
}

// CopyOut - Returns a deep copy of the current NumericIntFieldSpec
// instance.
func (numIntField *NumericIntFieldSpec) CopyOut() NumericIntFieldSpec {

	newNumField := NumericIntFieldSpec{}

	newNumField.LeftMargin = numIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = numIntField.LeftSpacer.CopyOut()
	newNumField.NumberSpec = numIntField.NumberSpec.CopyOut()
	newNumField.NumericPosition = numIntField.NumericPosition
	newNumField.RightSpacer = numIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = numIntField.TerminateWithNewLine.CopyOut()

	return newNumField
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (numIntField NumericIntFieldSpec) FieldType() string {
	return "NumericIntFieldSpec"
}

// NewPtr - Returns a pointer to a new NumericIntFieldSpec instance.
func (numIntField NumericIntFieldSpec) NewPtr() *NumericIntFieldSpec {

	newNumIntSpec := NumericIntFieldSpec{}

	return &newNumIntSpec
}


type StringFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	StrTxt                StringTextSpec
	StrPosition           FieldPositionSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  NewLineSpec
}

func (strField *StringFieldSpec) CopyOut() StringFieldSpec {

	newStrField := StringFieldSpec{}

	newStrField.LeftMargin = strField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = strField.LeftSpacer.CopyOut()
	newStrField.StrTxt = strField.StrTxt.CopyOut()
	newStrField.StrPosition = strField.StrPosition
	newStrField.RightSpacer = strField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = strField.TerminateWithNewLine

	return newStrField
}

// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (strField StringFieldSpec) FieldType() string {
	return "StringFieldSpec"
}


// NewPtr - Returns a pointer to a new StringFieldSpec instance.
func (strField StringFieldSpec) NewPtr() *StringFieldSpec {

	newStrFieldSpec := StringFieldSpec{}

	return &newStrFieldSpec
}


type StringTextSpec struct {
	StrValue              string
	StrFieldLength        int
	StrPadChar            rune
}

// CopyOut - Returns a deep copy of the current StringTextSpec
// instance.
func (strTxtSpec *StringTextSpec) CopyOut() StringTextSpec {

	newStrSpec := StringTextSpec{}
	newStrSpec.StrValue = strTxtSpec.StrValue
	newStrSpec.StrFieldLength = strTxtSpec.StrFieldLength
	newStrSpec.StrPadChar = strTxtSpec.StrPadChar

	return newStrSpec
}


// FieldType - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (strTxtSpec StringTextSpec) FieldType() string {
	return "StringTextSpec"
}

// NewPtr - Returns a pointer to a new StringTextSpec instance.
func (strTxtSpec StringTextSpec) NewPtr() *StringTextSpec {

	newStrTxtSpec := StringTextSpec{}

	return &newStrTxtSpec
}
