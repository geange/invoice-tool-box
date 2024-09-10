package pkg

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// Load 加载
func (p *Processor) Load() error {
	err := p.prepare()
	if err != nil {
		return err
	}
	err = p.loadResult()
	if err != nil {
		return err
	}
	err = p.loadFiles()
	if err != nil {
		return err
	}
	return p.getToken()
}

// 预处理
func (p *Processor) prepare() error {
	// 检查结果文件
	_, err := os.Stat(p.resultPath())
	if err == nil {
		p.existOldResult = true
	}

	// 检查重命名目录
	_, err = os.Stat(p.copyPath())
	if err == nil {
		return nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		err := os.Mkdir(p.copyPath(), 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

// 加载历史数据
func (p *Processor) loadResult() error {
	if !p.existOldResult {
		return nil
	}

	bs, err := os.ReadFile(p.resultPath())
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, &p.result); err != nil {
		return err
	}

	if p.invoiceField == "" {
		p.invoiceField = p.result.InvoiceField
	}

	p.result.ErrInvoices = []string{}
	p.newInvoices = map[string][]string{}

	if p.result.InvoiceField != p.invoiceField {
		p.result = NewResult(p.invoiceField)
		err := os.RemoveAll(p.copyPath())
		if err != nil {
			return err
		}
		err = os.Mkdir(p.copyPath(), 0755)
		if err != nil {
			return err
		}
	}

	for _, invoice := range p.result.Invoices {
		p.invoices[invoice.Name] = invoice
		p.pushDuplicate(invoice.Name, invoice.NewName)
	}

	return nil
}

// 加载当前目录的全部文件
func (p *Processor) loadFiles() error {
	entries, err := os.ReadDir(p.dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			continue
		}

		if !isImageByExt(entry.Name()) {
			continue
		}

		p.files = append(p.files, fileInfo)
	}

	for _, file := range p.files {
		if _, ok := p.invoices[file.Name()]; !ok {
			p.waitList = append(p.waitList, File{
				Name:    file.Name(),
				Ext:     filepath.Ext(file.Name()),
				Size:    file.Size(),
				ModTime: file.ModTime().Unix(),
			})
		}
	}
	return nil
}

func (p *Processor) dealPDF() error {
	entries, err := os.ReadDir(p.dir)
	if err != nil {
		return err
	}

	allFiles := make(map[string]bool)

	for _, entry := range entries {
		fileInfo, err := entry.Info()
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			continue
		}

		allFiles[entry.Name()] = true
	}

	for filename := range allFiles {

		if isPDF(filename) {
			jpegPath := p.srcPath(filename) + ".jpeg"
			if !allFiles[jpegPath] {
				err := convertPDFToJPEG(p.srcPath(filename), jpegPath)
				if err != nil {
					continue
				}
			}
		}
	}

	return nil
}

func (p *Processor) getToken() error {
	token, err := p.sdk.Token()
	if err != nil {
		return err
	}
	p.token = token.AccessToken
	return nil
}

func (p *Processor) srcPath(name string) string {
	return filepath.Join(p.dir, name)
}

func (p *Processor) copyPath() string {
	return filepath.Join(p.dir, CopyFileDir)
}

func (p *Processor) resultPath() string {
	return filepath.Join(p.dir, ResultFile)
}

type File struct {
	Name    string
	Ext     string
	Size    int64
	ModTime int64
}
