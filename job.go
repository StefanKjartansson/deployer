package deployer

import (
    "bytes"
    "fmt"
    "os/exec"
    "strings"
    "time"
)


type Job struct {
    ID string
    ProjectID string
    Status string
    Started time.Time
    Ended time.Time
}

func (j *Job) RunCommand () (err error){
    cmd := exec.Command("tr", "a-z", "A-Z")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err = cmd.Run()
    if err == nil {
        h.broadcast <- out.String()
    }
    return
}

func (j *Job) Start () () {
    message := fmt.Sprintf("Starting job %s at %v.", j.ID, j.Started)
    h.broadcast <- message
    err := j.RunCommand()
    if err != nil {
        h.broadcast <- fmt.Sprintf("%v", err)
    }
    (*j).Ended = time.Now()
    (*j).Status = "ended"
    message = fmt.Sprintf("Ending job %s at %v.", j.ID, j.Ended)
    h.broadcast <- message
}
