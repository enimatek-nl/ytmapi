package geckodriver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"ytmapi/mapper"
)

type GeckoDriver struct {
	port int
	host string
	path string
	cmd  *exec.Cmd
	url  string
	log  bool
}

func NewGeckoDriver(path string, verbose bool) (d *GeckoDriver) {
	d = &GeckoDriver{}
	d.path = path
	d.port = 9515
	d.host = "127.0.0.1"
	d.log = verbose
	return
}

func (d *GeckoDriver) Start() error {
	d.url = fmt.Sprintf("http://%s:%d/", d.host, d.port)

	d.cmd = exec.Command(d.path, "--port", strconv.Itoa(d.port), "--host", d.host)

	stdout, err := d.cmd.StdoutPipe()
	if err != nil {
		return errors.New(err.Error())
	}
	stderr, err := d.cmd.StderrPipe()
	if err != nil {
		return errors.New(err.Error())
	}
	if err := d.cmd.Start(); err != nil {
		return errors.New(err.Error())
	}

	timeout := 5 * time.Second
	now := time.Now()
	for {
		status, err := d.Get("/status")
		if err == nil {
			if status.M("value").B("ready") {
				log.Println("GeckoDriver is ready.")
				break
			}
		}
		if time.Since(now) > timeout {
			return errors.New("start failed: timeout expired")
		}
	}

	if d.log {
		go io.Copy(os.Stdout, stdout)
		go io.Copy(os.Stderr, stderr)
	}

	return nil
}

func (d *GeckoDriver) Stop() error {
	if d.cmd == nil {
		return errors.New("stop failed: geckodriver not running")
	}
	defer func() {
		d.cmd = nil
	}()
	d.cmd.Process.Signal(os.Interrupt)
	return nil
}

func (d *GeckoDriver) Get(uri string) (m mapper.Map, err error) {
	m = mapper.Map{}
	err = nil
	if d.cmd == nil {
		err = errors.New("get failed: geckodriver not running")
		return
	}
	resp, err := http.Get(fmt.Sprintf("http://%v:%v%v", d.host, d.port, uri))
	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			err = json.Unmarshal(body, &m)
		}
	} else {
		err = errors.New("not found")
	}
	return
}

func (d *GeckoDriver) Post(uri string, body []byte) (m mapper.Map, err error) {
	m = mapper.Map{}
	err = nil
	if d.cmd == nil {
		err = errors.New("post failed: geckodriver not running")
		return
	}
	resp, err := http.Post(fmt.Sprintf("http://%v:%v%v", d.host, d.port, uri), "application/json", bytes.NewBuffer(body))
	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(body), "\"value\":null") {
			return nil, nil
		}
		if err == nil {
			err = json.Unmarshal(body, &m)
		}
	} else {
		err = errors.New("not found")
	}
	return
}

func (d *GeckoDriver) Delete(uri string) (err error) {
	err = nil
	if d.cmd == nil {
		err = errors.New("delete failed: geckodriver not running")
		return
	}
	client := &http.Client{Timeout: time.Second * 2}
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://%v:%v%v", d.host, d.port, uri), nil)
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
	}
	return
}
