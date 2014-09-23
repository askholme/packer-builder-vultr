package main
import (
	"fmt"
	"log"
	"time"
  "github.com/askholme/vultr"
)

type retStruct struct {
  err error
  info *vultr.Server
}

// waitForState simply blocks until the droplet is in
// a state we expect, while eventually timing out.
func waitForServerState(state string,power string, serverId string, client *vultr.Client, timeout time.Duration) (*vultr.Server,error) {
	done := make(chan struct{})
	defer close(done)
	result := make(chan retStruct, 1)
  go func() {
      attempts := 0
  		for {
  			attempts += 1
  			log.Printf("Checking server status... (attempt: %d)", attempts)
  			serverInfo, err := client.GetServer(serverId)
  			if err != nil {
  				result <- retStruct{err,nil}
  				return
  			}
  			if serverInfo.Status == state && (serverInfo.Power == power || power == "") {
  				result <- retStruct{nil,serverInfo}
          return
  			}

  			// Wait 3 seconds in between
  			time.Sleep(3 * time.Second)

  			// Verify we shouldn't exit
  			select {
  			case <-done:
  				// We finished, so just exit the goroutine
  				return
  			default:
  				// Keep going
  			}
  		}
    }()
	log.Printf("Waiting for up to %d seconds for server", timeout/time.Second)
	select {
	case retval := <-result:
		return retval.info,retval.err
	case <-time.After(timeout):
		err := fmt.Errorf("Timeout while waiting to for server")
		return nil,err
	}
}
func waitForSnapshotState(state string, snapshotId string, client *vultr.Client, timeout time.Duration) (error) {
	done := make(chan struct{})
	defer close(done)
	result := make(chan error, 1)
  go func() {
      attempts := 0
  		for {
  			attempts += 1
  			log.Printf("Checking snapshot status... (attempt: %d)", attempts)
  			snapshotInfo, err := client.GetSnapshot(snapshotId)
  			if err != nil {
  				result <- err
  				return
  			}
  			if snapshotInfo.Status == state  {
  				result <- nil
          return
  			}

  			// Wait 3 seconds in between
  			time.Sleep(3 * time.Second)

  			// Verify we shouldn't exit
  			select {
  			case <-done:
  				// We finished, so just exit the goroutine
  				return
  			default:
  				// Keep going
  			}
  		}
    }()
	log.Printf("Waiting for up to %d seconds for snapshot", timeout/time.Second)
	select {
	case retval := <-result:
		return retval
	case <-time.After(timeout):
		err := fmt.Errorf("Timeout while waiting to for snapshot")
		return err
	}
}