package lgjx

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrclientrequest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrclientresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/jrapi/jrresponse"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/nonmigrate"
	lgjxReq "github.com/flipped-aurora/gin-vue-admin/server/model/lgjx/request"
	lgjx2 "github.com/flipped-aurora/gin-vue-admin/server/utils/lgjx"
	"github.com/go-resty/resty/v2"
	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"net/http"
	"strconv"
	"time"
)

type TestOrderService struct {
}

func (testOrderService *TestOrderService) CreateOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Create(&order).Error
	return err
}

func (testOrderService *TestOrderService) DeleteOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&order).Error
	return err
}

func (testOrderService *TestOrderService) DeleteOrderByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Delete(&[]lgjx.Order{}, "id in ?", ids.Ids).Error
	return err
}

func (testOrderService *TestOrderService) UpdateOrder(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Save(&order).Error
	return err
}

func (testOrderService *TestOrderService) GetOrder(id uint) (order lgjx.Order, err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Where("id = ?", id).First(&order).Error
	return
}

func (testOrderService *TestOrderService) GetOrderInfoList(info lgjxReq.OrderSearch) (list []lgjx.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Order{})
	db.Joins("Apply").Joins("Project").Joins("Letter")
	var orders []lgjx.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ApplyNo != nil {
		db = db.Where("Apply.apply_no = ?", info.ApplyNo)
	}
	if info.ProjectName != nil {
		db = db.Where("Apply.project_name = ?", info.ProjectName)
	}
	if info.InsureName != nil {
		db = db.Where("Apply.insure_name = ?", info.InsureName)
	}
	if info.ElogTemplateId != nil {
		db = db.Where("Project.template_id = ?", info.ElogTemplateId)
	}
	if info.ElogNo != nil {
		db = db.Where("Letter.elog_no = ?", info.ElogNo)
	}
	if info.OrderStatus != nil {
		if *info.OrderStatus == "已撤" {
			db = db.Where("order.revoke_id is not null")
		}
		if *info.OrderStatus == "销函" {
			db = db.Where("order.logout_id is not null")
		}
		if *info.OrderStatus == "理赔" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is not null AND Claim.audit_status = 2")
		}
		if *info.OrderStatus == "退函" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is not null AND Refund.audit_status = 2")
		}
		if *info.OrderStatus == "延期" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is not null AND Delay.audit_status = 2")
		}
		if *info.OrderStatus == "已开" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is null")
			db = db.Where("order.letter_id is not null")
		}
		if *info.OrderStatus == "未开" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is null")
			db = db.Where("order.letter_id is null")
		}
	}
	if info.AuditStatus != nil {
		db = db.Where("Apply.audit_status = ?", info.AuditStatus)
	}
	if info.OpenBeginDate != nil {
		db = db.Where("Apply.open_begin_date BETWEEN ? AND ?", info.OpenBeginDate[0], info.OpenBeginDate[1])
	}
	if info.ApplyCreatedAt != nil {
		db = db.Where("Apply.created_at BETWEEN ? AND ?", info.ApplyCreatedAt[0], info.ApplyCreatedAt[1])
	}
	if info.LetterCreatedAt != nil {
		db = db.Where("Letter.created_at BETWEEN ? AND ?", info.LetterCreatedAt[0], info.LetterCreatedAt[1])
	}
	if info.InsureDay != nil {
		db = db.Where("Letter.insure_day = ?", info.InsureDay)
	}
	if info.AuditDelay != nil {
		db = db.Where("order.delay_id is not null")
	}
	if info.AuditRefund != nil {
		db = db.Where("order.refund_id is not null")
	}
	if info.AuditClaim != nil {
		db = db.Where("order.claim_id is not null")
	}
	if info.IsPayed != nil {
		db = db.Where("order.pay_id is not null")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).
		Preload(clause.Associations).
		Preload("Project.Template").
		Preload("Letter.ElogFile").
		Preload("Letter.ElogEncryptFile").
		Preload("Delay.ElogFile").
		Preload("Delay.ElogEncryptFile").
		Order("order.created_at desc").Offset(offset).Find(&orders).Error
	return orders, total, err
}

func (testOrderService *TestOrderService) ApproveApply(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Apply.AuditStatus = &auditStatus
		order.Apply.AuditOpinion = &auditOpinion
		order.Apply.AuditDate = &auditDate
		realAmount := math.Trunc(*order.Project.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		if realAmount < global.GVA_CONFIG.Insurance.ElogMinAmount {
			order.Apply.RealElogAmount = &global.GVA_CONFIG.Insurance.ElogMinAmount
		} else {
			order.Apply.RealElogAmount = &realAmount
		}
		order.Apply.RealElogRate = &global.GVA_CONFIG.Insurance.ElogRate
		order.Apply.InsuranceName = &global.GVA_CONFIG.Insurance.Name
		order.Apply.InsuranceCreditCode = &global.GVA_CONFIG.Insurance.CreditCode
		err := tx.Save(&order.Apply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/applyPush"
			var applyPush = jrclientrequest.ApplyPush{
				OrderNo:             *order.OrderNo,
				ApplyNo:             *order.Apply.ApplyNo,
				AuditStatus:         *order.Apply.AuditStatus,
				AuditOpinion:        *order.Apply.AuditOpinion,
				AuditDate:           *order.Apply.AuditDate,
				RealElogAmount:      *order.Apply.RealElogAmount,
				RealElogRate:        *order.Apply.RealElogRate,
				TenderDeposit:       *order.Apply.TenderDeposit,
				InsuranceName:       *order.Apply.InsuranceName,
				InsuranceCreditCode: *order.Apply.InsuredCreditCode,
			}
			req, err := lgjx2.GenJRRequest(applyPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) RejectApply(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Apply.AuditStatus = &auditStatus
		order.Apply.AuditOpinion = &auditOpinion
		order.Apply.AuditDate = &auditDate
		realAmount := math.Trunc(*order.Project.TenderDeposit*global.GVA_CONFIG.Insurance.ElogRate*1e2+0.5) * 1e-2
		if realAmount < global.GVA_CONFIG.Insurance.ElogMinAmount {
			order.Apply.RealElogAmount = &global.GVA_CONFIG.Insurance.ElogMinAmount
		} else {
			order.Apply.RealElogAmount = &realAmount
		}
		order.Apply.RealElogRate = &global.GVA_CONFIG.Insurance.ElogRate
		order.Apply.InsuranceName = &global.GVA_CONFIG.Insurance.Name
		order.Apply.InsuranceCreditCode = &global.GVA_CONFIG.Insurance.CreditCode
		err := tx.Save(&order.Apply).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/applyPush"
			var applyPush = jrclientrequest.ApplyPush{
				OrderNo:             *order.OrderNo,
				ApplyNo:             *order.Apply.ApplyNo,
				AuditStatus:         *order.Apply.AuditStatus,
				AuditOpinion:        *order.Apply.AuditOpinion,
				AuditDate:           *order.Apply.AuditDate,
				RealElogAmount:      *order.Apply.RealElogAmount,
				RealElogRate:        *order.Apply.RealElogRate,
				TenderDeposit:       *order.Apply.TenderDeposit,
				InsuranceName:       *order.Apply.InsuranceName,
				InsuranceCreditCode: *order.Apply.InsuredCreditCode,
			}
			req, err := lgjx2.GenJRRequest(applyPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) ApproveDelay(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Delay.AuditStatus = &auditStatus
		order.Delay.AuditOpinion = &auditOpinion
		order.Delay.AuditDate = &auditDate

		var templateFile lgjx.File
		if err = tx.Model(&lgjx.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}

		var letter lgjx.Letter
		var file lgjx.File
		var encryptFile lgjx.File

		if letter, file, encryptFile, err = lgjx2.OpenLetter(order, templateFile, true); err != nil {
			return err
		}
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		if err = tx.Create(&encryptFile).Error; err != nil {
			return err
		}
		order.Delay.ElogFileID = &file.ID
		order.Delay.ElogEncryptFileID = &encryptFile.ID

		order.Delay.ElogUrl = letter.ElogUrl
		order.Delay.ElogEncryptUrl = letter.ElogEncryptUrl
		order.Delay.TenderDeposit = letter.TenderDeposit
		order.Delay.InsureStartDate = letter.InsureStartDate
		order.Delay.InsureEndDate = letter.InsureEndDate
		order.Delay.InsureDay = letter.InsureDay
		order.Delay.ValidateCode = letter.ValidateCode
		err := tx.Save(&order.Delay).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/delayPush"
			var delayPush = jrclientrequest.DelayPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Delay.ApplyNo,
				ElogNo:          *order.Delay.ElogNo,
				AuditStatus:     *order.Delay.AuditStatus,
				AuditOpinion:    *order.Delay.AuditOpinion,
				AuditDate:       *order.Delay.AuditDate,
				ElogUrl:         global.GVA_CONFIG.Insurance.APIDoaminTest + "/delayFileDownload?elog=" + *letter.ElogUrl,
				ElogEncryptUrl:  global.GVA_CONFIG.Insurance.APIDoaminTest + "/delayFileDownload?elog=" + *letter.ElogEncryptUrl + "&type=encrypt",
				TenderDeposit:   *letter.TenderDeposit,
				InsureStartDate: *letter.InsureStartDate,
				InsureEndDate:   *letter.InsureEndDate,
				InsureDay:       *letter.InsureDay,
				ValidateCode:    *letter.ValidateCode,
			}
			req, err := lgjx2.GenJRRequest(delayPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) RejectDelay(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Delay.AuditStatus = &auditStatus
		order.Delay.AuditOpinion = &auditOpinion
		order.Delay.AuditDate = &auditDate

		var templateFile lgjx.File
		if err = tx.Model(&lgjx.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		err := tx.Save(&order.Delay).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/delayPush"
			var delayPush = jrclientrequest.DelayPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Delay.ApplyNo,
				ElogNo:          *order.Letter.ElogNo,
				AuditStatus:     *order.Delay.AuditStatus,
				AuditOpinion:    *order.Delay.AuditOpinion,
				AuditDate:       *order.Delay.AuditDate,
				ElogUrl:         *order.Letter.ElogUrl,
				ElogEncryptUrl:  *order.Letter.ElogEncryptUrl,
				TenderDeposit:   *order.Letter.TenderDeposit,
				InsureStartDate: *order.Letter.InsureStartDate,
				InsureEndDate:   *order.Letter.InsureEndDate,
				InsureDay:       *order.Letter.InsureDay,
				ValidateCode:    *order.Letter.ValidateCode,
			}
			req, err := lgjx2.GenJRRequest(delayPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) ApproveRefund(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Refund.AuditStatus = &auditStatus
		order.Refund.AuditOpinion = &auditOpinion
		order.Refund.AuditDate = &auditDate
		order.Refund.PayAmount = order.Pay.PayAmount
		err := tx.Save(&order.Refund).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/refundPush"
			var refundPush = jrclientrequest.RefundPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Refund.ApplyNo,
				ElogNo:       *order.Refund.ElogNo,
				AuditStatus:  *order.Refund.AuditStatus,
				AuditOpinion: *order.Refund.AuditOpinion,
				AuditDate:    *order.Refund.AuditDate,
				PayAmount:    *order.Refund.PayAmount,
			}
			req, err := lgjx2.GenJRRequest(refundPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) RejectRefund(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Refund.AuditStatus = &auditStatus
		order.Refund.AuditOpinion = &auditOpinion
		order.Refund.AuditDate = &auditDate
		order.Refund.PayAmount = order.Pay.PayAmount
		err := tx.Save(&order.Refund).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/refundPush"
			var refundPush = jrclientrequest.RefundPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Refund.ApplyNo,
				ElogNo:       *order.Refund.ElogNo,
				AuditStatus:  *order.Refund.AuditStatus,
				AuditOpinion: *order.Refund.AuditOpinion,
				AuditDate:    *order.Refund.AuditDate,
			}
			req, err := lgjx2.GenJRRequest(refundPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) ApproveClaim(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(2)
		auditOpinion := "受理成功"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Claim.AuditStatus = &auditStatus
		order.Claim.AuditOpinion = &auditOpinion
		order.Claim.AuditDate = &auditDate
		order.Claim.RealClaimAmount = order.Claim.ClaimAmount
		order.Claim.RealClaimDate = &auditDate
		err := tx.Save(&order.Claim).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/claimPush"
			var claimPush = jrclientrequest.ClaimPush{
				OrderNo:         *order.OrderNo,
				ApplyNo:         *order.Claim.ApplyNo,
				ElogNo:          *order.Claim.ElogNo,
				AuditStatus:     *order.Claim.AuditStatus,
				AuditOpinion:    *order.Claim.AuditOpinion,
				AuditDate:       *order.Claim.AuditDate,
				RealClaimAmount: *order.Claim.RealClaimAmount,
				RealClaimDate:   *order.Claim.RealClaimDate,
			}
			req, err := lgjx2.GenJRRequest(claimPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) RejectClaim(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		auditStatus := int64(3)
		auditOpinion := "受理失败"
		auditDate := time.Now().Format("2006-01-02 15:04:05")
		order.Claim.AuditStatus = &auditStatus
		order.Claim.AuditOpinion = &auditOpinion
		order.Claim.AuditDate = &auditDate
		err := tx.Save(&order.Claim).Error
		if err != nil {
			return err
		}

		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/claimPush"
			var claimPush = jrclientrequest.ClaimPush{
				OrderNo:      *order.OrderNo,
				ApplyNo:      *order.Claim.ApplyNo,
				ElogNo:       *order.Claim.ElogNo,
				AuditStatus:  *order.Claim.AuditStatus,
				AuditOpinion: *order.Claim.AuditOpinion,
				AuditDate:    *order.Claim.AuditDate,
			}
			req, err := lgjx2.GenJRRequest(claimPush)
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) OpenLetter(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var templateFile lgjx.File
		if err = tx.Model(&lgjx.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		if err = tx.Where("order_no = ?", order.OrderNo).Delete(&lgjx.Letter{}).Error; err != nil {
			return err
		}
		var letter lgjx.Letter
		var file lgjx.File
		var encryptFile lgjx.File
		if letter, file, encryptFile, err = lgjx2.OpenLetter(order, templateFile, true); err != nil {
			return err
		}
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		if err = tx.Create(&encryptFile).Error; err != nil {
			return err
		}
		letter.ElogFileID = &file.ID
		letter.ElogEncryptFileID = &encryptFile.ID
		if err = tx.Create(&letter).Error; err != nil {
			return err
		}
		order.LetterID = &letter.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/lgResultPush"
			var lgResultPush = jrclientrequest.LgResultPush{
				OrderNo:             *letter.OrderNo,
				ElogNo:              *letter.ElogNo,
				InsuranceName:       *letter.InsuranceName,
				InsuranceCreditCode: *letter.InsuranceCreditCode,
				ElogOutDate:         *letter.ElogOutDate,
				ElogUrl:             global.GVA_CONFIG.Insurance.APIDoaminTest + "/letterFileDownload?elog=" + *letter.ElogUrl,
				ElogEncryptUrl:      global.GVA_CONFIG.Insurance.APIDoaminTest + "/letterFileDownload?elog=" + *letter.ElogEncryptUrl + "&type=encrypt",
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) RePush(order lgjx.Order) (err error) {
	err = global.MustGetGlobalDBByDBName("lg-jx-test").Transaction(func(tx *gorm.DB) error {
		var templateFile lgjx.File
		if err = tx.Model(&lgjx.File{}).Where("id = ?", *order.Project.Template.TemplateFileID).First(&templateFile).Error; err != nil {
			return err
		}
		if err = tx.Where("order_no = ?", order.OrderNo).Delete(&lgjx.Letter{}).Error; err != nil {
			return err
		}
		var letter lgjx.Letter
		var file lgjx.File
		var encryptFile lgjx.File
		if letter, file, encryptFile, err = lgjx2.OpenLetter(order, templateFile, true); err != nil {
			return err
		}
		if err = tx.Create(&file).Error; err != nil {
			return err
		}
		if err = tx.Create(&encryptFile).Error; err != nil {
			return err
		}
		letter.ElogFileID = &file.ID
		letter.ElogEncryptFileID = &encryptFile.ID
		if err = tx.Create(&letter).Error; err != nil {
			return err
		}
		order.LetterID = &letter.ID
		if err = tx.Save(&order).Error; err != nil {
			return err
		}
		if global.GVA_CONFIG.Insurance.JRAPIDomainTest != "" {
			apiPath := "/jrapi/lg/lgResultPush"
			var lgResultPush = jrclientrequest.LgResultPush{
				OrderNo:             *letter.OrderNo,
				ElogNo:              *letter.ElogNo,
				InsuranceName:       *letter.InsuranceName,
				InsuranceCreditCode: *letter.InsuranceCreditCode,
				ElogOutDate:         *letter.ElogOutDate,
				ElogUrl:             global.GVA_CONFIG.Insurance.APIDoaminTest + "/letterFileDownload?elog=" + *letter.ElogUrl,
				ElogEncryptUrl:      global.GVA_CONFIG.Insurance.APIDoaminTest + "/letterFileDownload?elog=" + *letter.ElogEncryptUrl + "&type=encrypt",
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
					err := errors.New(res.Msg)
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
	return err
}

func (testOrderService *TestOrderService) GetOrderStatisticData() (orderStatisticData nonmigrate.OrderStatisticData, err error) {
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Order{})
	db.Joins("Pay").Joins("Letter").Joins("Refund").Joins("Claim")
	//未撤函，未理赔，未退函
	db = db.Where("order.revoke_id is null")
	db = db.Where("order.claim_id is null")
	db = db.Where("order.refund_id is null")

	err = db.Select("COALESCE(SUM(Letter.tender_deposit), 0) as TotalGuaranteeAmount, COALESCE(SUM(Pay.pay_amount), 0) as TotalElogAmount").Scan(&orderStatisticData).Error

	return orderStatisticData, err
}

func (testOrderService *TestOrderService) ExportExcel(info lgjxReq.OrderSearch) (excelData []byte, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("lg-jx-test").Model(&lgjx.Order{})
	db.Joins("Apply").Joins("Pay").Joins("Letter").Joins("Revoke").Joins("Delay").Joins("Refund").Joins("Claim").Joins("Logout").Joins("Invoice").Joins("Project")
	var orders []lgjx.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ApplyNo != nil {
		db = db.Where("Apply.apply_no = ?", info.ApplyNo)
	}
	if info.ProjectName != nil {
		db = db.Where("Apply.project_name = ?", info.ProjectName)
	}
	if info.InsureName != nil {
		db = db.Where("Apply.insure_name = ?", info.InsureName)
	}
	if info.ElogTemplateId != nil {
		db = db.Where("Project.template_id = ?", info.ElogTemplateId)
	}
	if info.ElogNo != nil {
		db = db.Where("Letter.elog_no = ?", info.ElogNo)
	}
	if info.OrderStatus != nil {
		if *info.OrderStatus == "已撤" {
			db = db.Where("order.revoke_id is not null")
		}
		if *info.OrderStatus == "销函" {
			db = db.Where("order.logout_id is not null")
		}
		if *info.OrderStatus == "理赔" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is not null AND Claim.audit_status = 2")
		}
		if *info.OrderStatus == "退函" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is not null AND Refund.audit_status = 2")
		}
		if *info.OrderStatus == "延期" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is not null AND Delay.audit_status = 2")
		}
		if *info.OrderStatus == "已开" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is null")
			db = db.Where("order.letter_id is not null")
		}
		if *info.OrderStatus == "未开" {
			db = db.Where("order.logout_id is null")
			db = db.Where("order.claim_id is null")
			db = db.Where("order.refund_id is null")
			db = db.Where("order.delay_id is null")
			db = db.Where("order.letter_id is null")
		}
	}
	if info.AuditStatus != nil {
		db = db.Where("Apply.audit_status = ?", info.AuditStatus)
	}
	if info.OpenBeginDate != nil {
		db = db.Where("Apply.open_begin_date BETWEEN ? AND ?", info.OpenBeginDate[0], info.OpenBeginDate[1])
	}
	if info.ApplyCreatedAt != nil {
		db = db.Where("Apply.created_at BETWEEN ? AND ?", info.ApplyCreatedAt[0], info.ApplyCreatedAt[1])
	}
	if info.LetterCreatedAt != nil {
		db = db.Where("Letter.created_at BETWEEN ? AND ?", info.LetterCreatedAt[0], info.LetterCreatedAt[1])
	}
	if info.InsureDay != nil {
		db = db.Where("Letter.insure_day = ?", info.InsureDay)
	}
	if info.AuditDelay != nil {
		db = db.Where("order.delay_id is not null")
	}
	if info.AuditRefund != nil {
		db = db.Where("order.refund_id is not null")
	}
	if info.AuditClaim != nil {
		db = db.Where("order.claim_id is not null")
	}
	if info.IsPayed != nil {
		db = db.Where("order.pay_id is not null")
	}

	err = db.Limit(limit).Preload(clause.Associations).Order("order.created_at desc").Offset(offset).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	excel := excelize.NewFile()
	_ = excel.SetSheetRow("Sheet1", "A1", &[]string{"保函文件下载", "交易中心", "保函申请编码", "申请企业", "标段名称", "标段编号", "受益人名称", "担保金额（元）", "保函起始日期", "保函截止日期", "订单状态", "开标时间", "保费金额", "所属市", "所属县", "审核时间", "申请日期", "开函日期", "保函编码", "审核状态", "付款时间", "付款金额", "交易单号"})
	for i, order := range orders {
		axis := fmt.Sprintf("A%d", i+2)
		elogAmount, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", *order.Apply.TenderDeposit**order.Apply.ProductRate), 64)
		var elogUrl string
		var insureStartDate string
		var insureEndDate string
		var letterOpenDate string
		var elogNo string
		if order.LetterID != nil {
			elogUrl = global.GVA_CONFIG.Insurance.APIDoaminTest + "letterFileDownload?elog=" + *order.Letter.ElogUrl
			insureStartDate = *order.Letter.InsureStartDate
			insureEndDate = *order.Letter.InsureEndDate
			letterOpenDate = order.Letter.CreatedAt.Format("2006-01-02 15:04:05")
			elogNo = *order.Letter.ElogNo
		} else {
			elogUrl = ""
			insureStartDate = ""
			insureEndDate = ""
			letterOpenDate = ""
			elogNo = ""
		}
		var projectCity string
		var projectCounty string
		if order.ProjectID != nil {
			projectCity = *order.Project.ProjectCity
			projectCounty = *order.Project.ProjectCounty
		} else {
			projectCity = ""
			projectCounty = ""
		}
		var payTime string
		var payAmount float64
		var payTransNo string
		if order.PayID != nil {
			payTime = *order.Pay.PayTime
			payAmount = *order.Pay.PayAmount
			payTransNo = *order.Pay.PayTransNo
		} else {
			payTime = ""
			payAmount = 0.0
			payTransNo = ""
		}
		_ = excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			elogUrl,
			"江西云平台",
			*order.Apply.ApplyNo,
			*order.Apply.InsureName,
			*order.Apply.ProjectName,
			*order.Apply.ProjectNo,
			*order.Apply.InsuredName,
			*order.Apply.TenderDeposit,
			insureStartDate,
			insureEndDate,
			lgjx2.OrderStatus(order),
			*order.Apply.OpenBeginDate,
			elogAmount,
			projectCity,
			projectCounty,
			*order.Apply.AuditDate,
			order.Apply.CreatedAt.Format("2006-01-02 15:04:05"),
			letterOpenDate,
			elogNo,
			lgjx2.AuditStatus(*order.Apply.AuditStatus),
			payTime,
			payAmount,
			payTransNo,
		})
	}
	excelBuffer, err := excel.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	excelData = excelBuffer.Bytes()
	return
}
