package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/tidwall/gjson"
	"os"
	"path/filepath"
	"time"

	"github.com/geange/invoice-toolbox/sdk"
)

func (p *Processor) Run() error {
	p.bar = pb.Start64(int64(len(p.waitList)))
	for _, v := range p.waitList {
		p.bar.Increment()
		err := p.deal(&v)
		if err != nil {
			//log.Println(err.Error())
			p.pushFailFile(v.Name)
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func (p *Processor) Close() error {
	defer p.bar.Finish()

	invoices := make([]InvoiceInfo, 0)
	for _, invoice := range p.invoices {
		invoices = append(invoices, invoice)
	}
	p.result.Invoices = invoices

	for k, items := range p.newInvoices {
		if len(items) > 1 {
			p.result.Duplicate[k] = items
		}
	}

	bs, err := json.Marshal(p.result)
	if err != nil {
		return err
	}

	result, err := os.Create(p.resultPath())
	if err != nil {
		return err
	}

	_, err = result.Write(bs)
	return err
}

func (p *Processor) deal(v *File) error {
	picContent, err := os.ReadFile(p.srcPath(v.Name))
	if err != nil {
		return err
	}

	if !isImage(picContent, v.Name) {
		return err
	}

	invoice, err := p.sdk.VatInvoice(p.token, picContent)
	if err != nil {
		return err
	}

	newName := p.newName(v, invoice)

	info := InvoiceInfo{
		Name:    v.Name,
		NewName: newName,
		Modify:  v.ModTime,
		Result:  invoice,
	}

	// 重命名文件
	if p.needRename {

		// 如果存在了就不处理
		if _, ok := p.newInvoices[newName]; ok {
			p.pushDuplicate(v.Name, newName)
			return ErrExistDuplicateFile
		}
		p.pushDuplicate(v.Name, newName)

		newFilePath := p.rename(v, invoice)
		newFile, err := os.Create(newFilePath)
		if err != nil {
			return err
		}
		defer newFile.Close()

		_, err = newFile.Write(picContent)
		if err != nil {
			return err
		}
	}

	p.invoices[v.Name] = info
	return nil
}

func (p *Processor) pushFailFile(name string) {
	if exist(name, p.result.ErrInvoices) {
		return
	}
	p.result.ErrInvoices = append(p.result.ErrInvoices, name)
}

func (p *Processor) pushDuplicate(name, newName string) {
	if exist(name, p.newInvoices[newName]) {
		return
	}
	p.newInvoices[newName] = append(p.newInvoices[newName], name)
}

func exist(name string, vars []string) bool {
	for _, s := range vars {
		if s == name {
			return true
		}
	}
	return false
}

func (p *Processor) newName(file *File, invoice *sdk.VatInvoiceResponse) string {
	bytes, _ := json.Marshal(invoice)
	return gjson.GetBytes(bytes,
		fmt.Sprintf("words_result.%s", p.invoiceField)).String() +
		file.Ext
}

func (p *Processor) rename(file *File, invoice *sdk.VatInvoiceResponse) string {
	return filepath.Join(p.dir, CopyFileDir, p.newName(file, invoice))
}
