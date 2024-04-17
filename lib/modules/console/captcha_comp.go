package console

import (
	"fmt"
	"time"

	"github.com/zmloong/lotus/core"
	"github.com/zmloong/lotus/core/cbase"
	"github.com/zmloong/lotus/sys/sdks/email"
)

type CaptchaComp struct {
	cbase.ModuleCompBase
	module IConsole
	email  email.IEmail
}

func (this *CaptchaComp) Init(service core.IService, module core.IModule, comp core.IModuleComp, options core.IModuleOptions) (err error) {
	err = this.ModuleCompBase.Init(service, module, comp, options)
	this.module = module.(IConsole)
	this.email, err = email.NewSys(
		email.SetServerhost(this.module.Options().GetMailServerhost()),
		email.SetFromemail(this.module.Options().GetMailFromemail()),
		email.SetFompasswd(this.module.Options().GetMailFompasswd()),
		email.SetServerport(this.module.Options().GetMailServerport()))
	return
}

// 发送邮箱验证码
func (this *CaptchaComp) SendEmailCaptcha(email, captcha string) (err error) {
	return this.email.SendMail(email, "Verification Code", fmt.Sprintf("Your %s console verification code:<b>%s</b>, please do not forward others", this.module.Options().GetProjectName(), captcha))
}

// 获取验证码
func (this *CaptchaComp) QueryCaptcha(cId string, ctype CaptchaType) (captcha string, err error) {
	Id := fmt.Sprintf(string(Cache_ConsoleCaptcha), cId, ctype)
	redis := this.module.Cache().GetRedis()
	err = redis.Get(Id, &captcha)
	return
}

// 写入验证码
func (this *CaptchaComp) WriteCaptcha(cId, captcha string, ctype CaptchaType) {
	Id := fmt.Sprintf(string(Cache_ConsoleCaptcha), cId, ctype)
	redis := this.module.Cache().GetRedis()
	redis.Set(Id, captcha, time.Second*time.Duration(this.module.Options().GetCaptchaExpirationdate()))
}
