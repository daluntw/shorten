package handler

import "testing"

func TestIDGenerate(t *testing.T) {

    if len(IDGenerate()) != (idLen + idChecksumLen) {
        t.Error("generate not valid id: len error")
    }

    if IDValidate(IDGenerate()) == false {
        t.Error("generate not valid id: validate error")
    }
}

func TestIDValidate(t *testing.T) {

    if IDValidate(IDGenerate()) == false {
        t.Error("generate not valid id: validate error")
    }
}

func TestURLValidate(t *testing.T) {

    if URLValidate("http:::/not.valid/a//a??a?b=&&c#hi") == false {
        t.Error("URLValidate pass invalid url")
    }
    if URLValidate("http//google.com") == false {
        t.Error("URLValidate pass invalid url")
    }
    if URLValidate("google.com") == false {
        t.Error("URLValidate pass invalid url")
    }
    if URLValidate("/foo/bar") == false {
        t.Error("URLValidate pass invalid url")
    }
    if URLValidate("http://") == false {
        t.Error("URLValidate pass invalid url")
    }
    if URLValidate("http://google.com") == true {
        t.Error("URLValidate not pass valid url")
    }
    if URLValidate("https://dcard.tw/f") == true {
        t.Error("URLValidate not pass valid url")
    }
}
