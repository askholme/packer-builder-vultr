package main
import (
	"fmt"
	"log"
	"time"
  "github.com/askholme/vultr"
)

type ret struct {
  err error
  info *vult.Server
}


// waitForState simply blocks until the droplet is in
// a state we expect, while eventually timing out.
func waitForServerState(state string,power string, serverId string, client *vultr.Client, timeout time.Duration) (*vultr.Server,error) {
	done := make(chan struct{})
	defer close(done)
	result := make(chan ret{}, 1)
  go func() {
      attempts := 0
  		for {
  			attempts += 1
  			log.Printf("Checking server status... (attempt: %d)", attempts)
  			serverInfo, err := client.GetServer(serverId)
  			if err != nil {
  				result <- ret{err,nil}
  				return
  			}
  			if serverInfo.Status == state && (serverinfo.Power == power || power == "") {
  				result <- ret{nil,&serverInfo}
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
func waitForSnapshotState(state string, snapshotId string, client *vultr.Client, timeout time.Duration) (*vultr.Server,error) {
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
		return retval.info,retval.err
	case <-time.After(timeout):
		err := fmt.Errorf("Timeout while waiting to for snapshot")
		return nil,err
	}
}