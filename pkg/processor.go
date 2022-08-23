package pkg

import (
	"github.com/cheggaaa/pb/v3"
	"github.com/geange/invoice-toolbox/sdk"
	"io/fs"
	"os"
)

type Processor struct {
	bar *pb.ProgressBar

	existOldResult       bool
	needRename           bool
	isChangeFileName     bool // 修改
	isChangeInvoiceField bool //
	invoiceField         string
	dir                  string

	result      *Result
	invoices    map[string]InvoiceInfo // 文件名 => 处理结果
	newInvoices map[string][]string    // 新名
	waitList    []File                 // 待处理的文件
	files       []fs.FileInfo          //

	sdk   *sdk.SDK
	token string
}

type ProcessorConfig struct {
	Dir          string
	InvoiceField string
}

var (
	baiduAK = os.Getenv("AK_BAIDU_INVOICE")
	baiduSK = os.Getenv("SK_BAIDU_INVOICE")
)

func NewProcessor(cfg *ProcessorConfig) *Processor {
	return &Processor{
		existOldResult:   false,
		needRename:       true,
		isChangeFileName: true,
		invoiceField:     cfg.InvoiceField,
		dir:              cfg.Dir,

		result:      NewResult(cfg.InvoiceField),
		invoices:    make(map[string]InvoiceInfo),
		newInvoices: make(map[string][]string),
		waitList:    make([]File, 0),
		files:       make([]fs.FileInfo, 0),

		sdk:   sdk.NewSDK(baiduAK, baiduSK),
		token: "",
	}
}
