package pkg

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"
)

func (p *Processor) DefaultCSV(fileName string) error {
	err := p.Load()
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(p.dir, fileName))
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)

	title := []string{"发票代号", "发票号码", "日期", "货物名", "金额", "税额", "价税合计", "销售方名称"}
	err = writer.Write(title)
	if err != nil {
		return err
	}

	for _, invoice := range p.result.Invoices {
		v := invoice.Result
		commodityName := make([]string, 0)
		for _, item := range v.WordsResult.CommodityName {
			commodityName = append(commodityName, item.Word)
		}

		values := []string{
			v.WordsResult.InvoiceCode,
			v.WordsResult.InvoiceNum,
			v.WordsResult.InvoiceDate,
			strings.Join(commodityName, ","),
			v.WordsResult.TotalAmount,
			v.WordsResult.TotalTax,
			v.WordsResult.AmountInFiguers,
			v.WordsResult.SellerName,
		}
		err := writer.Write(values)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
