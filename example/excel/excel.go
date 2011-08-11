package main

import "github.com/mattn/go-ole"
import "github.com/mattn/go-ole/oleutil"
import "syscall"

func main() {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("Excel.Application")
	excel, _ := unknown.QueryInterface(ole.IID_IDispatch)
	oleutil.PutProperty(excel, "Visible", true)
	result, _ := oleutil.GetProperty(excel, "Workbooks")
	workbooks := result.ToIDispatch()
	result, _ = oleutil.CallMethod(workbooks, "Add", nil)
	workbook := result.ToIDispatch()
	result, _ = oleutil.GetProperty(workbook, "Worksheets", 1)
	worksheet := result.ToIDispatch()
	result, _ = oleutil.GetProperty(worksheet, "Cells", 1, 1)
	cell := result.ToIDispatch()
	oleutil.PutProperty(cell, "Value", 12345)

	syscall.Sleep(2000000000)

	oleutil.PutProperty(workbook, "Saved", true)
	oleutil.CallMethod(excel, "Quit")
	excel.Release()
}
