package deployer

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

type Job struct {
	ID        string
	ProjectID string
	Status    string
	Started   time.Time
	Ended     time.Time
}

func (j *Job) RunCommand() (err error) {
	log.Println("Running job")
	cmdLines := strings.Split(Projects[j.ProjectID].Provisioner, "\n")
	for _, l := range cmdLines {
		foo := strings.Split(l, " ")
		cmd := exec.Command(foo[0], foo[1:]...)
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err == nil {
			h.broadcast <- out.String()
		}

	}
	return
}

func (j *Job) Start() {
	message := fmt.Sprintf("Starting job %s at %v.", j.ID, j.Started)
	log.Println(message)
	h.broadcast <- message
	err := j.RunCommand()
	if err != nil {
		h.broadcast <- fmt.Sprintf("%v", err)
	}
	(*j).Ended = time.Now()
	(*j).Status = "ended"
	message = fmt.Sprintf("Ending job %s at %v.", j.ID, j.Ended)

	log.Println(message)

	h.broadcast <- message
}
