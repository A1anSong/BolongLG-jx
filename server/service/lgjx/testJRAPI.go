package lgjx

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrclientrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrclientresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	lgjx2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lgjx"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"time"
)

type TestJRAPIService struct {
}

func (testJRAPIService *TestJRAPIService) ApplyOrder(reApply jrrequest.JRAPIApply) (resApply jrresponse.JRAPIApply, err error) {
	if reApply.OrderNo == nil ||
		reApply.ApplyNo == nil ||
		reApply.ProductNo == nil ||
		reApply.ProductType == nil ||
		reApply.ProductRate == nil ||
		reApply.ElogAmount == nil ||
		reApply.ProjectGuid == nil ||
		reApply.ProjectName == nil ||
		reApply.ProjectNo == nil ||
		reApply.TenderDeposit == nil ||
		reApply.DepositStartDate == nil ||
		reApply.DepositEndDate == nil ||
		reApply.OpenBeginDate == nil ||
		reApply.ElogTemplateNo == nil ||
		reApply.ElogTemplateName == nil ||
		reApply.InsuredName == nil ||
		reApply.InsuredCreditCode == nil ||
		reApply.InsureName == nil ||
		reApply.InsureCreditCode == nil ||
		reApply.InsureLegalName == nil ||
		reApply.InsureLegalIdCard == nil ||
		reApply.InsureAddress == nil ||
		reApply.ApplicantName == nil ||
		reApply.ApplicantIdCard == nil ||
		reApply.ApplicantTel == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		order := &lgjx.Order{
			OrderNo: reApply.OrderNo,
		}
		if err = tx.Create(&order).Error; err != nil {
			return err
		}
		var auditStatus int64 = 1
		auditOpinion := "todo:意见审批"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		attachInfo, _ := json.Marshal(reApply.AttachInfo)
		attachInfoString := string(attachInfo)
		productType, _ := strconv.ParseInt(*reApply.ProductType, 10, 64)
		apply := &lgjx.Apply{
			OrderID:           &order.ID,
			OrderNo:           reApply.OrderNo,
			ApplyNo:           reApply.ApplyNo,
			ProductNo:         reApply.ProductNo,
			ProductType:       &productType,
			ProductRate:       reApply.ProductRate,
			ElogAmount:        reApply.ElogAmount,
			ProjectGuid:       reApply.ProjectGuid,
			ProjectName:       reApply.ProjectName,
			ProjectNo:         reApply.ProjectNo,
			TenderDeposit:     reApply.TenderDeposit,
			DepositStartDate:  reApply.DepositStartDate,
			DepositEndDate:    reApply.DepositEndDate,
			OpenBeginDate:     reApply.OpenBeginDate,
			ElogTemplateNo:    reApply.ElogTemplateNo,
			ElogTemplateName:  reApply.ElogTemplateName,
			InsuredName:       reApply.InsuredName,
			InsuredCreditCode: reApply.InsuredCreditCode,
			InsuredAddress:    reApply.InsuredAddress,
			InsureName:        reApply.InsureName,
			InsureCreditCode:  reApply.InsureCreditCode,
			InsureLegalName:   reApply.InsureLegalName,
			InsureLegalIdCard: reApply.InsureLegalIdCard,
			InsureAddress:     reApply.InsureAddress,
			ApplicantName:     reApply.ApplicantName,
			ApplicantIdCard:   reApply.ApplicantIdCard,
			ApplicantTel:      reApply.ApplicantTel,
			ApplicantAuthCode: reApply.ApplicantAuthCode,
			AttachInfo:        &attachInfoString,
			AuditStatus:       &auditStatus,
			AuditOpinion:      &auditOpinion,
			AuditDate:         &auditDate,
		}
		if err = tx.Create(&apply).Error; err != nil {
			return err
		}
		var project lgjx.Project
		err = tx.Model(&lgjx.Project{}).Where("project_no = ? AND is_enable = TRUE", apply.ProjectNo).First(&project).Error
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
		} else {
			order.ProjectID = &project.ID
		}
		order.ApplyID = &apply.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		resApply = jrresponse.JRAPIApply{
			OrderNo:      reApply.OrderNo,
			ApplyNo:      reApply.ApplyNo,
			AuditStatus:  &auditStatus,
			AuditOpinion: &auditOpinion,
			AuditDate:    &auditDate,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) PayPush(rePayPush jrrequest.JRAPIPayPush) (resPayPush jrresponse.JRAPIPayPush, err error) {
	if rePayPush.OrderNo == nil ||
		rePayPush.PayNo == nil ||
		rePayPush.PayAmount == nil ||
		rePayPush.PayTime == nil ||
		rePayPush.PayTransNo == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var order lgjx.Order
		if err = tx.Where("order_no = ?", rePayPush.OrderNo).Preload(clause.Associations).Preload("Project.Template").First(&order).Error; err != nil {
			return err
		}
		pay := &lgjx.Pay{
			OrderID:    &order.ID,
			OrderNo:    rePayPush.OrderNo,
			PayNo:      rePayPush.PayNo,
			PayAmount:  rePayPush.PayAmount,
			PayTime:    rePayPush.PayTime,
			PayTransNo: rePayPush.PayTransNo,
		}
		if err = tx.Create(&pay).Error; err != nil {
			return err
		}
		order.PayID = &pay.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resPayPush = jrresponse.JRAPIPayPush{
			ReceiveResult: &receiveResult,
		}

		err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(subTx *gorm.DB) error {
			var templateFile lgjx.File
			if err = subTx.Model(&lgjx.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
				return err
			}
			var letter lgjx.Letter
			var file lgjx.File
			var encryptFile lgjx.File
			if letter, file, encryptFile, err = lgjx2.OpenLetter(order, templateFile, true); err != nil {
				return err
			}
			if err = subTx.Create(&file).Error; err != nil {
				return err
			}
			if err = subTx.Create(&encryptFile).Error; err != nil {
				return err
			}
			letter.ElogFileID = &file.ID
			letter.ElogEncryptFileID = &encryptFile.ID
			if err = subTx.Create(&letter).Error; err != nil {
				return err
			}
			order.LetterID = &letter.ID
			if err = subTx.Save(&order).Error; err != nil {
				return err
			}
			if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
				apiPath := "/jrapi/lg/lgResultPush"
				var lgResultPush = jrclientrequest.LgResultPush{
					OrderNo:             *order.OrderNo,
					ElogNo:              *letter.ElogNo,
					InsuranceName:       *letter.InsuranceName,
					InsuranceCreditCode: *letter.InsuranceCreditCode,
					ElogOutDate:         *letter.ElogOutDate,
					ElogUrl:             *letter.ElogUrl,
					ElogEncryptUrl:      *letter.ElogEncryptUrl,
					TenderDeposit:       *letter.TenderDeposit,
					InsureStartDate:     *letter.InsureStartDate,
					InsureEndDate:       *letter.InsureEndDate,
					InsureDay:           *letter.InsureDay,
					ValidateCode:        *letter.ValidateCode,
				}
				req, err := lgjx2.GenJRRequest(lgResultPush)
				if err != nil {
					return err
				}
				var res jrresponse.JRResponse
				client := resty.New()
				resp, err := client.R().
					SetBody(&req).
					SetResult(&res).
					Post(global.GVA_CONFIG.Insurance.JRAPIDomainTest + apiPath)
				if err != nil {
					return err
				}
				if resp.StatusCode() == http.StatusOK {
					if res.Code != 0 {
						code := (jrapi.ResponseCode)(res.Code)
						err := errors.New(code.String())
						global.GVA_LOG.Error("调用"+apiPath+"失败", zap.Error(err))
						return err
					} else {
						byteEncryptData, err := base64.StdEncoding.DecodeString(res.Data)
						if err != nil {
							return err
						}
						jsonData, err := lgjx2.Sm4Decrypt(byteEncryptData)
						if err != nil {
							return err
						}
						var resData jrclientresponse.Response
						err = json.Unmarshal([]byte(jsonData), &resData)
						if err != nil {
							return err
						}
						if resData.ReceiveResult != "success" {
							global.GVA_LOG.Error("调用"+apiPath+"结果不为success", zap.Error(err))
							return errors.New("接收结果不为success")
						}
					}
				} else {
					return errors.New("交易中心响应失败")
				}
			}

			return nil
		})
		if err != nil {
			fmt.Println(err.Error())
		}

		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) QueryInfo(reQueryInfo jrrequest.JRAPIQueryInfo) (resQueryInfo jrresponse.JRAPIQueryInfo, err error) {
	if reQueryInfo.ElogNo == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var letter lgjx.Letter
		if err = tx.Where("elog_no = ?", reQueryInfo.ElogNo).Preload("Order").Preload("Order.Apply").First(&letter).Error; err != nil {
			return err
		}
		resQueryInfo = jrresponse.JRAPIQueryInfo{
			OrderNo:             letter.OrderNo,
			ElogNo:              letter.ElogNo,
			ProductNo:           letter.Order.Apply.ProductNo,
			ProductType:         letter.Order.Apply.ProductType,
			ProductRate:         letter.Order.Apply.ProductRate,
			InsuranceName:       letter.InsuranceName,
			InsuranceCreditCode: letter.InsuranceCreditCode,
			EloOutDate:          letter.ElogOutDate,
			EloUrl:              letter.ElogUrl,
			EloEncryptUrl:       letter.ElogEncryptUrl,
			TenderDeposit:       letter.TenderDeposit,
			InsureStartDate:     letter.InsureStartDate,
			InsureEndDate:       letter.InsureEndDate,
			InsureDay:           letter.InsureDay,
			ValidateCode:        letter.ValidateCode,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) RevokePush(reRevokePush jrrequest.JRAPIRevokePush) (resRevokePush jrresponse.JRAPIRevokePush, err error) {
	if reRevokePush.OrderNo == nil ||
		reRevokePush.RevokeOrigin == nil ||
		reRevokePush.RevokeReason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var order lgjx.Order
		if err = tx.Where("order_no = ?", reRevokePush.OrderNo).First(&order).Error; err != nil {
			return err
		}
		revoke := &lgjx.Revoke{
			OrderID:      &order.ID,
			OrderNo:      reRevokePush.OrderNo,
			RevokeOrigin: reRevokePush.RevokeOrigin,
			RevokeReason: reRevokePush.RevokeReason,
		}
		if err = tx.Create(&revoke).Error; err != nil {
			return err
		}
		order.RevokeID = &revoke.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resRevokePush = jrresponse.JRAPIRevokePush{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) ApplyDelay(reApplyDelay jrrequest.JRAPIApplyDelay) (resApplyDelay jrresponse.JRAPIApplyDelay, err error) {
	if reApplyDelay.OrderNo == nil ||
		reApplyDelay.ApplyNo == nil ||
		reApplyDelay.ElogNo == nil ||
		reApplyDelay.ProjectGuid == nil ||
		reApplyDelay.ProjectName == nil ||
		reApplyDelay.ProjectNo == nil ||
		reApplyDelay.TenderDeposit == nil ||
		reApplyDelay.DepositStartDate == nil ||
		reApplyDelay.DepositEndDate == nil ||
		reApplyDelay.OpenBeginDate == nil ||
		reApplyDelay.InsureName == nil ||
		reApplyDelay.InsureCreditCode == nil ||
		reApplyDelay.ApplicantName == nil ||
		reApplyDelay.ApplicantIdCard == nil ||
		reApplyDelay.ApplicantTel == nil ||
		reApplyDelay.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var order lgjx.Order
		if err = tx.Where("order_no = ?", reApplyDelay.OrderNo).First(&order).Error; err != nil {
			return err
		}
		attachInfo, _ := json.Marshal(reApplyDelay.AttachInfo)
		attachInfoString := string(attachInfo)
		delay := &lgjx.Delay{
			OrderID:          &order.ID,
			OrderNo:          reApplyDelay.OrderNo,
			ApplyNo:          reApplyDelay.ApplyNo,
			ElogNo:           reApplyDelay.ElogNo,
			ProjectGuid:      reApplyDelay.ProjectGuid,
			ProjectName:      reApplyDelay.ProjectName,
			ProjectNo:        reApplyDelay.ProjectNo,
			TenderDeposit:    reApplyDelay.TenderDeposit,
			DepositStartDate: reApplyDelay.DepositStartDate,
			DepositEndDate:   reApplyDelay.DepositEndDate,
			OpenBeginDate:    reApplyDelay.OpenBeginDate,
			InsureName:       reApplyDelay.InsureName,
			InsureCreditCode: reApplyDelay.InsureCreditCode,
			ApplicantName:    reApplyDelay.ApplicantName,
			ApplicantIdCard:  reApplyDelay.ApplicantIdCard,
			ApplicantTel:     reApplyDelay.ApplicantTel,
			Reason:           reApplyDelay.Reason,
			AttachInfo:       &attachInfoString,
		}
		if err = tx.Create(&delay).Error; err != nil {
			return err
		}
		order.DelayID = &delay.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resApplyDelay = jrresponse.JRAPIApplyDelay{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) ApplyRefund(reApplyRefund jrrequest.JRAPIApplyRefund) (resApplyRefund jrresponse.JRAPIApplyRefund, err error) {
	if reApplyRefund.OrderNo == nil ||
		reApplyRefund.ApplyNo == nil ||
		reApplyRefund.ElogNo == nil ||
		reApplyRefund.InsureName == nil ||
		reApplyRefund.InsureCreditCode == nil ||
		reApplyRefund.ApplicantName == nil ||
		reApplyRefund.ApplicantIdCard == nil ||
		reApplyRefund.ApplicantTel == nil ||
		reApplyRefund.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var order lgjx.Order
		if err = tx.Where("order_no = ?", reApplyRefund.OrderNo).First(&order).Error; err != nil {
			return err
		}
		attachInfo, _ := json.Marshal(reApplyRefund.AttachInfo)
		attachInfoString := string(attachInfo)
		refund := &lgjx.Refund{
			OrderID:          &order.ID,
			OrderNo:          reApplyRefund.OrderNo,
			ApplyNo:          reApplyRefund.ApplyNo,
			ElogNo:           reApplyRefund.ElogNo,
			InsureName:       reApplyRefund.InsureName,
			InsureCreditCode: reApplyRefund.InsureCreditCode,
			ApplicantName:    reApplyRefund.ApplicantName,
			ApplicantIdCard:  reApplyRefund.ApplicantIdCard,
			ApplicantTel:     reApplyRefund.ApplicantTel,
			Reason:           reApplyRefund.Reason,
			AttachInfo:       &attachInfoString,
		}
		if err = tx.Create(&refund).Error; err != nil {
			return err
		}
		order.RefundID = &refund.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resApplyRefund = jrresponse.JRAPIApplyRefund{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) ApplyClaim(reApplyClaim jrrequest.JRAPIApplyClaim) (resApplyClaim jrresponse.JRAPIApplyClaim, err error) {
	if reApplyClaim.OrderNo == nil ||
		reApplyClaim.ApplyNo == nil ||
		reApplyClaim.ElogNo == nil ||
		reApplyClaim.InsuredName == nil ||
		reApplyClaim.InsuredCreditCode == nil ||
		reApplyClaim.InsuredBankNo == nil ||
		reApplyClaim.InsuredBankName == nil ||
		reApplyClaim.ApplicantName == nil ||
		reApplyClaim.ApplicantIdCard == nil ||
		reApplyClaim.ApplicantTel == nil ||
		reApplyClaim.ClaimAmount == nil ||
		reApplyClaim.Reason == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var order lgjx.Order
		if err = tx.Where("order_no = ?", reApplyClaim.OrderNo).First(&order).Error; err != nil {
			return err
		}
		attachInfo, _ := json.Marshal(reApplyClaim.AttachInfo)
		attachInfoString := string(attachInfo)
		claim := &lgjx.Claim{
			OrderID:           &order.ID,
			OrderNo:           reApplyClaim.OrderNo,
			ApplyNo:           reApplyClaim.ApplyNo,
			ElogNo:            reApplyClaim.ElogNo,
			InsuredName:       reApplyClaim.InsuredName,
			InsuredCreditCode: reApplyClaim.InsuredCreditCode,
			InsuredBankNo:     reApplyClaim.InsuredBankNo,
			InsuredBankName:   reApplyClaim.InsuredBankName,
			ApplicantName:     reApplyClaim.ApplicantName,
			ApplicantIdCard:   reApplyClaim.ApplicantIdCard,
			ApplicantTel:      reApplyClaim.ApplicantTel,
			ClaimAmount:       reApplyClaim.ClaimAmount,
			Reason:            reApplyClaim.Reason,
			AttachInfo:        &attachInfoString,
		}
		if err = tx.Create(&claim).Error; err != nil {
			return err
		}
		order.ClaimID = &claim.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resApplyClaim = jrresponse.JRAPIApplyClaim{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) LogoutPush(reLogoutPush jrrequest.JRAPILogoutPush) (resLogoutPush jrresponse.JRAPILogoutPush, err error) {
	if reLogoutPush.ProjectGuid == nil ||
		reLogoutPush.ProjectName == nil ||
		reLogoutPush.ProjectNo == nil ||
		reLogoutPush.Reason == nil ||
		reLogoutPush.LogoutType == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		logout := &lgjx.Logout{
			ProjectGuid:         reLogoutPush.ProjectGuid,
			ProjectName:         reLogoutPush.ProjectName,
			ProjectNo:           reLogoutPush.ProjectNo,
			Reason:              reLogoutPush.Reason,
			LogoutType:          reLogoutPush.LogoutType,
			WinBidderName:       reLogoutPush.WinBidderName,
			WinBidderCreditCode: reLogoutPush.WinBidderCreditCode,
		}
		if err = tx.Create(&logout).Error; err != nil {
			return err
		}
		if err = tx.Model(&lgjx.Order{}).Joins("Apply").Where("Apply.project_guid = ?", reLogoutPush.ProjectGuid).Update("logout_id", logout.ID).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resLogoutPush = jrresponse.JRAPILogoutPush{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) ApplyInvoice(reApplyInvoice jrrequest.JRAPIApplyInvoice) (resApplyInvoice jrresponse.JRAPIApplyInvoice, err error) {
	if reApplyInvoice.ApplyNo == nil ||
		reApplyInvoice.InvoiceTotalAmount == nil ||
		reApplyInvoice.InvoiceType == nil ||
		reApplyInvoice.InvoiceTileType == nil ||
		reApplyInvoice.InvoiceTile == nil ||
		reApplyInvoice.OrderList == nil {
		err = errors.New(jrapi.MissingServiceParam.String())
		return
	}
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		orderList, _ := json.Marshal(reApplyInvoice.OrderList)
		orderListString := string(orderList)
		invoiceApply := &lgjx.InvoiceApply{
			ApplyNo:            reApplyInvoice.ApplyNo,
			InvoiceTotalAmount: reApplyInvoice.InvoiceTotalAmount,
			InvoiceType:        reApplyInvoice.InvoiceType,
			InvoiceTileType:    reApplyInvoice.InvoiceTileType,
			InvoiceTile:        reApplyInvoice.InvoiceTile,
			TaxNo:              reApplyInvoice.TaxNo,
			BankName:           reApplyInvoice.BankName,
			BankNo:             reApplyInvoice.BankNo,
			CompanyAddress:     reApplyInvoice.CompanyAddress,
			CompanyTel:         reApplyInvoice.CompanyTel,
			Remarks:            reApplyInvoice.Remarks,
			OrderList:          &orderListString,
		}
		if err = tx.Create(&invoiceApply).Error; err != nil {
			return err
		}
		receiveResult := "success"
		resApplyInvoice = jrresponse.JRAPIApplyInvoice{
			ReceiveResult: &receiveResult,
		}
		return nil
	})
	return
}

func (testJRAPIService *TestJRAPIService) LetterFileDownload(elog string, encrypt bool) (file lgjx.File, err error) {
	var letter lgjx.Letter
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Letter{})
	if encrypt {
		db = db.Where("elog_encrypt_url = ?", elog)
	} else {
		db = db.Where("elog_url = ?", elog)
	}
	err = db.Preload("ElogFile").Preload("ElogEncryptFile").First(&letter).Error
	if encrypt {
		file = *letter.ElogEncryptFile
	} else {
		file = *letter.ElogFile
	}
	return
}
