package lgjx

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"github.com/nguyenthenguyen/docx"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func OpenLetter(order lgjx.Order, templateFile lgjx.File) (letter lgjx.Letter, file lgjx.File, encryptFile lgjx.File, err error) {
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	out, err := os.Create(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
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
	insuranceName := global.GVA_CONFIG.Insurance.Name
	insuranceCreditCode := global.GVA_CONFIG.Insurance.CreditCode
	insuranceAddress := global.GVA_CONFIG.Insurance.Address
	insuranceZipCode := global.GVA_CONFIG.Insurance.ZipCode
	insuranceTel := global.GVA_CONFIG.Insurance.Tel
	tenderDepositString := fmt.Sprintf("%.2f", *order.Apply.TenderDeposit)
	tenderDepositStringFixed, _ := strconv.ParseFloat(tenderDepositString, 64)
	tenderDepositCNYString := ConvertNumToCny(tenderDepositStringFixed)
	year, month, day := time.Now().Date()

	letterReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
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
	_ = letterDocx.WriteToFile(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".docx")

	letterEncryptReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
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
	_ = letterEncryptDocx.WriteToFile(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.docx")

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")

	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+".docx", "--outdir", global.GVA_CONFIG.Insurance.TempDir).Run()
	if err != nil {
		return
	}
	err = exec.Command("libreoffice", "--headless", "--convert-to", "pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encrypt.docx", "--outdir", global.GVA_CONFIG.Insurance.TempDir).Run()
	if err != nil {
		return
	}

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".docx")
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.docx")

	err = exec.Command("java", "-jar", global.GVA_CONFIG.Insurance.SignProgram, global.GVA_CONFIG.Insurance.KeyFile, global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+".pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"Signed.pdf", global.GVA_CONFIG.Insurance.StampFile).Run()
	if err != nil {
		return
	}
	err = exec.Command("java", "-jar", global.GVA_CONFIG.Insurance.SignProgram, global.GVA_CONFIG.Insurance.KeyFile, global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encrypt.pdf", global.GVA_CONFIG.Insurance.TempDir+"letter"+fileName+"encryptSigned.pdf", global.GVA_CONFIG.Insurance.StampFile).Run()
	if err != nil {
		return
	}

	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".pdf")
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + "encrypt.pdf")

	letterFileName := "letter" + fileName + "Signed.pdf"
	encryptLetterFileName := "letter" + fileName + "encryptSigned.pdf"
	letterFile, err := os.Open(global.GVA_CONFIG.Insurance.TempDir + letterFileName)
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
	encryptLetterFile, err := os.Open(global.GVA_CONFIG.Insurance.TempDir + encryptLetterFileName)
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
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + letterFileName)
	_ = os.Remove(global.GVA_CONFIG.Insurance.TempDir + encryptLetterFileName)
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
