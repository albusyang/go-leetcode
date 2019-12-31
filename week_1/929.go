package main

import (
	"fmt"
)

func main() {
	emails := []string{"fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com", "fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com", "r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com", "r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com", "r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com", "fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com", "fg.r.u.uzj+fek@kziczvh.com"}
	fmt.Println(numUniqueEmails(emails))
}

func numUniqueEmails(emails []string) int {
    emailMap := make(map[string]int)
	for _, v := range emails {
		meetAt := false
		meetPlus := false
		emailAddr := ""
		for _, ev := range v {
			if meetAt {
				emailAddr += string(ev)
			} else {
				if ev == 64 {
					emailAddr += string(ev)
					meetAt = true
					continue
				}
				if meetPlus {
					continue
				} else {
					if ev == 43 {
						meetPlus = true
						continue
					}
					if ev == 46 {
						continue
					}
					emailAddr += string(ev)
				}
			}
		}

        _, ok := emailMap[emailAddr]
        if ok {
            continue
        } else {
            emailMap[emailAddr] = 1
        }
	}
	return len(emailMap)
}
