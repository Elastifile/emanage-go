package emanage

import (
	"fmt"

	"github.com/go-errors/errors"

	"github.com/elastifile/emanage-go/src/remote"
	"github.com/elastifile/emanage-go/src/rest"
)

const (
	packagesUri = "/api/sw_packages"
)

type swPackages struct {
	conn *rest.Session
}

type SwPackageInfo struct {
	Version          string `json:"version"`
	ReleaseTimestamp string `json:"release_timestamp"`
	Filename         string `json:"filename"`
	Compatibility    string `json:"compatibility"`
	VerifyStatus     string `json:"verify_status"`
}
type SwPackageVerifyStatus struct {
	Status string `json:"status"`
}

type SwPackageVerify string

const (
	SwPackageVerifySuccess string = "success"
)

func (packs *swPackages) Info() (*SwPackageInfo, error) { //TODO(eduard): change to Sw_packages
	var pack SwPackageInfo
	uri := fmt.Sprintf("%s/info", packagesUri)
	err := packs.conn.Request(rest.MethodGet, uri, nil, &pack)
	return &pack, err
}
func (packs *swPackages) Verify() (*SwPackageVerifyStatus, error) {
	var pack SwPackageVerifyStatus
	uri := fmt.Sprintf("%s/verify", packagesUri)
	err := packs.conn.Request(rest.MethodPost, uri, nil, &pack)
	return &pack, err
}

func (packs *swPackages) Delete() error {
	uri := fmt.Sprintf("%s/delete", packagesUri)
	err := packs.conn.Request(rest.MethodDelete, uri, nil, nil)
	return err
}

func (packs *swPackages) Upload(uri, ip string) error {
	upgradeDir := "/elastifile/emanage/public/sw_package" //TODO(eduard): temp w/a , may be later to do upload by stream
	pkgName := "pkg000.tar"
	r, err := remote.NewRemote(string(ip))
	if err != nil {
		return err
	}
	command := fmt.Sprintf("mkdir %v;rm -rf %v/*", upgradeDir, upgradeDir)
	_, err = r.Run(command)
	if err != nil {
		return err
	}
	command = fmt.Sprintf("curl -Sso %v/%v %v; chown -R apache:apache %v", upgradeDir, pkgName, uri, upgradeDir)
	_, err = r.Run(command)
	return err
}

func (packs *swPackages) UploadByCli(uri, ip string) (string, error) {
	upgradeDir := "~/sw_package" //TODO(eduard): temp w/a , may be later to do upload by stream
	pkgName := "pkg000.tar"
	r, err := remote.NewRemote(string(ip))
	if err != nil {
		return "", errors.WrapPrefix(err, "Failed to create new remote", 0)
	}

	// Cleanup the upgrade path
	command := fmt.Sprintf("mkdir -p %v && rm -rf %v/*", upgradeDir, upgradeDir)
	output, err := r.Run(command)
	if err != nil {
		return string(output), err
	}

	// Check the upgrade path is a dir
	command = fmt.Sprintf("test -d %v", upgradeDir)
	output, err = r.Run(command)
	if err != nil {
		return string(output), errors.WrapPrefix(err, fmt.Sprintf("directory %v not exist", upgradeDir), 0)
	}

	// Fetch the upgrade package
	command = fmt.Sprintf("curl -Sso %v/%v %v && chown -R apache:apache %v", upgradeDir, pkgName, uri, upgradeDir)
	_, err = r.Run(command)
	if err != nil {
		return "", err
	}

	// Trigger the upgrade
	command = fmt.Sprintf("source ~root/elfs_admin && elfs-cli sw_package upload --content %v/%v", upgradeDir, pkgName)
	output, err = r.Run(command)
	if err != nil {
		return string(output), errors.WrapPrefix(err, fmt.Sprintf("sw_package upload failed. Command: %v Output: %v", command, output), 0)
	}

	// Cleanup the upgrade path
	command = fmt.Sprintf("rm -rf %v/*", upgradeDir)
	_, err1 := r.Run(command)
	if err1 != nil {
		logger.Warn(fmt.Sprintf("Failed to remove dir %v contents. Command: %v Error: %v", upgradeDir, command, err1))
	}
	return string(output), err
}
