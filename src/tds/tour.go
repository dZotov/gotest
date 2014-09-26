package tour

import{
	"net/http"
}


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



} 