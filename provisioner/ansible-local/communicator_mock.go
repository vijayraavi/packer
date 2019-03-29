package ansiblelocal

import (
	"context"
	"io"
	"os"

	"github.com/hashicorp/packer/packer"
)

type communicatorMock struct {
	startCommand      []string
	uploadDestination []string
}

func (c *communicatorMock) Start(ctx context.Context, cmd *packer.RemoteCmd) error {
	c.startCommand = append(c.startCommand, cmd.Command)
	cmd.SetExited(0)
	return nil
}

func (c *communicatorMock) Upload(ctx context.Context, dst string, _ io.Reader, _ *os.FileInfo) error {
	c.uploadDestination = append(c.uploadDestination, dst)
	return nil
}

func (c *communicatorMock) UploadDir(ctx context.Context, dst, src string, exclude []string) error {
	return nil
}

func (c *communicatorMock) Download(ctx context.Context, src string, dst io.Writer) error {
	return nil
}

func (c *communicatorMock) DownloadDir(ctx context.Context, src, dst string, exclude []string) error {
	return nil
}

func (c *communicatorMock) verify() {
}
