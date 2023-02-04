package utils

import (
    "github.com/go-gomail/gomail"
    "github.com/gofrs/uuid"
    "github.com/matcornic/hermes/v2"
    "mathe-learn-platform/config"
)

// GetUUID 根据 idtype 生成32位随机数
func GetUUID(idtype string) string {
    uid := uuid.NewV5(uuid.Must(uuid.NewV4()), idtype).String()
    uid = uid[:8] + uid[9:13] + uid[14:18] + uid[19:23] + uid[24:]
    return uid
}

// GenetateEmailBody 生成邮件主体
func GenetateEmailBody(name, code string) hermes.Email {
    body := hermes.Email{
        Body: hermes.Body{
            Greeting: "尊敬的",
            Name:     name,
            Intros: []string{
                "系统检测到您最近使用了验证码服务",
                "验证码:" + code + ",有效期为5分钟, 请勿回复此邮件。",
                "如果非您本人操作, 请忽略。",
            },
            Signature: "您忠诚的",
        },
    }
    return body
}

// SendEmail 给用户发送
func SendEmail(address, code, name string) error {
    h := hermes.Hermes{
        Product: hermes.Product{
            Name:      "MLP团队敬上。",
            Logo:      "https://up.enterdesk.com/edpic/45/f0/71/45f07166bddcd586f364b249fcd802c0.jpg",
            Copyright: "Copyright © 2023 Dharma Initiative. All rights reserved.",
        },
    }
    emailBody, _ := h.GenerateHTML(GenetateEmailBody(name, code))
    m := gomail.NewMessage()
    m.SetHeader("From", config.MLPEmailAddress)
    m.SetHeader("To", address)
    m.SetHeader("Subject", "MLP用户身份确认")
    m.SetBody("text/html", emailBody)
    //m.AddAlternative("text/html", emailBody)
    d := gomail.NewDialer("smtp.qq.com", 25, config.MLPEmailAddress, config.EmailCode)
    err := d.DialAndSend(m)
    return err
}
