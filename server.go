package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Server struct {
	config *Config
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rpath := r.URL.Path
	if rpath == "/tasks" {
		s.showTasks(w)
		return
	}
	task := s.getTask(rpath)
	if task == nil {
		w.WriteHeader(404)
		return
	}
	r.ParseForm()
	if r.Form.Get("secret") != task.Secret {
		log.Println("User provided bad secret", r.Form.Get("secret"))
		w.WriteHeader(401)
		return
	}
	cmdstrs := task.render(map[string]string{})
	for _, cmdstr := range cmdstrs {
		cmd := exec.Command("/bin/bash", "-c", cmdstr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func (s *Server) showTasks(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(s.config.toCleanJSON())
}

func (s *Server) getTask(path string) *Task {
	for _, t := range s.config.Tasks {
		if t.Name == path[1:] {
			return &t
		}
	}
	return nil
}
