package handler

import (
    "github.com/cespare/xxhash"
    "github.com/valyala/fastrand"
    "math/rand"
    "net/url"
    "strings"
    "time"
)

var validChar string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var idLen, idChecksumLen int = 5, 1

func init() {
    rand.Seed(time.Now().Unix())
}

func IDGenerate() string {

    // create {{idLen}} char + {{idChecksumLen}} checksum
    var sb strings.Builder

    for i := 0; i < idLen; i++ {
        sb.WriteRune(rune(validChar[fastrand.Uint32n(uint32(len(validChar)))]))
    }

    // use xxHash64 result as checksum
    // TODO: there is more efficiency way to do this

    for i := 0; i < idChecksumLen; i++ {
        r := validChar[xxhash.Sum64String(sb.String())%uint64(len(validChar))]
        sb.WriteRune(rune(r))
    }

    return sb.String()
}

func IDValidate(id string) bool {

    if len(id) != (idLen + idChecksumLen) {
        return false
    }

    var sb strings.Builder
    sb.WriteString(id[:idLen])

    for i := 0; i < idChecksumLen; i++ {
        if r := validChar[xxhash.Sum64String(sb.String())%uint64(len(validChar))]; r != id[idLen+i] {
            return false
        } else {
            sb.WriteRune(rune(r))
        }
    }

    return true
}

func URLValidate(str string) bool {
    //source: https://stackoverflow.com/questions/31480710/validate-url-with-standard-package-in-go
    u, err := url.Parse(str)
    return err == nil && u.Scheme != "" && u.Host != ""
}
