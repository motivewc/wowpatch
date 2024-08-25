//go:build darwin

package platform

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RemoveCodesigningSignature(path string) error {
	cmd := exec.Command("/usr/bin/codesign", "--remove-signature", path)
	var out bytes.Buffer
	cmd.Stdout = nil
	cmd.Stderr = &out

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("unable to remove codesigning attributes: %s\n%w", out.String(), err)
	}

	return nil
}
