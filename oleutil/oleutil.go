package oleutil

import "ole"
import "os"
import "unsafe"

func CreateDispatch(progId string) (dispatch *ole.IDispatch, err os.Error) {
	var clsid *ole.GUID
	clsid, err = ole.CLSIDFromProgID(progId)
	if err != nil {
		clsid, err = ole.CLSIDFromString(progId)
		if err != nil {
			return
		}
	}

	var unknown *ole.IUnknown
	unknown, err = ole.CreateInstance(clsid, ole.IID_IDispatch)
	if err != nil {
		return
	}
	dispatch = (*ole.IDispatch)(unsafe.Pointer(unknown))
	return
}

func CallMethod(disp *ole.IDispatch, name string, params ...interface{}) (result *ole.VARIANT, err os.Error) {
	var dispid []int32
	dispid, err = disp.GetIDsOfName([]string{name})
	if err != nil {
		return
	}
	result, err = disp.Invoke(dispid[0], ole.DISPATCH_METHOD, params...)
	return
}

func GetProperty(disp *ole.IDispatch, name string, params ...interface{}) (result *ole.VARIANT, err os.Error) {
	var dispid []int32
	dispid, err = disp.GetIDsOfName([]string{name})
	if err != nil {
		return
	}
	result, err = disp.Invoke(dispid[0], ole.DISPATCH_PROPERTYGET, params...)
	return
}

func PutProperty(disp *ole.IDispatch, name string, params ...interface{}) (result *ole.VARIANT, err os.Error) {
	var dispid []int32
	dispid, err = disp.GetIDsOfName([]string{name})
	if err != nil {
		return
	}
	result, err = disp.Invoke(dispid[0], ole.DISPATCH_PROPERTYPUT, params...)
	return
}
