package common

import (
	"context"
)

// A driver is able to talk to HyperV and perform certain
// operations with it. Some of the operations on here may seem overly
// specific, but they were built specifically in mind to handle features
// of the HyperV builder for Packer, and to abstract differences in
// versions out of the builder steps, so sometimes the methods are
// extremely specific.
type Driver interface {

	// Checks if the VM named is running.
	IsRunning(context.Context, string) (bool, error)

	// Checks if the VM named is off.
	IsOff(context.Context, string) (bool, error)

	//How long has VM been on
	Uptime(ctx context.Context, vmName string) (uint64, error)

	// Start starts a VM specified by the name given.
	Start(context.Context, string) error

	// Stop stops a VM specified by the name given.
	Stop(context.Context, string) error

	// Verify checks to make sure that this driver should function
	// properly. If there is any indication the driver can't function,
	// context.Context,  this will return an error.
	Verify(context.Context) error

	// Finds the MAC address of the NIC nic0
	Mac(context.Context, string) (string, error)

	// Finds the IP address of a VM connected that uses DHCP by its MAC address
	IpAddress(context.Context, string) (string, error)

	// Finds the hostname for the ip address
	GetHostName(context.Context, string) (string, error)

	// Finds the IP address of a host adapter connected to switch
	GetHostAdapterIpAddressForSwitch(context.Context, string) (string, error)

	// Type scan codes to virtual keyboard of vm
	TypeScanCodes(context.Context, string, string) error

	//Get the ip address for network adaptor
	GetVirtualMachineNetworkAdapterAddress(context.Context, string) (string, error)

	//Set the vlan to use for switch
	SetNetworkAdapterVlanId(context.Context, string, string) error

	//Set the vlan to use for machine
	SetVirtualMachineVlanId(context.Context, string, string) error

	SetVmNetworkAdapterMacAddress(context.Context, string, string) error

	//Replace the network adapter with a (non-)legacy adapter
	ReplaceVirtualMachineNetworkAdapter(context.Context, string, bool) error

	UntagVirtualMachineNetworkAdapterVlan(context.Context, string, string) error

	CreateExternalVirtualSwitch(context.Context, string, string) error

	GetVirtualMachineSwitchName(context.Context, string) (string, error)

	ConnectVirtualMachineNetworkAdapterToSwitch(context.Context, string, string) error

	CreateVirtualSwitch(context.Context, string, string) (bool, error)

	DeleteVirtualSwitch(context.Context, string) error

	CreateVirtualMachine(context.Context, string, string, string, int64, int64, int64, string, uint, bool, bool, string) error

	AddVirtualMachineHardDrive(context.Context, string, string, string, int64, int64, string) error

	CloneVirtualMachine(context.Context, string, string, string, bool, string, string, string, int64, string, bool) error

	DeleteVirtualMachine(context.Context, string) error

	GetVirtualMachineGeneration(context.Context, string) (uint, error)

	SetVirtualMachineCpuCount(context.Context, string, uint) error

	SetVirtualMachineMacSpoofing(context.Context, string, bool) error

	SetVirtualMachineDynamicMemory(context.Context, string, bool) error

	SetVirtualMachineSecureBoot(context.Context, string, bool, string) error

	SetVirtualMachineVirtualizationExtensions(context.Context, string, bool) error

	EnableVirtualMachineIntegrationService(context.Context, string, string) error

	ExportVirtualMachine(context.Context, string, string) error

	PreserveLegacyExportBehaviour(context.Context, string, string) error

	MoveCreatedVHDsToOutputDir(context.Context, string, string) error

	CompactDisks(context.Context, string) (string, error)

	RestartVirtualMachine(context.Context, string) error

	CreateDvdDrive(context.Context, string, string, uint) (uint, uint, error)

	MountDvdDrive(context.Context, string, string, uint, uint) error

	SetBootDvdDrive(context.Context, string, uint, uint, uint) error

	UnmountDvdDrive(context.Context, string, uint, uint) error

	DeleteDvdDrive(context.Context, string, uint, uint) error

	MountFloppyDrive(context.Context, string, string) error

	UnmountFloppyDrive(context.Context, string) error

	// Connect connects to a VM specified by the name given.
	Connect(context.Context, string) (context.CancelFunc, error)

	// Disconnect disconnects to a VM specified by the context cancel function.
	Disconnect(context.CancelFunc)
}
