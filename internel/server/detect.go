package server

import (
	"log"
	"time"
)

func (s *Server) detect() {
	for {
		time.Sleep(time.Second * 3)
		s.detectInternel()
	}
}

func (s *Server) detectInternel() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, v := range s.rNode {
		if v.Timeout == nil {
			continue
		}
		if !time.Now().Before(*v.Timeout) {
			s.rNode[i].TimeoutCount += 1
			s.rNode[i].TimeoutBl = true
			s.rNode[i].TimeoutSec = time.Now().Unix() - v.Timeout.Unix()

			if !v.SendEmail {
				s.rNode[i].SendEmail = true
				err := s.email.SendEmail(*v)
				if err != nil {
					log.Println(err)
				}
			}
		}

		if time.Now().Before(*v.Timeout) && v.TimeoutBl == true {
			s.rNode[i].SendEmail = false
			s.rNode[i].TimeoutBl = false
			s.rNode[i].TimeoutSec = 0
		}
	}
}
