package models

import "encoding/xml"

type ServiceResponse struct {
	XMLName               xml.Name              `xml:"serviceResponse" json:"-"`
	AuthenticationSuccess AuthenticationSuccess `xml:"authenticationSuccess"`
}

type AuthenticationSuccess struct {
	XMLName    xml.Name   `xml:"authenticationSuccess" json:"-"`
	User       string     `xml:"user" json:"user"`
	Attributes Attributes `xml:"attributes" json:"attributes"`
}

type Attributes struct {
	XMLName    xml.Name `xml:"attributes" json:"-"`
	Ldap_cn    string   `xml:"ldap_cn" xml:"ldap_cn"`
	Kd_org     string   `xml:"kd_org" json:"kd_org"`
	Peran_user string   `xml:"peran_user" json:"peran_user"`
	Nama       string   `xml:"nama" json:"nama"`
	Npm        string   `xml:"npm" json:"npm"`
	Jurusan    Jurusan  `json:"jurusan"`
}

type Jurusan struct {
	Faculty      string `json:"faculty"`
	ShortFaculty string `json:"shortFaculty"`
	Major        string `json:"major"`
	Program      string `json:"program"`
}
