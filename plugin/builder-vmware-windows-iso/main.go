package main

import (
	"github.com/packer-community/packer-windows-plugins/builder/vmware-windows/iso"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(iso.Builder))
	server.Serve()
}
