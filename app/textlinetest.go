package main

import (
	"fmt"
	"github.com/MikeAustin71/pathfileopsgo/pathfileops/v2"
	"local.com/amarillomike/TextLineOps/textlinebuilder"
	"os"
)

func main() {

	ePrefix := "main() "

	dMgr, err := mainTest{}.getCurrentDirectory(ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}


	var fMgr pathfileops.FileMgr

	fMgr, err = mainTest{}.createOutputFile(dMgr, ePrefix)


	txtLnAry := make([]textlinebuilder.TextFieldSpec, 0)

	blankLines := textlinebuilder.BlankLinesSpec{}

	blankLines.NumBlankLines = 3

	txtLnAry = append(txtLnAry, blankLines)

	stringField := textlinebuilder.StringFieldSpec{}

	txtLnAry = append(txtLnAry, stringField )

	fmt.Printf("Testing Text Line Builder")

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
			"output" + string(os.PathSeparator)

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