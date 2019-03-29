package main

import (
	"fmt"
	"strings"
)

func main() {
	emails := []string{"fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com", "fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com", "r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com", "r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com", "r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com", "fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com", "fg.r.u.uzj+fek@kziczvh.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
	emailList := make([]string, 0)
	for _, v := range emails {
		emailAddr := ""
		big := strings.Split(v, "@")
		
		middle := strings.Split(big[0], "+")
		
		small := strings.Split(middle[0], ".")
		for _, sv := range small {
			emailAddr += sv
		}
		emailAddr += "@"
		emailAddr += big[1]
		
		if len(emailList) == 0 {
			emailList = append(emailList, emailAddr)
		} else {
			exist := false
			for _, em := range emailList {
				if em == emailAddr {
					exist = true
				}
			}
			if !exist {
				emailList = append(emailList, emailAddr)
			}
		}
	}
	return len(emailList)
}
