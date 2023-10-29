package resp

import (
	"api/utils"
	"encoding/xml"
)

type M map[string]any

func Code(code int) *M {
	return &M{"code": code}
}

func (m *M) Code(code int) *M {
	(*m)["code"] = code
	return m
}

func Msg(msg string) *M {
	return &M{"msg": msg}
}

func (m *M) Msg(msg string) *M {
	(*m)["msg"] = msg
	return m
}

func Data(data interface{}) *M {
	return &M{"data": data}
}

func (m *M) Data(data interface{}) *M {
	(*m)["data"] = data
	return m
}

func CodeMsg(code int, msg string) *M {
	return Code(code).Msg(msg)
}

func (m *M) CodeMsg(code int, msg string) *M {
	return m.Code(code).Msg(msg)
}

func (m *M) GetCode() int {
	return utils.CInt((*m)["code"])
}

func (m *M) GetMsg() string {
	return utils.CString((*m)["msg"])
}

func (m *M) GetData() interface{} {
	return (*m)["data"]
}

func (m *M) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "map",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range *m {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
