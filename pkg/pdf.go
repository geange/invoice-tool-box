package pkg

import (
	"errors"
	"fmt"
	"os/exec"
)

func convertPDFToJPEG(inputPDF, outputJPEG string) error {
	_, err := exec.LookPath("gs")
	if err != nil {
		return errors.New("gs not found")
	}

	// 使用 Ghostscript 将 PDF 转换为 JPEG
	cmd := exec.Command("gs", "-dNOPAUSE", "-dBATCH", "-sDEVICE=jpeg", "-r300", "-sOutputFile="+outputJPEG, inputPDF)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run Ghostscript command: %v", err)
	}
	return nil
}
