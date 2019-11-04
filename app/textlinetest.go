package main

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/TextLineOps/textlinebuilder"
	"os"
	"strings"
)

func main() {

	ePrefix := "main() "

	dMgr, err := mainTest{}.getCurrentDirectory(ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	var b strings.Builder

	b, err = mainTest{}.createTextLines(ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}


	var fMgr pathfileops.FileMgr

	fMgr, err = mainTest{}.createOutputFile(dMgr, ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}


	errs := make([]error, 0)

	_, err = fMgr.WriteBytesToFile([]byte(b.String()))

	if err != nil {
		err2 := fmt.Errorf("%v", err.Error())

		errs = append(errs, err2)
	}

	err = fMgr.FlushBytesToDisk()

	if err != nil {
		err2 := fmt.Errorf("%v", err.Error())

		errs = append(errs, err2)
	}

	err = fMgr.CloseThisFile()

	if err != nil {
		err2 := fmt.Errorf("%v", err.Error())

		errs = append(errs, err2)
	}

	err = pathfileops.FileHelper{}.ConsolidateErrors(errs)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmt.Printf(" ** SUCCESS!! **")
	fmt.Printf("Text Line Builder")
}

type mainTest struct {
	input      string
	output     string
}


// getCurrentDirectory() - Returns the current directory
// where this executable resides.
func (mtest mainTest) getCurrentDirectory(ePrefix string) (pathfileops.DirMgr, error) {

	ePrefix += "mainTest.getCurrentDirectory() "
	dMgr := pathfileops.DirMgr{}
	currDir, err := pathfileops.FileHelper{}.GetCurrentDir()

	if err != nil {
		return dMgr, fmt.Errorf(ePrefix +
			"\nError returned by pathfileops.FileHelper{}.GetCurrentDir()\n" +
			"Error='%v'\n", err.Error())
	}

	dMgr, err = pathfileops.DirMgr{}.New(currDir)

	if err != nil {
		return dMgr, fmt.Errorf(ePrefix +
			"\nError returned by DirMgr{}.New(currDir)\n" +
			"Error='%v'\n", err.Error())
	}

	return dMgr, nil
}

func (mtest mainTest) createOutputFile(
	curDir pathfileops.DirMgr,
	ePrefix string) (pathfileops.FileMgr, error) {

		ePrefix += "mainTest.createOutputFile() "

		outputFilePathName := curDir.GetAbsolutePathWithSeparator() +
			"output" + string(os.PathSeparator) + "textlinebuilder.txt"

		f, err := pathfileops.FileMgr{}.New(outputFilePathName)

		if err != nil {
			return pathfileops.FileMgr{},
				fmt.Errorf(ePrefix +
					"\nError returned by FileMgr{}.New(outputFilePathName)\n" +
					"Error='%v'\n", err.Error())
		}

	var err2 error

	err = f.IsFileMgrValid(ePrefix)

	if err != nil {
		return f, err
	}

	fileExists, err2 := f.DoesThisFileExist()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	if fileExists {

		err2 = f.DeleteThisFile()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		fileExists, err2 = f.DoesThisFileExist()

		if err2 != nil {
			err = fmt.Errorf(ePrefix+"%v", err2.Error())
			return f, err
		}

		if fileExists {
			err = fmt.Errorf(ePrefix+"Existing Output File FAILED to Delete! "+
				"Output File= '%v' ", f.GetAbsolutePathFileName())
			return f, err
		}

	}

	err2 = f.CreateThisFile()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err2 = f.OpenThisFileReadWrite()

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"%v", err2.Error())
		return f, err
	}

	err = nil

	return f, err
}

// createTextLines - Creates the test lines of text and returns them in
// a string builder.
func (mtest mainTest) createTextLines(ePrefix string) (strings.Builder, error) {

	ePrefix += "mainTest.createTextLines() "

	b := strings.Builder{}

	b.Grow(2048)


	blankLines := textlinebuilder.BlankLinesSpec{NumBlankLines:2}

	leftMargin := textlinebuilder.MarginSpec{MarginStr:"", MarginLength: 2, MarginChar: ' '}
	leftSpacer := textlinebuilder.MarginSpec{}
	labelStr := textlinebuilder.StringSpec{StrValue:"SomeTitle",StrFieldLength:60,StrPadChar:' ',StrPosition: textlinebuilder.FieldPos.Center()}
	rightSpacer := textlinebuilder.MarginSpec{}
	newLine := textlinebuilder.NewLineSpec{AddNewLine:true}

	err := textlinebuilder.TextLineBuilder{}.Build(&b, ePrefix,
		blankLines,
		leftMargin,
		leftSpacer,
		labelStr,
		rightSpacer,
		newLine)

	if err != nil {
		return strings.Builder{}, err
	}

	blankLines.NumBlankLines = 0

	lineSpec := textlinebuilder.LineSpec{
		LineChar:         '*',
		LineLength:       9,
		LineFieldLength:  60,
		LineFieldPadChar: ' ',
		LinePosition:     textlinebuilder.FieldPos.Center(),
	}



	lineBreak := textlinebuilder.LineBreakField{
		CreateLineBreak:      true,
		LeadingBlankLines:    blankLines.CopyOut(),
		LeftMargin:           leftMargin.CopyOut(),
		LeftSpacer:           leftSpacer.CopyOut(),
		LineSpec:             lineSpec.CopyOut(),
		RightSpacer:          rightSpacer.CopyOut(),
		TerminateWithNewLine: newLine.CopyOut(),
		TrailingBlankLines:   textlinebuilder.BlankLinesSpec{NumBlankLines:2},
	}

	err = textlinebuilder.TextLineBuilder{}.Build(&b,
		ePrefix,
		lineBreak)

	if err != nil {
		return strings.Builder{}, err
	}

	return b, nil
}