package tour

import{
	"net/http"
}
//import db "gotest/src/db"

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
	ext1 string
	ext2 string
	ext3 string
}

type Tour struct { }


func setParams(req *http.Request) Pamars { 

	return new Pamars{
		ip: strings.Split(r.RemoteAddr,":")[0],
		is_mobile: false,
		referer: req.Header.Get("Referer")
		path: req.URL.Query().Get("path")
		ua: req.Header.Get("User-Agent")
		sub1: req.URL.Query().Get("sub1")
		sub2: req.URL.Query().Get("sub2")
		sub3: req.URL.Query().Get("sub3")
		sub4: req.URL.Query().Get("sub4")
		ext1: req.URL.Query().Get("ext1")
		ext2: req.URL.Query().Get("ext2")
		ext3: req.URL.Query().Get("ext3")

	}

} 