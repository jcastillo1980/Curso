package sendmail

import (
	"log"
	"net/http"
	"net/smtp"
	"os"
)

// SendMailObj ????
type SendMailObj struct {
	Host     string
	Port     string
	From     string
	NameFrom string
	Pass     string
	To       []string
}

func finalizaHTTP(w http.ResponseWriter, retorno int, respuesta string) {
	w.WriteHeader(retorno)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte(respuesta))
}

// ServeHTTP interface para resolver servidor
func (s SendMailObj) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		s.To = []string{}
		if vk := r.Form["TO1"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if vk := r.Form["TO2"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if vk := r.Form["TO3"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if vk := r.Form["TO4"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if vk := r.Form["TO5"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if vk := r.Form["TO6"]; vk != nil {
			if len(vk[0]) > 0 {
				s.To = append(s.To, vk[0])
			}
		}
		if len(s.To) == 0 {
			finalizaHTTP(w, 401, "0FAIL_NO_TO")
			return
		}

		body := ""
		subj := ""
		if vk := r.Form["BODY"]; vk != nil {
			if len(vk[0]) > 0 {
				body = body + vk[0]
			}
		}
		if vk := r.Form["TITULO"]; vk != nil {
			if len(vk[0]) > 0 {
				subj = subj + vk[0]
			}
		}
		if s.SendTextMail(subj, body) == true {
			finalizaHTTP(w, 200, "1OK")
		} else {
			finalizaHTTP(w, 200, "0FAIL_SEND")
		}

	} else {
		finalizaHTTP(w, 400, "0FAIL")
	}
}

// GetParamSendMail Recoge los parametros de las variables del sistema
func (s *SendMailObj) GetParamSendMail() {
	s.Host = os.Getenv("LAUNCHER_HOST")
	s.Port = os.Getenv("LAUNCHER_PORT")
	s.From = os.Getenv("LAUNCHER_FROM")
	s.NameFrom = os.Getenv("LAUNCHER_NAMEFROM")
	s.Pass = os.Getenv("LAUNCHER_PASS")

	valor := os.Getenv("LAUNCHER_TO1")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
	valor = os.Getenv("LAUNCHER_TO2")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
	valor = os.Getenv("LAUNCHER_TO3")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
	valor = os.Getenv("LAUNCHER_TO4")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
	valor = os.Getenv("LAUNCHER_TO5")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
	valor = os.Getenv("LAUNCHER_TO6")
	if len(valor) > 0 {
		s.To = append(s.To, valor)
	}
}

// SendTextMail envia un texto por email
func (s *SendMailObj) SendTextMail(subject string, body string) bool {
	if len(s.Host) <= 0 {
		return false
	}
	if len(s.Port) <= 0 {
		return false
	}
	if len(s.From) <= 0 {
		return false
	}
	if len(s.Pass) <= 0 {
		return false
	}
	if len(s.To) <= 0 {
		return false
	}
	if len(body) <= 0 {
		return false
	}

	msg := "From: " + s.NameFrom + " <" + s.From + ">\n" + "To: "
	for i, vv := range s.To {
		if i > 0 {
			msg = msg + ";"
		}
		msg = msg + "<" + vv + ">"
	}

	msg = msg + "\n" + "Subject: " + subject + "\n\n" + body

	err := smtp.SendMail(s.Host+":"+s.Port,
		smtp.PlainAuth("", s.From, s.Pass, s.Host),
		s.From, s.To, []byte(msg))

	if err != nil {
		log.Printf("SMTP Error: %s\r\n", err)
		return false
	}

	return true
}

// PrintObj Imprime los datos de la estructura
func (s SendMailObj) PrintObj() {
	log.Printf("%#v\r\n", s)
}
