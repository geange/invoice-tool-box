package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

	title := []string{"旧文件名", "新文件名", "发票代号", "发票号码", "日期", "货物名", "金额", "税额", "价税合计", "销售方名称"}
	err = writer.Write(title)
	if err != nil {
		return err
	}

	for _, invoice := range p.result.Invoices {
		v := invoice.Result
		commodityName := make(map[string]struct{})
		for _, item := range v.WordsResult.CommodityName {
			text := removeType(item.Word)
			commodityName[text] = struct{}{}
		}

		values := []string{
			fmt.Sprintf(`"%s"`, invoice.Name),
			fmt.Sprintf(`"%s"`, invoice.NewName),
			fmt.Sprintf(`"%s"`, v.WordsResult.InvoiceCode),
			fmt.Sprintf(`"%s"`, v.WordsResult.InvoiceNum),
			fmt.Sprintf(`"%s"`, v.WordsResult.InvoiceDate),
			join(commodityName),
			fmt.Sprintf(`"%s"`, v.WordsResult.TotalAmount),
			fmt.Sprintf(`"%s"`, v.WordsResult.TotalTax),
			fmt.Sprintf(`"%s"`, v.WordsResult.AmountInFiguers),
			fmt.Sprintf(`"%s"`, v.WordsResult.SellerName),
		}
		err := writer.Write(values)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}

func removeType(name string) string {
	values := []rune(name)
	for i := len(values) - 1; i >= 0; i-- {
		if values[i] == '*' {
			return string(values[i+1:])
		}
	}
	return name
}

func join(values map[string]struct{}) string {
	names := make([]string, 0)
	for v := range values {
		names = append(names, v)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}
