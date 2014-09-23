package main
import (
	"fmt"
	"log"
  "github.com/askholme/vultr"
)

type Artifact struct {
	// The name of the snapshot
	snapshotName string
	// The ID of the image
	snapshotId string
	// The name of the region
	regionName string

	// The client for making API calls
	client *vultr.Client
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Files() []string {
	// No files with Vultr
	return nil
}

func (a *Artifact) Id() string {
	// mimicing the aws builder
	return fmt.Sprintf("%s:%s", a.regionName, a.snapshotName)
}

func (a *Artifact) String() string {
	return fmt.Sprintf("A snapshot was created: '%v' in region '%v'", a.snapshotName, a.regionName)
}

func (a *Artifact) Destroy() error {
	log.Printf("Destroying image: %d (%s)", a.snapshotId, a.snapshotName)
	return a.client.DeleteSnapshot(a.snapshotId)
}