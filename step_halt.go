package main
import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
  "github.com/askholme/vultr"
)

type stepHalt struct{}

func (s *stepHalt) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*vultr.Client)
	c := state.Get("config").(config)
	ui := state.Get("ui").(packer.Ui)
	serverId := state.Get("server_id").(string)

	serverInfo, err := client.GetServer(serverId)
	if err != nil {
		err := fmt.Errorf("Error checking server state: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
  if serverInfo.Status != "active" {
		err := fmt.Errorf("Error server is not active, status is: ", serverInfo.Status)
		state.Put("error", err)
		ui.Error(err.Error())
    return multistep.ActionHalt
  }
	if serverInfo.Power != "running"  {
		// Already off, don't do anything
		return multistep.ActionContinue
	}
  
	// Pull the plug on the Droplet
	ui.Say("Forcefully shutting down Server...")
	err = client.HaltServer(serverId)
	if err != nil {
		err := fmt.Errorf("Error powering off server: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	log.Println("Waiting for halt to complete...")
	_,err = waitForServerState("active","stopped", serverId, client, c.stateTimeout)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	return multistep.ActionContinue
}

func (s *stepHalt) Cleanup(state multistep.StateBag) {
	// no cleanup
}