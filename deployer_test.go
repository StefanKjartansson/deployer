package deployer

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net"
    "net/http"
    "net/http/httptest"
    "sync"
    "testing"
    ws "code.google.com/p/go.net/websocket"
)

var serverAddr string
var once sync.Once

func newConfig(t *testing.T, path string) *ws.Config {
        config, _ := ws.NewConfig(fmt.Sprintf("ws://%s%s", serverAddr, path), "http://localhost")
        return config
}

func startServer() {
        configServer()
        server := httptest.NewServer(nil)
        serverAddr = server.Listener.Addr().String()
        log.Print("Test Server running on ", serverAddr)
}

func makeTestData() {
    for _, s := range []string{"foo","bar","boz"} {
        _ = NewProject(s)
    }
}

func makeRequest(method, url string, v interface{}) (error) {

    client := &http.Client{}

    req, err := http.NewRequest(method, url, nil)

    if err != nil {
        return err
    }

    res, err := client.Do(req)
    if err != nil {
        return err
    }

    if res.StatusCode != 200 {
        log.Printf("Wrong status code: %d\n", res.StatusCode)
    }

    content, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    log.Println(string(content))
    if err != nil {
        return err
    }

    err = json.Unmarshal(content, v)
    return err
}

func TestIndexHandler(t *testing.T) {
    once.Do(startServer)

    r, err := http.Get(fmt.Sprintf("http://%s/", serverAddr))

    if err != nil {
        t.Errorf("Index: %v\n", err)
    }

    if r.StatusCode != 200 {
        t.Errorf("Wrong status code: %d\n", r.StatusCode)
    }

}

func TestDetailHandler(t *testing.T) {
    once.Do(startServer)

    foo := NewProject("foobar")

    returned := Project{}

    url := fmt.Sprintf("http://%s/projects/%s/", serverAddr, foo.ID)

    log.Println(url)

    err := makeRequest("GET", url, &returned)

    if err != nil {
        t.Errorf("Index: %v", err)
    }

    if foo.ID != returned.ID {
        t.Errorf("Got: %v, Expected: %s\n", returned, foo)
    }
}

func TestListHandler(t *testing.T) {
    once.Do(startServer)
    makeTestData()

    url := fmt.Sprintf("http://%s/projects/", serverAddr)

    returned := []Project{}

    log.Println(url)

    err := makeRequest("GET", url, &returned)

    if err != nil {
        t.Errorf("Index: %v", err)
    }

}

func TestDeployHandler(t *testing.T) {
    once.Do(startServer)

    go h.run()

    real_test := NewProject("real_test")
    url := fmt.Sprintf("http://%s/projects/%s/deploy", serverAddr, real_test.ID)

    client, err := net.Dial("tcp", serverAddr)
    if err != nil {
        t.Fatal("dialing", err)
    }
    conn, err := ws.NewClient(newConfig(t, "/ws"), client)
    if err != nil {
        t.Errorf("WebSocket handshake error: %v", err)
        return
    }

    log.Println(url)
    returned := Job{}
    err = makeRequest("POST", url, &returned)

    if err != nil {
        t.Errorf("Index: %v", err)
    }

    for i:=0; i<3; i++ {

        var actual_msg = make([]byte, 512)
        n, err := conn.Read(actual_msg)
        if err != nil {
            t.Errorf("Read: %v", err)
        }
        actual_msg = actual_msg[0:n]
        log.Println(string(actual_msg))

    }

}
