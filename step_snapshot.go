package main

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"log"
  "github.com/askholme/vultr"
)

type stepSnapshot struct{}

func (s *stepSnapshot) Run(state multistep.StateBag) multistep.StepAction {
	client := state.Get("client").(*vultr.Client)
	ui := state.Get("ui").(packer.Ui)
	c := state.Get("config").(config)
	serverId := state.Get("droplet_id").(string)

	ui.Say(fmt.Sprintf("Creating snapshot: %v", c.SnapshotName))
	snapshotId,err := client.CreateSnapshot(serverId, c.SnapshotName)
	if err != nil {
		err := fmt.Errorf("Error creating snapshot: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say("Waiting for snapshot to complete...")
	err = waitForSnapshotState("complete", snapshotId, client, c.stateTimeout)
	if err != nil {
		err := fmt.Errorf("Error waiting for snapshot to complete: %s", err)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	log.Printf("Snapshot image ID: %d", snapshotId)

	state.Put("snapshot_id", snapshotId)
	state.Put("snapshot_name", c.SnapshotName)
	state.Put("region", c.Region)
	return multistep.ActionContinue
}

func (s *stepSnapshot) Cleanup(state multistep.StateBag) {
	// no cleanup
}