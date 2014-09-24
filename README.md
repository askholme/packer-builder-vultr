# Vultr builder for packer

See Bintray.md for binary downloads

a packer plugin enabling building of vultr snapshots. 
The code structure etc. borrows heavily from the packer digitalocean plugin. 

Communication with Vultr is done using the go API from github.com/askholme/vultr

Compile and place the binary in
* the same directory as packer, or
* `~/.packer.d/plugins`, or
* run packer from the directory with the binary in it

## Configuration

The builder supports the following parameters:
* `api_key` , the Vultr v1 API key
* `region`, id or name of a Vultr region, defaults to Atlanta (both enterend as a string)
* `plan`, id or name of a Vultr plan, defaults to the small 768MB ram plan (both entered as a string)
* `os`, id or name of a Vultr OS, defaults to Debian Wheezy (both entered as a string)
* `os_snapshot`, id or name of a Vultr Snapshot (string)
* `ipxe`, URL to boot from using Ipxe (string)
* `snapshot_name`, name of the snapshot that packer creates, defaults to `packer-{{timestamp}}` (string)
* `private_networking`, turn on private networking, default is false (bool)
* `IPv6`, turn on IPv6, default is false (bool)
* `ssh_username`, ssh username, default is `root` (string)
* `ssh_password`, ssh password, only for snapshots/custom OS (string)
* `ssh_key`, ssh private key for snapshots/custom OS (string)
* `ssh_port`, ssh port for snapshots/custom OS (int)

The user must provide `api_key`. Typically providing `region,plan` and one of `os,os_snapshot,ipxe` 
would be needed as well. If `os_snapshot` or `ipxe` is used ssh-connection information is also required

Example:
```
{
	"type"				: "vultr",
	"api_key_"			: "foo",
	"region"			: "Atlanta",
	"os"				: "Debian 7 x64 (wheezy)",
	"plan"				: "768 MB RAM,15 GB SSD,0.20 TB BW",
	"private_networking": true,
	"snapshot_name"		: "mysnapsnot-{{timestamp}}"	
}
```