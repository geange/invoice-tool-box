package sdk

import (
	"encoding/base64"
	"encoding/json"
	"strings"
)

type VatInvoice struct {
}

func (s *SDK) VatInvoice(accessToken string, imageBytes []byte) (*VatInvoiceResponse, error) {
	host := "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"

	image := base64.StdEncoding.EncodeToString(imageBytes)

	resp, err := s.client.R().SetQueryParam("access_token", accessToken).
		SetFormData(map[string]string{
			"image": image,
		}).Post(host)
	if err != nil {
		return nil, err
	}

	var response VatInvoiceResponse
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type VatInvoiceResponse struct {
	LogId          int `json:"log_id"`
	WordsResultNum int `json:"words_result_num"`
	WordsResult    struct {
		InvoiceNum        string `json:"InvoiceNum"`
		InvoiceNumConfirm string `json:"InvoiceNumConfirm"`
		SellerName        string `json:"SellerName"`
		CommodityTaxRate  []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityTaxRate"`
		SellerBank      string `json:"SellerBank"`
		Checker         string `json:"Checker"`
		TotalAmount     string `json:"TotalAmount"`
		CommodityAmount []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityAmount"`
		InvoiceDate  string `json:"InvoiceDate"`
		CommodityTax []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityTax"`
		PurchaserName string `json:"PurchaserName"`
		CommodityNum  []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityNum"`
		Province           string `json:"Province"`
		City               string `json:"City"`
		SheetNum           string `json:"SheetNum"`
		Agent              string `json:"Agent"`
		PurchaserBank      string `json:"PurchaserBank"`
		Remarks            string `json:"Remarks"`
		Password           string `json:"Password"`
		SellerAddress      string `json:"SellerAddress"`
		PurchaserAddress   string `json:"PurchaserAddress"`
		InvoiceCode        string `json:"InvoiceCode"`
		InvoiceCodeConfirm string `json:"InvoiceCodeConfirm"`
		CommodityUnit      []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityUnit"`
		Payee                string `json:"Payee"`
		PurchaserRegisterNum string `json:"PurchaserRegisterNum"`
		CommodityPrice       []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityPrice"`
		NoteDrawer        string `json:"NoteDrawer"`
		AmountInWords     string `json:"AmountInWords"`
		AmountInFiguers   string `json:"AmountInFiguers"`
		TotalTax          string `json:"TotalTax"`
		InvoiceType       string `json:"InvoiceType"`
		SellerRegisterNum string `json:"SellerRegisterNum"`
		CommodityName     []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityName"`
		CommodityType []struct {
			Word string `json:"word"`
			Row  string `json:"row"`
		} `json:"CommodityType"`
	} `json:"words_result"`
}

func (v *VatInvoiceResponse) CSVTitle() []string {
	commodityName := make([]string, 0)
	for _, item := range v.WordsResult.CommodityName {
		commodityName = append(commodityName, item.Word)
	}

	return []string{
		"发票代号",
		"发票号码",
		"日期",
		"货物名",
		"金额",
		"税额",
		"价税合计",
		"销售方名称",
	}
}

func (v *VatInvoiceResponse) CSV() []string {
	commodityName := make([]string, 0)
	for _, item := range v.WordsResult.CommodityName {
		commodityName = append(commodityName, item.Word)
	}

	return []string{
		v.WordsResult.InvoiceCode,
		v.WordsResult.InvoiceNum,
		v.WordsResult.InvoiceDate,
		strings.Join(commodityName, ","),
		v.WordsResult.TotalAmount,
		v.WordsResult.TotalTax,
		v.WordsResult.AmountInFiguers,
		v.WordsResult.SellerName,
	}
}
