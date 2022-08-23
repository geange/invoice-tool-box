package pkg

import (
	"encoding/json"
	"github.com/geange/invoice-toolbox/sdk"
	"os"
)

type Result struct {
	AbsPath      string              `json:"abs_path"`
	InvoiceField string              `json:"invoice_field"` // 用于重命名的字段
	Invoices     []InvoiceInfo       `json:"invoices"`      //
	ErrInvoices  []string            `json:"err_invoices"`  // 处理失败的图片
	Duplicate    map[string][]string `json:"duplicate"`     // 重复的文件
}

func NewResult(invoiceField string) *Result {
	return &Result{
		InvoiceField: invoiceField,
		Invoices:     make([]InvoiceInfo, 0),
		ErrInvoices:  make([]string, 0),
		Duplicate:    make(map[string][]string),
	}
}

type InvoiceInfo struct {
	Name    string                  `json:"name"`     // 文件名
	NewName string                  `json:"new_name"` // 复制的文件名
	Modify  int64                   `json:"modify"`   // 修改时间
	Result  *sdk.VatInvoiceResponse `json:"result"`   // 识别结果
}

func (c *Result) Name() string {
	return "config.json"
}

func (c *Result) Init() error {
	bs, err := os.ReadFile(c.Name())
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, c)
}

func (c *Result) Save() error {
	file, err := os.Create(c.Name())
	if err != nil {
		return err
	}
	defer file.Close()

	bs, _ := json.Marshal(c)

	_, err = file.Write(bs)
	return err
}
