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
	insuredName := *order.Project.TendereeName
	insureName := *order.Apply.InsureName
	projectNo := *order.Project.ProjectNo
	projectName := *order.Project.ProjectName
	tenderDeposit := *order.Project.TenderDeposit
	tenderDepositString := fmt.Sprintf("%.2f", tenderDeposit)
	tenderDepositStringFixed, _ := strconv.ParseFloat(tenderDepositString, 64)
	tenderDepositCNY := ConvertNumToCny(tenderDepositStringFixed)
	insuranceName := global.GVA_CONFIG.Insurance.Name
	insuranceAddress := global.GVA_CONFIG.Insurance.Address
	insuranceZipCode := global.GVA_CONFIG.Insurance.ZipCode
	insuranceTel := global.GVA_CONFIG.Insurance.Tel
	year, month, day := currentTime.Date()
	tenderEndDate, _ := time.ParseInLocation("2006-01-02 15:04:05", *order.Project.TenderEndDate, time.Local)
	insureEndDate := tenderEndDate.AddDate(0, 0, int(*order.Project.ProjectDay))
	insureEndDayeYear, insureEndDayeMonth, insureEndDayeDay := insureEndDate.Date()
	insureDay := *order.Project.ProjectDay

	insuranceCreditCode := global.GVA_CONFIG.Insurance.CreditCode
	urlString := GenerateMD5String([]byte(time.Now().String()))
	elogOutDate := currentTime.Format("2006-01-02 15:04:05")
	elogUrl := urlString[:16]
	elogEncryptUrl := urlString[16:]
	insureStartDateString := currentTime.Format("2006-01-02 15:04:05")
	insureEndDateString := insureEndDate.Format("2006-01-02 15:04:05")
	validateCode := urlString[12:20]

	letterReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
	letterDocx := letterReader.Editable()
	_ = letterDocx.Replace("{elogNo}", elogNo, -1)
	_ = letterDocx.Replace("{insuredName}", insuredName, -1)
	_ = letterDocx.Replace("{insureName}", insureName, -1)
	_ = letterDocx.Replace("{projectNo}", projectNo, -1)
	_ = letterDocx.Replace("{projectName}", projectName, -1)
	_ = letterDocx.Replace("{tenderDeposit}", tenderDepositString, -1)
	_ = letterDocx.Replace("{tenderDepositCNY}", tenderDepositCNY, -1)
	_ = letterDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterDocx.Replace("{insureStartDayeYear}", strconv.Itoa(year), -1)
	_ = letterDocx.Replace("{insureStartDayeMonth}", strconv.Itoa(int(month)), -1)
	_ = letterDocx.Replace("{insureStartDayeDay}", strconv.Itoa(day), -1)
	_ = letterDocx.Replace("{insureEndDayeYear}", strconv.Itoa(insureEndDayeYear), -1)
	_ = letterDocx.Replace("{insureEndDayeMonth}", strconv.Itoa(int(insureEndDayeMonth)), -1)
	_ = letterDocx.Replace("{insureEndDayeDay}", strconv.Itoa(insureEndDayeDay), -1)
	_ = letterDocx.Replace("{insureDay}", strconv.FormatInt(insureDay, 10), -1)
	_ = letterDocx.WriteToFile(global.GVA_CONFIG.Insurance.TempDir + "letter" + fileName + ".docx")

	letterEncryptReader, _ := docx.ReadDocxFile(global.GVA_CONFIG.Insurance.TempDir + fileName + ".docx")
	letterEncryptDocx := letterEncryptReader.Editable()
	_ = letterEncryptDocx.Replace("{elogNo}", elogNo, -1)
	_ = letterEncryptDocx.Replace("{insuredName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insureName}", "**********", -1)
	_ = letterEncryptDocx.Replace("{projectNo}", *order.Project.ProjectNo, -1)
	_ = letterEncryptDocx.Replace("{projectName}", *order.Project.ProjectName, -1)
	_ = letterEncryptDocx.Replace("{tenderDeposit}", "**********", -1)
	_ = letterEncryptDocx.Replace("{tenderDepositCNY}", "**********", -1)
	_ = letterEncryptDocx.Replace("{insuranceName}", insuranceName, -1)
	_ = letterEncryptDocx.Replace("{insuranceAddress}", insuranceAddress, -1)
	_ = letterEncryptDocx.Replace("{insuranceZipCode}", insuranceZipCode, -1)
	_ = letterEncryptDocx.Replace("{insuranceTel}", insuranceTel, -1)
	_ = letterEncryptDocx.Replace("{year}", strconv.Itoa(year), -1)
	_ = letterEncryptDocx.Replace("{month}", strconv.Itoa(int(month)), -1)
	_ = letterEncryptDocx.Replace("{day}", strconv.Itoa(day), -1)
	_ = letterEncryptDocx.Replace("{insureStartDayeYear}", strconv.Itoa(year), -1)
	_ = letterEncryptDocx.Replace("{insureStartDayeMonth}", strconv.Itoa(int(month)), -1)
	_ = letterEncryptDocx.Replace("{insureStartDayeDay}", strconv.Itoa(day), -1)
	_ = letterEncryptDocx.Replace("{insureEndDayeYear}", strconv.Itoa(insureEndDayeYear), -1)
	_ = letterEncryptDocx.Replace("{insureEndDayeMonth}", strconv.Itoa(int(insureEndDayeMonth)), -1)
	_ = letterEncryptDocx.Replace("{insureEndDayeDay}", strconv.Itoa(insureEndDayeDay), -1)
	_ = letterEncryptDocx.Replace("{insureDay}", strconv.FormatInt(insureDay, 10), -1)
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
		TenderDeposit:       order.Project.TenderDeposit,
		InsureStartDate:     &insureStartDateString,
		InsureEndDate:       &insureEndDateString,
		InsureDay:           order.Project.ProjectDay,
		ValidateCode:        &validateCode,
	}
	return
}
