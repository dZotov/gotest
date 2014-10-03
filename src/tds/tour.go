package tour

import{
	"net/http"
}
import db "gotest/src/db"

type Pamars struct{
	ip string
	is_mobile bool
	referer string
	path string
	l string
	referer_hash string
	ua string
	sub1 string
	sub2 string
	sub3 string
	sub4 string


}

type Tour struct { }


func (h Tour) SetParams(req *http.Request) Pamars { 

	return new Pamars{
		ip: strings.Split(r.RemoteAddr,":")[0],
		is_mobile: false,
		referer: req.Header.Get("Referer"),
		path: req.
		ua: req.Header.Get("User-Agent"),
		sub1: string
		sub2: string
		sub3: string
		sub4: string

	}

} 