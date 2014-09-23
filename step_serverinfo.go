package main
import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
  "time"
)

type stepServerInfo struct{}

func (s *stepServerInfo) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*DigitalOceanClient)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("config").(config)
	serverId := state.Get("server_id").(string)

	ui.Say("Waiting for server to be powered on...")
  type waitOpts struct {
    state       string
    power       string
    serverId    string
    client      *vultr.Client
    serverIp    string
    serverPort  int
  }
	serverInfo, err := waitForServerState("active","running",serverId,%client, c.stateTimeout)
	if err != nil {
		err := fmt.Errorf("Error waiting for server to be running: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	state.Put("server_ip", serverInfo.Ip)
  state.Put("default_password", serverInfo.Password)
	return multistep.ActionContinue
}

func (s *stepServerInfo) Cleanup(state multistep.StateBag) {
	// no cleanup
}