package lgjx

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"github.com/nguyenthenguyen/docx"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func OpenLetter(order lgjx.Order, templateFile lgjx.File) (letter lgjx.Letter, file lgjx.File, encryptFile lgjx.File, err error) {
	basePath := "./tmp/"
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	out, err := os.Create(basePath + fileName + ".docx")
	if err != nil {
		return
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	_, err = out.Write(templateFile.FileSteam)
	if err != nil {
		return
	}

	currentTime := time.Now()
	elogNo := GenerateMD5String([]byte(*order.OrderNo + time.Now().String()))
	urlString := GenerateMD5String([]byte(time.Now().String()))
	elogUrl := urlString[:16]
	elogEncryptUrl := urlString[16:]
	elogOutDate := currentTime.Format("2006-01-02 15:04:05")
	insureStartDate := currentTime.Format("2006-01-02 15:04:05")
	insureEndDate := currentTime.AddDate(0, 0, 15).Format("2006-01-02 15:04:05")
	var insureDay int64 = 15
	validateCode := urlString[12:20]
	insuranceName := "XXXX担保公司"
	insuranceCreditCode := "CNXKeCQaEDydwjK5tC"
	insuranceAddress := "XXX担保公司地址"
	insuranceZipCode := "330000"
	insuranceTel := "13813813800"
	tenderDepositString := fmt.Sprintf("%.2f", *order.Apply.TenderDeposit)
	tenderDepositStringFixed, _ := strconv.ParseFloat(tenderDepositString, 64)
	tenderDepositCNYString := ConvertNumToCny(tenderDepositStringFixed)
	year, month, day := time.Now().Date()

	letterReader, _ := docx.ReadDocxFile(basePath + fileName + ".docx")
	letterDocx := letterReader.Editable()
	_ = letterDocx.Replace("{elogNo}", elogNo, -1)
	_ = letterDocx.Replace("{insuredName}", *order.Apply.InsuredName, -1)
	_ = letterDocx.Replace("{insureName}", *order.Apply.InsureName, -1)
	_ = letterDocx.Replace("{projectNo}", *order.Apply.ProjectNo, -1)
	_ = letterDocx.Replace("{projectName}", *order.Apply.ProjectName, -1)
	_ = letterDocx.Replace("{tenderDeposit}", tenderDepositString, -1)
	_ = letterDocx.Replace("{tenderDepositCNY}", tenderDepositCNYString, -1)
	_ = letterDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterDocx.WriteToFile(basePath + "letter" + fileName + ".docx")

	letterEncryptReader, _ := docx.ReadDocxFile(basePath + fileName + ".docx")
	letterEncryptDocx := letterEncryptReader.Editable()
	_ = letterEncryptDocx.Replace("{elogNo}", elogNo, -1)
	_ = letterEncryptDocx.Replace("{insuredName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insureName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{projectNo}", *order.Apply.ProjectNo, -1)
	_ = letterEncryptDocx.Replace("{projectName}", *order.Apply.ProjectName, -1)
	_ = letterEncryptDocx.Replace("{tenderDeposit}", "**********", -1)
	_ = letterEncryptDocx.Replace("{tenderDepositCNY}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterEncryptDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterEncryptDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterEncryptDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterEncryptDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterEncryptDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterEncryptDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterEncryptDocx.WriteToFile(basePath + "letter" + fileName + "encrypt.docx")

	_ = os.Remove(basePath + fileName + ".docx")

	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", basePath+"letter"+fileName+".docx", "--outdir", basePath).Run()
	if err != nil {
		return
	}
	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", basePath+"letter"+fileName+"encrypt.docx", "--outdir", basePath).Run()
	if err != nil {
		return
	}

	_ = os.Remove(basePath + "letter" + fileName + ".docx")
	_ = os.Remove(basePath + "letter" + fileName + "encrypt.docx")

	letterFileName := "letter" + fileName + ".pdf"
	encryptLetterFileName := "letter" + fileName + "encrypt.pdf"
	letterFile, err := os.Open(basePath + letterFileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(letterFile)
	letterFileContent, err := io.ReadAll(letterFile)
	if err != nil {
		return
	}
	encryptLetterFile, err := os.Open(basePath + encryptLetterFileName)
	if err != nil {
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(encryptLetterFile)
	encryptLetterFileContent, err := io.ReadAll(encryptLetterFile)
	if err != nil {
		return
	}

	file = lgjx.File{
		FileName:  &letterFileName,
		FileSteam: letterFileContent,
	}
	encryptFile = lgjx.File{
		FileName:  &encryptLetterFileName,
		FileSteam: encryptLetterFileContent,
	}
	_ = os.Remove(basePath + letterFileName)
	_ = os.Remove(basePath + encryptLetterFileName)
	letter = lgjx.Letter{
		OrderID:             &order.ID,
		OrderNo:             order.OrderNo,
		ElogNo:              &elogNo,
		InsuranceName:       &insuranceName,
		InsuranceCreditCode: &insuranceCreditCode,
		ElogOutDate:         &elogOutDate,
		ElogUrl:             &elogUrl,
		ElogEncryptUrl:      &elogEncryptUrl,
		TenderDeposit:       order.Apply.TenderDeposit,
		InsureStartDate:     &insureStartDate,
		InsureEndDate:       &insureEndDate,
		InsureDay:           &insureDay,
		ValidateCode:        &validateCode,
	}
	return
}
