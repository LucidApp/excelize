// Copyright 2016 - 2023 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.16 or later.

package excelize

import (
	"errors"
	"fmt"
)

// newInvalidColumnNameError defined the error message on receiving the
// invalid column name.
func newInvalidColumnNameError(col string) error {
	return fmt.Errorf("invalid column name %q", col)
}

// newInvalidRowNumberError defined the error message on receiving the invalid
// row number.
func newInvalidRowNumberError(row int) error {
	return fmt.Errorf("invalid row number %d", row)
}

// newInvalidCellNameError defined the error message on receiving the invalid
// cell name.
func newInvalidCellNameError(cell string) error {
	return fmt.Errorf("invalid cell name %q", cell)
}

// newInvalidExcelDateError defined the error message on receiving the data
// with negative values.
func newInvalidExcelDateError(dateValue float64) error {
	return fmt.Errorf("invalid date value %f, negative values are not supported", dateValue)
}

// newInvalidNameError defined the error message on receiving the invalid
// defined name or table name.
func newInvalidNameError(name string) error {
	return fmt.Errorf("invalid name %q, the name should be starts with a letter or underscore, can not include a space or character, and can not conflict with an existing name in the workbook", name)
}

// newUnsupportedChartType defined the error message on receiving the chart
// type are unsupported.
func newUnsupportedChartType(chartType ChartType) error {
	return fmt.Errorf("unsupported chart type %d", chartType)
}

// newUnzipSizeLimitError defined the error message on unzip size exceeds the
// limit.
func newUnzipSizeLimitError(unzipSizeLimit int64) error {
	return fmt.Errorf("unzip size exceeds the %d bytes limit", unzipSizeLimit)
}

// newInvalidStyleID defined the error message on receiving the invalid style
// ID.
func newInvalidStyleID(styleID int) error {
	return fmt.Errorf("invalid style ID %d", styleID)
}

// newFieldLengthError defined the error message on receiving the field length
// overflow.
func newFieldLengthError(name string) error {
	return fmt.Errorf("field %s must be less than or equal to 255 characters", name)
}

// newCellNameToCoordinatesError defined the error message on converts
// alphanumeric cell name to coordinates.
func newCellNameToCoordinatesError(cell string, err error) error {
	return fmt.Errorf("cannot convert cell %q to coordinates: %v", cell, err)
}

// newNoExistSheetError defined the error message on receiving the non existing
// sheet name.
func newNoExistSheetError(name string) error {
	return fmt.Errorf("sheet %s does not exist", name)
}

// newNotWorksheetError defined the error message on receiving a sheet which
// not a worksheet.
func newNotWorksheetError(name string) error {
	return fmt.Errorf("sheet %s is not a worksheet", name)
}

// newStreamSetRowError defined the error message on the stream writer
// receiving the non-ascending row number.
func newStreamSetRowError(row int) error {
	return fmt.Errorf("row %d has already been written", row)
}

// newViewIdxError defined the error message on receiving a invalid sheet view
// index.
func newViewIdxError(viewIndex int) error {
	return fmt.Errorf("view index %d out of range", viewIndex)
}

var (
	// ErrStreamSetColWidth defined the error message on set column width in
	// stream writing mode.
	ErrStreamSetColWidth = errors.New("must call the SetColWidth function before the SetRow function")
	// ErrStreamSetPanes defined the error message on set panes in stream
	// writing mode.
	ErrStreamSetPanes = errors.New("must call the SetPanes function before the SetRow function")
	// ErrColumnNumber defined the error message on receive an invalid column
	// number.
	ErrColumnNumber = fmt.Errorf(`the column number must be greater than or equal to %d and less than or equal to %d`, MinColumns, MaxColumns)
	// ErrColumnWidth defined the error message on receive an invalid column
	// width.
	ErrColumnWidth = fmt.Errorf("the width of the column must be less than or equal to %d characters", MaxColumnWidth)
	// ErrOutlineLevel defined the error message on receive an invalid outline
	// level number.
	ErrOutlineLevel = errors.New("invalid outline level")
	// ErrCoordinates defined the error message on invalid coordinates tuples
	// length.
	ErrCoordinates = errors.New("coordinates length must be 4")
	// ErrExistsSheet defined the error message on given sheet already exists.
	ErrExistsSheet = errors.New("the same name sheet already exists")
	// ErrTotalSheetHyperlinks defined the error message on hyperlinks count
	// overflow.
	ErrTotalSheetHyperlinks = errors.New("over maximum limit hyperlinks in a worksheet")
	// ErrInvalidFormula defined the error message on receive an invalid
	// formula.
	ErrInvalidFormula = errors.New("formula not valid")
	// ErrAddVBAProject defined the error message on add the VBA project in
	// the workbook.
	ErrAddVBAProject = errors.New("unsupported VBA project")
	// ErrMaxRows defined the error message on receive a row number exceeds maximum limit.
	ErrMaxRows = errors.New("row number exceeds maximum limit")
	// ErrMaxRowHeight defined the error message on receive an invalid row
	// height.
	ErrMaxRowHeight = fmt.Errorf("the height of the row must be less than or equal to %d points", MaxRowHeight)
	// ErrImgExt defined the error message on receive an unsupported image
	// extension.
	ErrImgExt = errors.New("unsupported image extension")
	// ErrWorkbookFileFormat defined the error message on receive an
	// unsupported workbook file format.
	ErrWorkbookFileFormat = errors.New("unsupported workbook file format")
	// ErrMaxFilePathLength defined the error message on receive the file path
	// length overflow.
	ErrMaxFilePathLength = errors.New("file path length exceeds maximum limit")
	// ErrUnknownEncryptMechanism defined the error message on unsupported
	// encryption mechanism.
	ErrUnknownEncryptMechanism = errors.New("unknown encryption mechanism")
	// ErrUnsupportedEncryptMechanism defined the error message on unsupported
	// encryption mechanism.
	ErrUnsupportedEncryptMechanism = errors.New("unsupported encryption mechanism")
	// ErrUnsupportedHashAlgorithm defined the error message on unsupported
	// hash algorithm.
	ErrUnsupportedHashAlgorithm = errors.New("unsupported hash algorithm")
	// ErrUnsupportedNumberFormat defined the error message on unsupported number format
	// expression.
	ErrUnsupportedNumberFormat = errors.New("unsupported number format token")
	// ErrPasswordLengthInvalid defined the error message on invalid password
	// length.
	ErrPasswordLengthInvalid = errors.New("password length invalid")
	// ErrParameterRequired defined the error message on receive the empty
	// parameter.
	ErrParameterRequired = errors.New("parameter is required")
	// ErrParameterInvalid defined the error message on receive the invalid
	// parameter.
	ErrParameterInvalid = errors.New("parameter is invalid")
	// ErrDefinedNameScope defined the error message on not found defined name
	// in the given scope.
	ErrDefinedNameScope = errors.New("no defined name on the scope")
	// ErrDefinedNameDuplicate defined the error message on the same name
	// already exists on the scope.
	ErrDefinedNameDuplicate = errors.New("the same name already exists on the scope")
	// ErrCustomNumFmt defined the error message on receive the empty custom number format.
	ErrCustomNumFmt = errors.New("custom number format can not be empty")
	// ErrFontLength defined the error message on the length of the font
	// family name overflow.
	ErrFontLength = fmt.Errorf("the length of the font family name must be less than or equal to %d", MaxFontFamilyLength)
	// ErrFontSize defined the error message on the size of the font is invalid.
	ErrFontSize = fmt.Errorf("font size must be between %d and %d points", MinFontSize, MaxFontSize)
	// ErrSheetIdx defined the error message on receive the invalid worksheet
	// index.
	ErrSheetIdx = errors.New("invalid worksheet index")
	// ErrUnprotectSheet defined the error message on worksheet has set no
	// protection.
	ErrUnprotectSheet = errors.New("worksheet has set no protect")
	// ErrUnprotectSheetPassword defined the error message on remove sheet
	// protection with password verification failed.
	ErrUnprotectSheetPassword = errors.New("worksheet protect password not match")
	// ErrGroupSheets defined the error message on group sheets.
	ErrGroupSheets = errors.New("group worksheet must contain an active worksheet")
	// ErrDataValidationFormulaLength defined the error message for receiving a
	// data validation formula length that exceeds the limit.
	ErrDataValidationFormulaLength = fmt.Errorf("data validation must be 0-%d characters", MaxFieldLength)
	// ErrDataValidationRange defined the error message on set decimal range
	// exceeds limit.
	ErrDataValidationRange = errors.New("data validation range exceeds limit")
	// ErrCellCharsLength defined the error message for receiving a cell
	// characters length that exceeds the limit.
	ErrCellCharsLength = fmt.Errorf("cell value must be 0-%d characters", TotalCellChars)
	// ErrOptionsUnzipSizeLimit defined the error message for receiving
	// invalid UnzipSizeLimit and UnzipXMLSizeLimit.
	ErrOptionsUnzipSizeLimit = errors.New("the value of UnzipSizeLimit should be greater than or equal to UnzipXMLSizeLimit")
	// ErrSave defined the error message for saving file.
	ErrSave = errors.New("no path defined for file, consider File.WriteTo or File.Write")
	// ErrAttrValBool defined the error message on marshal and unmarshal
	// boolean type XML attribute.
	ErrAttrValBool = errors.New("unexpected child of attrValBool")
	// ErrSparklineType defined the error message on receive the invalid
	// sparkline Type parameters.
	ErrSparklineType = errors.New("parameter 'Type' must be 'line', 'column' or 'win_loss'")
	// ErrSparklineLocation defined the error message on missing Location
	// parameters
	ErrSparklineLocation = errors.New("parameter 'Location' is required")
	// ErrSparklineRange defined the error message on missing sparkline Range
	// parameters
	ErrSparklineRange = errors.New("parameter 'Range' is required")
	// ErrSparkline defined the error message on receive the invalid sparkline
	// parameters.
	ErrSparkline = errors.New("must have the same number of 'Location' and 'Range' parameters")
	// ErrSparklineStyle defined the error message on receive the invalid
	// sparkline Style parameters.
	ErrSparklineStyle = errors.New("parameter 'Style' must between 0-35")
	// ErrWorkbookPassword defined the error message on receiving the incorrect
	// workbook password.
	ErrWorkbookPassword = errors.New("the supplied open workbook password is not correct")
	// ErrSheetNameInvalid defined the error message on receive the sheet name
	// contains invalid characters.
	ErrSheetNameInvalid = errors.New("the sheet can not contain any of the characters :\\/?*[or]")
	// ErrSheetNameSingleQuote defined the error message on the first or last
	// character of the sheet name was a single quote.
	ErrSheetNameSingleQuote = errors.New("the first or last character of the sheet name can not be a single quote")
	// ErrSheetNameBlank defined the error message on receive the blank sheet
	// name.
	ErrSheetNameBlank = errors.New("the sheet name can not be blank")
	// ErrSheetNameLength defined the error message on receiving the sheet
	// name length exceeds the limit.
	ErrSheetNameLength = fmt.Errorf("the sheet name length exceeds the %d characters limit", MaxSheetNameLength)
	// ErrNameLength defined the error message on receiving the defined name or
	// table name length exceeds the limit.
	ErrNameLength = fmt.Errorf("the name length exceeds the %d characters limit", MaxFieldLength)
	// ErrExistsTableName defined the error message on given table already exists.
	ErrExistsTableName = errors.New("the same name table already exists")
	// ErrCellStyles defined the error message on cell styles exceeds the limit.
	ErrCellStyles = fmt.Errorf("the cell styles exceeds the %d limit", MaxCellStyles)
	// ErrUnprotectWorkbook defined the error message on workbook has set no
	// protection.
	ErrUnprotectWorkbook = errors.New("workbook has set no protect")
	// ErrUnprotectWorkbookPassword defined the error message on remove workbook
	// protection with password verification failed.
	ErrUnprotectWorkbookPassword = errors.New("workbook protect password not match")
)
