package cmd

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

//Prober Interface for probes
type Prober interface {
	SetEndpoint(string)
	Run(chan bool)
}

//NewProbeMachine Configures a ProbeMachine
func NewProbeMachine(
	timeout int,
	probe Prober,
) *ProbeMachine {
	return &ProbeMachine{
		Timeout: timeout,
		Probe:   probe,
	}
}

//ProbeMachine It probes
type ProbeMachine struct {
	//Max time a probing can take
	Timeout int

	//The probe does the heavy lifting
	Probe Prober
}

// Run Starts the healthcheck process
func (pm *ProbeMachine) Run() bool {
	resultChannel := make(chan bool, 1)

	go func() {
		for true {
			pm.Probe.Run(resultChannel)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	select {
	case <-time.After(time.Second * time.Duration(pm.Timeout)):
		break
	case result := <-resultChannel:
		return result
	}

	return false
}

//HTTPProbe Basic HTTP probe
type HTTPProbe struct {
	Endpoint string
}

//SetEndpoint Sets endpoint according to Prober interface
func (probe *HTTPProbe) SetEndpoint(endpoint string) {
	probe.Endpoint = endpoint
}

//Run Run the probe
func (probe *HTTPProbe) Run() chan bool {
	resultChannel := make(chan bool, 1)

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(probe.Endpoint)

	if err != nil {
		log.Errorf("Request failed to %s with message: %s", probe.Endpoint, err)
		resultChannel <- false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		resultChannel <- true
	} else {
		resultChannel <- false
	}

	return resultChannel
}
