package main
import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
  "time"
)

type stepShutdown struct {
}

func (s *stepShutdown) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
//	c := state.Get("config").(config)
  comm := state.Get("communicator").(packer.Communicator)
	ui.Say("Shutting down the server...")
	cmd := &packer.RemoteCmd{
				Command: fmt.Sprintf("shutdown -P"),
	}
	if err := comm.Start(cmd); err != nil {
  	cmd = &packer.RemoteCmd{
  				Command: fmt.Sprintf("sudo shutdown -P"),
    }
    if errs := comm.Start(cmd); errs != nil {
  		err_msg := fmt.Errorf("Error could not shutdown server safely: %s (%s)", err,errs)
  		state.Put("error", err_msg)
  		ui.Error(err_msg.Error())
  		return multistep.ActionHalt
    }
  }
  ui.Say("Waiting 20 seconds for server to properly shutdown...")
  time.Sleep(20*time.Second)
	return multistep.ActionContinue
}

func (s *stepShutdown) Cleanup(state multistep.StateBag) {
}