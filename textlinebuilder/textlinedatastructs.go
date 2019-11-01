package textlinebuilder


type BlankLinesSpec struct {
	NumBlankLines        int
}

func (blkLine *BlankLinesSpec) CopyOut() BlankLinesSpec {
	newBlankLineSpec := BlankLinesSpec{}

	newBlankLineSpec.NumBlankLines = blkLine.NumBlankLines

	return newBlankLineSpec
}

type LineBreakSpec struct {
	CreateLineBreak       bool
	LeadingBlankLines     int
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	LineBreakChar         rune
	LineBreakLength       int
	RightMargin           MarginFieldSpec
	TerminateWithNewLine  bool
	TrailingBlankLines    int
}

func (lineBrk *LineBreakSpec) CopyOut() LineBreakSpec {

	newLineBreak := LineBreakSpec{}

	newLineBreak.CreateLineBreak = lineBrk.CreateLineBreak
	newLineBreak.LeftMargin = lineBrk.LeftMargin.CopyOut()
	newLineBreak.LeftSpacer = lineBrk.LeftSpacer.CopyOut()
	newLineBreak.LineBreakChar = lineBrk.LineBreakChar
	newLineBreak.LineBreakLength = lineBrk.LineBreakLength
	newLineBreak.RightMargin = lineBrk.RightMargin.CopyOut()
	newLineBreak.TerminateWithNewLine = lineBrk.TerminateWithNewLine

	return newLineBreak
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

type NewLineSpec struct {
	AddNewLine        bool
}

func (newLine *NewLineSpec) CopyOut() NewLineSpec {
	newNewLine := NewLineSpec{}

	newNewLine.AddNewLine = newLine.AddNewLine

	return newNewLine
}

type NumericIntFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	NumericValue          int
	NumericFieldSpec      string
	NumericFieldLength    int
	NumericPadChar        rune
	NumericPosition       FieldPositionSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  bool
}

func (numIntField *NumericIntFieldSpec) CopyOut() NumericIntFieldSpec {

	newNumField := NumericIntFieldSpec{}

	newNumField.LeftMargin = numIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = numIntField.LeftSpacer.CopyOut()

	newNumField.NumericValue = numIntField.NumericValue
	newNumField.NumericFieldSpec = numIntField.NumericFieldSpec
	newNumField.NumericFieldLength = numIntField.NumericFieldLength
	newNumField.NumericPadChar = numIntField.NumericPadChar
	newNumField.NumericPosition = numIntField.NumericPosition
	newNumField.RightSpacer = numIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = numIntField.TerminateWithNewLine

	return newNumField
}

type StringFieldSpec struct {
	LeftMargin            MarginFieldSpec
	LeftSpacer            MarginFieldSpec
	StrValue              string
	StrFieldLength        int
	StrPadChar            rune
	StrPosition           FieldPositionSpec
	RightSpacer           MarginFieldSpec
	TerminateWithNewLine  bool
}

func (strField *StringFieldSpec) CopyOut() StringFieldSpec {

	newStrField := StringFieldSpec{}

	newStrField.LeftMargin = strField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = strField.LeftSpacer.CopyOut()
	newStrField.StrValue = strField.StrValue
	newStrField.StrFieldLength = strField.StrFieldLength
	newStrField.StrPadChar = strField.StrPadChar
	newStrField.StrPosition = strField.StrPosition
	newStrField.RightSpacer = strField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = strField.TerminateWithNewLine

	return newStrField
}


type TwoLabelStrLine struct {
	LeadingBlankLines  int
	TopLineBreak       LineBreakSpec
	Label1             StringFieldSpec
	Label2             StringFieldSpec
	BottomLineBreak    LineBreakSpec
	TrailingBlankLines int
}

type OneLabelOneIntLine struct {
	LeadingBlankLines  int
	TopLineBreak       LineBreakSpec
	Label1             StringFieldSpec
	Number1            NumericIntFieldSpec
	BottomLineBreak    LineBreakSpec
	TrailingBlankLines int
}

func (labelIntLine *OneLabelOneIntLine) CopyOut() OneLabelOneIntLine {

	newLabelInt := OneLabelOneIntLine{}

	newLabelInt.LeadingBlankLines = labelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = labelIntLine.TopLineBreak.CopyOut()
	newLabelInt.Label1 = labelIntLine.Label1.CopyOut()
	newLabelInt.Number1 = labelIntLine.Number1.CopyOut()
	newLabelInt.BottomLineBreak = labelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = labelIntLine.TrailingBlankLines
	return newLabelInt
}

