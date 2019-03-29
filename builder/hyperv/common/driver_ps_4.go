package common

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/hashicorp/packer/common/powershell"
	"github.com/hashicorp/packer/common/powershell/hyperv"
)

type HypervPS4Driver struct {
}

func NewHypervPS4Driver(ctx context.Context) (Driver, error) {
	appliesTo := "Applies to Windows 8.1, Windows PowerShell 4.0, Windows Server 2012 R2 only"

	// Check this is Windows
	if runtime.GOOS != "windows" {
		err := fmt.Errorf("%s", appliesTo)
		return nil, err
	}

	ps4Driver := &HypervPS4Driver{}

	if err := ps4Driver.Verify(ctx); err != nil {
		return nil, err
	}

	return ps4Driver, nil
}

func (d *HypervPS4Driver) IsRunning(ctx context.Context, vmName string) (bool, error) {
	return hyperv.IsRunning(ctx, vmName)
}

func (d *HypervPS4Driver) IsOff(ctx context.Context, vmName string) (bool, error) {
	return hyperv.IsOff(ctx, vmName)
}

func (d *HypervPS4Driver) Uptime(ctx context.Context, vmName string) (uint64, error) {
	return hyperv.Uptime(ctx, vmName)
}

// Start starts a VM specified by the name given.
func (d *HypervPS4Driver) Start(ctx context.Context, vmName string) error {
	return hyperv.StartVirtualMachine(ctx, vmName)
}

// Stop stops a VM specified by the name given.
func (d *HypervPS4Driver) Stop(ctx context.Context, vmName string) error {
	return hyperv.StopVirtualMachine(ctx, vmName)
}

func (d *HypervPS4Driver) Verify(ctx context.Context) error {

	if err := d.verifyPSVersion(ctx); err != nil {
		return err
	}

	if err := d.verifyPSHypervModule(ctx); err != nil {
		return err
	}

	if err := d.verifyHypervPermissions(ctx); err != nil {
		return err
	}

	return nil
}

// Get mac address for VM.
func (d *HypervPS4Driver) Mac(ctx context.Context, vmName string) (string, error) {
	res, err := hyperv.Mac(ctx, vmName)

	if err != nil {
		return res, err
	}

	if res == "" {
		err := fmt.Errorf("%s", "No mac address.")
		return res, err
	}

	return res, err
}

// Get ip address for mac address.
func (d *HypervPS4Driver) IpAddress(ctx context.Context, mac string) (string, error) {
	res, err := hyperv.IpAddress(ctx, mac)

	if err != nil {
		return res, err
	}

	if res == "" {
		err := fmt.Errorf("%s", "No ip address.")
		return res, err
	}
	return res, err
}

// Get host name from ip address
func (d *HypervPS4Driver) GetHostName(ctx context.Context, ip string) (string, error) {
	return powershell.GetHostName(ctx, ip)
}

func (d *HypervPS4Driver) GetVirtualMachineGeneration(ctx context.Context, vmName string) (uint, error) {
	return hyperv.GetVirtualMachineGeneration(ctx, vmName)
}

// Finds the IP address of a host adapter connected to switch
func (d *HypervPS4Driver) GetHostAdapterIpAddressForSwitch(ctx context.Context, switchName string) (string, error) {
	res, err := hyperv.GetHostAdapterIpAddressForSwitch(ctx, switchName)

	if err != nil {
		return res, err
	}

	if res == "" {
		err := fmt.Errorf("%s", "No ip address.")
		return res, err
	}
	return res, err
}

// Type scan codes to virtual keyboard of vm
func (d *HypervPS4Driver) TypeScanCodes(ctx context.Context, vmName string, scanCodes string) error {
	return hyperv.TypeScanCodes(ctx, vmName, scanCodes)
}

// Get network adapter address
func (d *HypervPS4Driver) GetVirtualMachineNetworkAdapterAddress(ctx context.Context, vmName string) (string, error) {
	return hyperv.GetVirtualMachineNetworkAdapterAddress(ctx, vmName)
}

//Set the vlan to use for switch
func (d *HypervPS4Driver) SetNetworkAdapterVlanId(ctx context.Context, switchName string, vlanId string) error {
	return hyperv.SetNetworkAdapterVlanId(ctx, switchName, vlanId)
}

//Set the vlan to use for machine
func (d *HypervPS4Driver) SetVirtualMachineVlanId(ctx context.Context, vmName string, vlanId string) error {
	return hyperv.SetVirtualMachineVlanId(ctx, vmName, vlanId)
}

func (d *HypervPS4Driver) SetVmNetworkAdapterMacAddress(ctx context.Context, vmName string, mac string) error {
	return hyperv.SetVmNetworkAdapterMacAddress(ctx, vmName, mac)
}

//Replace the network adapter with a (non-)legacy adapter
func (d *HypervPS4Driver) ReplaceVirtualMachineNetworkAdapter(ctx context.Context, vmName string, virtual bool) error {
	return hyperv.ReplaceVirtualMachineNetworkAdapter(ctx, vmName, virtual)
}

func (d *HypervPS4Driver) UntagVirtualMachineNetworkAdapterVlan(ctx context.Context, vmName string, switchName string) error {
	return hyperv.UntagVirtualMachineNetworkAdapterVlan(ctx, vmName, switchName)
}

func (d *HypervPS4Driver) CreateExternalVirtualSwitch(ctx context.Context, vmName string, switchName string) error {
	return hyperv.CreateExternalVirtualSwitch(ctx, vmName, switchName)
}

func (d *HypervPS4Driver) GetVirtualMachineSwitchName(ctx context.Context, vmName string) (string, error) {
	return hyperv.GetVirtualMachineSwitchName(ctx, vmName)
}

func (d *HypervPS4Driver) ConnectVirtualMachineNetworkAdapterToSwitch(ctx context.Context, vmName string, switchName string) error {
	return hyperv.ConnectVirtualMachineNetworkAdapterToSwitch(ctx, vmName, switchName)
}

func (d *HypervPS4Driver) DeleteVirtualSwitch(ctx context.Context, switchName string) error {
	return hyperv.DeleteVirtualSwitch(ctx, switchName)
}

func (d *HypervPS4Driver) CreateVirtualSwitch(ctx context.Context, switchName string, switchType string) (bool, error) {
	return hyperv.CreateVirtualSwitch(ctx, switchName, switchType)
}

func (d *HypervPS4Driver) AddVirtualMachineHardDrive(ctx context.Context, vmName string, vhdFile string, vhdName string,
	vhdSizeBytes int64, diskBlockSize int64, controllerType string) error {
	return hyperv.AddVirtualMachineHardDiskDrive(ctx, vmName, vhdFile, vhdName, vhdSizeBytes,
		diskBlockSize, controllerType)
}

func (d *HypervPS4Driver) CreateVirtualMachine(ctx context.Context, vmName string, path string, harddrivePath string, ram int64,
	diskSize int64, diskBlockSize int64, switchName string, generation uint, diffDisks bool,
	fixedVHD bool, version string) error {
	return hyperv.CreateVirtualMachine(ctx, vmName, path, harddrivePath, ram, diskSize, diskBlockSize, switchName,
		generation, diffDisks, fixedVHD, version)
}

func (d *HypervPS4Driver) CloneVirtualMachine(ctx context.Context, cloneFromVmcxPath string, cloneFromVmName string,
	cloneFromSnapshotName string, cloneAllSnapshots bool, vmName string, path string, harddrivePath string,
	ram int64, switchName string, copyTF bool) error {
	return hyperv.CloneVirtualMachine(ctx, cloneFromVmcxPath, cloneFromVmName, cloneFromSnapshotName,
		cloneAllSnapshots, vmName, path, harddrivePath, ram, switchName, copyTF)
}

func (d *HypervPS4Driver) DeleteVirtualMachine(ctx context.Context, vmName string) error {
	return hyperv.DeleteVirtualMachine(ctx, vmName)
}

func (d *HypervPS4Driver) SetVirtualMachineCpuCount(ctx context.Context, vmName string, cpu uint) error {
	return hyperv.SetVirtualMachineCpuCount(ctx, vmName, cpu)
}

func (d *HypervPS4Driver) SetVirtualMachineMacSpoofing(ctx context.Context, vmName string, enable bool) error {
	return hyperv.SetVirtualMachineMacSpoofing(ctx, vmName, enable)
}

func (d *HypervPS4Driver) SetVirtualMachineDynamicMemory(ctx context.Context, vmName string, enable bool) error {
	return hyperv.SetVirtualMachineDynamicMemory(ctx, vmName, enable)
}

func (d *HypervPS4Driver) SetVirtualMachineSecureBoot(ctx context.Context, vmName string, enable bool, templateName string) error {
	return hyperv.SetVirtualMachineSecureBoot(ctx, vmName, enable, templateName)
}

func (d *HypervPS4Driver) SetVirtualMachineVirtualizationExtensions(ctx context.Context, vmName string, enable bool) error {
	return hyperv.SetVirtualMachineVirtualizationExtensions(ctx, vmName, enable)
}

func (d *HypervPS4Driver) EnableVirtualMachineIntegrationService(ctx context.Context, vmName string,
	integrationServiceName string) error {
	return hyperv.EnableVirtualMachineIntegrationService(ctx, vmName, integrationServiceName)
}

func (d *HypervPS4Driver) ExportVirtualMachine(ctx context.Context, vmName string, path string) error {
	return hyperv.ExportVirtualMachine(ctx, vmName, path)
}

func (d *HypervPS4Driver) PreserveLegacyExportBehaviour(ctx context.Context, srcPath string, dstPath string) error {
	return hyperv.PreserveLegacyExportBehaviour(ctx, srcPath, dstPath)
}

func (d *HypervPS4Driver) MoveCreatedVHDsToOutputDir(ctx context.Context, srcPath string, dstPath string) error {
	return hyperv.MoveCreatedVHDsToOutputDir(ctx, srcPath, dstPath)
}

func (d *HypervPS4Driver) CompactDisks(ctx context.Context, path string) (result string, err error) {
	return hyperv.CompactDisks(ctx, path)
}

func (d *HypervPS4Driver) RestartVirtualMachine(ctx context.Context, vmName string) error {
	return hyperv.RestartVirtualMachine(ctx, vmName)
}

func (d *HypervPS4Driver) CreateDvdDrive(ctx context.Context, vmName string, isoPath string, generation uint) (uint, uint, error) {
	return hyperv.CreateDvdDrive(ctx, vmName, isoPath, generation)
}

func (d *HypervPS4Driver) MountDvdDrive(ctx context.Context, vmName string, path string, controllerNumber uint,
	controllerLocation uint) error {
	return hyperv.MountDvdDrive(ctx, vmName, path, controllerNumber, controllerLocation)
}

func (d *HypervPS4Driver) SetBootDvdDrive(ctx context.Context, vmName string, controllerNumber uint, controllerLocation uint,
	generation uint) error {
	return hyperv.SetBootDvdDrive(ctx, vmName, controllerNumber, controllerLocation, generation)
}

func (d *HypervPS4Driver) UnmountDvdDrive(ctx context.Context, vmName string, controllerNumber uint, controllerLocation uint) error {
	return hyperv.UnmountDvdDrive(ctx, vmName, controllerNumber, controllerLocation)
}

func (d *HypervPS4Driver) DeleteDvdDrive(ctx context.Context, vmName string, controllerNumber uint, controllerLocation uint) error {
	return hyperv.DeleteDvdDrive(ctx, vmName, controllerNumber, controllerLocation)
}

func (d *HypervPS4Driver) MountFloppyDrive(ctx context.Context, vmName string, path string) error {
	return hyperv.MountFloppyDrive(ctx, vmName, path)
}

func (d *HypervPS4Driver) UnmountFloppyDrive(ctx context.Context, vmName string) error {
	return hyperv.UnmountFloppyDrive(ctx, vmName)
}

func (d *HypervPS4Driver) verifyPSVersion(ctx context.Context) error {

	log.Printf("Enter method: %s", "verifyPSVersion")
	// check PS is available and is of proper version
	versionCmd := "$host.version.Major"

	var ps powershell.PowerShellCmd
	cmdOut, err := ps.Output(ctx, versionCmd)
	if err != nil {
		return err
	}

	versionOutput := strings.TrimSpace(cmdOut)
	log.Printf("%s output: %s", versionCmd, versionOutput)

	ver, err := strconv.ParseInt(versionOutput, 10, 32)

	if err != nil {
		return err
	}

	if ver < 4 {
		err := fmt.Errorf("%s", "Windows PowerShell version 4.0 or higher is expected")
		return err
	}

	return nil
}

func (d *HypervPS4Driver) verifyPSHypervModule(ctx context.Context) error {

	log.Printf("Enter method: %s", "verifyPSHypervModule")

	versionCmd := "function foo(){try{ $commands = Get-Command -Module Hyper-V;if($commands.Length -eq 0){return $false} }catch{return $false}; return $true} foo"

	var ps powershell.PowerShellCmd
	cmdOut, err := ps.Output(ctx, versionCmd)
	if err != nil {
		return err
	}

	if powershell.IsFalse(cmdOut) {
		err := fmt.Errorf("%s", "PS Hyper-V module is not loaded. Make sure Hyper-V feature is on.")
		return err
	}

	return nil
}

func (d *HypervPS4Driver) isCurrentUserAHyperVAdministrator(ctx context.Context) (bool, error) {
	//SID:S-1-5-32-578 = 'BUILTIN\Hyper-V Administrators'
	//https://support.microsoft.com/en-us/help/243330/well-known-security-identifiers-in-windows-operating-systems

	var script = `
$identity = [System.Security.Principal.WindowsIdentity]::GetCurrent()
$principal = new-object System.Security.Principal.WindowsPrincipal($identity)
$hypervrole = [System.Security.Principal.SecurityIdentifier]"S-1-5-32-578"
return $principal.IsInRole($hypervrole)
`

	var ps powershell.PowerShellCmd
	cmdOut, err := ps.Output(ctx, script)
	if err != nil {
		return false, err
	}

	return powershell.IsTrue(cmdOut), nil
}

func (d *HypervPS4Driver) verifyHypervPermissions(ctx context.Context) error {

	log.Printf("Enter method: %s", "verifyHypervPermissions")

	hyperVAdmin, err := d.isCurrentUserAHyperVAdministrator(ctx)
	if err != nil {
		log.Printf("Error discovering if current is is a Hyper-V Admin: %s", err)
	}
	if !hyperVAdmin {

		isAdmin, _ := powershell.IsCurrentUserAnAdministrator(ctx)

		if !isAdmin {
			err := fmt.Errorf("%s", "Current user is not a member of 'Hyper-V Administrators' or 'Administrators' group")
			return err
		}
	}

	return nil
}

// Connect connects to a VM specified by the name given.
func (d *HypervPS4Driver) Connect(ctx context.Context, vmName string) (context.CancelFunc, error) {
	return hyperv.ConnectVirtualMachine(ctx, vmName)
}

// Disconnect disconnects to a VM specified by calling the context cancel function returned
// from Connect.
func (d *HypervPS4Driver) Disconnect(cancel context.CancelFunc) {
	hyperv.DisconnectVirtualMachine(cancel)
}
