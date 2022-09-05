package translator

import (
	"fmt"
	"github.com/ettle/strcase"
	"modernc.org/cc/v3"
	"reflect"
)

type StructType struct {
	Ident string
	Elems []StructElem
}

type StructElem struct {
	Ident string
	Type  string
}

func convertStructElemType(t cc.Type) StructType {
	var st StructType

	var ses []StructElem
	st.Ident = strcase.ToPascal(t.Name().String())
	for i := 0; i < t.NumField(); i++ {
		var se StructElem
		f := t.FieldByIndex([]int{i})
		if f.Type().Kind() == cc.Struct || f.Type().Kind() == cc.Union {
			se.Ident = strcase.ToPascal(f.Name().String())
			se.Type = f.Type().Name().String()
		} else {
			s := detectTypes(f.Type())
			se.Ident = strcase.ToPascal(f.Name().String())
			se.Type = s
		}

		ses = append(ses, se)
	}

	st.Elems = ses
	return st
}

func detectTypes(t cc.Type) string {
	switch kind := t.Kind(); kind {
	case cc.UInt64, cc.UInt32, cc.UInt16, cc.UInt8:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.Int64, cc.Int32, cc.Int16, cc.Int8:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.SChar:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.UChar:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.Short:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.UShort:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.Int:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.UInt:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.Long:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.ULong:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.LongLong:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.ULongLong:
		s := setString(UintT(int(t.Size())), "")
		return s
	case cc.Float:
		s := setString(FloatT(int(t.Size())), "")
		return s
	case cc.Double:
		s := setString(FloatT(int(t.Size())), "")
		return s
	case cc.LongDouble:
		s := setString(FloatT(int(t.Size())), "")
		return s
	case cc.Char:
		s := setString(IntT(int(t.Size())), "")
		return s
	case cc.Bool:
		s := setString(reflect.TypeOf(true).Kind(), "")
		return s
	case cc.Ptr:
		s := setString(IntT(int(t.Elem().Size())), "*")
		return s
	case cc.Array:
		var typ string
		if t.Elem().Kind() == cc.Char {
			typ = "byte"
		} else {
			typ = detectTypes(t.Elem())
		}

		bT := fmt.Sprintf("[%d]", t.Len())

		at := fmt.Sprintf("%s%s", bT, typ)
		return at
	default:
		//panic(fmt.Errorf("%T, %s (%s)", t, kind, t.String()))
	}

	return ""
}

func setString(k reflect.Kind, beforeTag string) string {
	typeS := fmt.Sprintf("%s%s", beforeTag, k.String())
	return typeS
}
