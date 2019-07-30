package mail

import (
	"Spider/config"
	"Spider/model"
	"Spider/utils"
	"fmt"
	"net/smtp"
	"strings"
)

var (
	user     = config.Conf.Get("email.username").(string)
	password = config.Conf.Get("email.password").(string)
	host     = config.Conf.Get("email.host").(string)
)

func Bilibili(subscribers []model.User, author model.BilibiliUp, list []model.BilibiliVideo) error {
	html := ""
	html += htmlHeader
	for _, video := range list {
		html += fmt.Sprintf(htmlInfo, video.Pic, video.Title, video.Description, video.ID, video.ID)
	}
	html += htmlFooter
	var errorAmount []model.User
	for _, subscriber := range subscribers {
		if err := SendMail(subscriber.Mail, fmt.Sprint(author.Name, " 的B站视频更新了"), html, "html"); err != nil {
			errorAmount = append(errorAmount, subscriber)
		}
	}
	utils.Logger.Info(fmt.Sprint("消息推送! 发送成功至:[]"))
	return nil
}

/*
 *    user : example@example.com login smtp server user
 *    password: xxxxx login smtp server password
 *    host: smtp.example.com:port   smtp.163.com:25
 *    to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */
func SendMail(to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

//func main() {
//	subject := "Test send email by golang"
//
//	body := `
//    <html>
//    <body>
//    <h3>
//    "这是GO语言写的测试邮件。"
//    </h3>
//    </body>
//    </html>
//    `
//	fmt.Println("send email")
//	err := SendMail(subject, body, "html")
//	if err != nil {
//		fmt.Println("send mail error!")
//		fmt.Println(err)
//	} else {
//		fmt.Println("send mail success!")
//	}
//
//}

var htmlInfo = "<div style=\"background-color: #ffffff;\">\n" +
	"                    <div class=\"layout two-col\"\n" +
	"                         style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000%% - 167400px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;\">\n" +
	"                        <div class=\"layout__inner\" style=\"border-collapse: collapse;display: table;width: 100%%;\">\n" +
	"                            <!--[if (mso)|(IE)]>\n" +
	"                            <table width=\"100%%\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                                <tr class=\"layout-full-width\" style=\"background-color: #ffffff;\">\n" +
	"                                    <td class=\"layout__edges\">&nbsp;</td>\n" +
	"                                    <td style=\"width: 300px\" valign=\"top\" class=\"w260\"><![endif]-->\n" +
	"                            <div class=\"column\"\n" +
	"                                 style=\"Float: left;max-width: 320px;min-width: 300px; width: 320px;width: calc(12300px - 2000%%);text-align: left;color: #f6f6f8;font-size: 16px;line-height: 24px;font-family: Roboto,Tahoma,sans-serif;\">\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"font-size: 12px;font-style: normal;font-weight: normal;line-height: 19px;\"\n" +
	"                                         align=\"center\">\n" +
	"                                        <img style=\"border: 0;display: block;height: auto;width: 100%%;max-width: 300px;\"\n" +
	"                                             alt=\"Tours\" width=\"260\" src=\"%s\"/>\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td>\n" +
	"                        <td style=\"width: 300px\" valign=\"top\" class=\"w260\"><![endif]-->\n" +
	"                            <div class=\"column\"\n" +
	"                                 style=\"Float: left;max-width: 320px;min-width: 300px; width: 320px;width: calc(12300px - 2000%%);text-align: left;color: #f6f6f8;font-size: 16px;line-height: 24px;font-family: Roboto,Tahoma,sans-serif;\">\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 40px;font-size: 1px;\">\n" +
	"                                        &nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;mso-text-raise: 4px;\">\n" +
	"                                        <h2 style=\"Margin-top: 0;Margin-bottom: 0;font-style: normal;font-weight: normal;color: #fff;font-size: 18px;line-height: 26px;\">\n" +
	"                                            <span style=\"color:#001b71\"><strong>%s</strong></span></h2>\n" +
	"                                        <p style=\"Margin-top: 16px;Margin-bottom: 20px;\"><span style=\"color:#001b71\">%s</span>\n" +
	"                                        </p>\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div class=\"btn btn--flat btn--large\" style=\"Margin-bottom: 20px;text-align: left;\">\n" +
	"                                        <![if !mso]><a\n" +
	"                                            style=\"border-radius: 4px;display: inline-block;font-size: 14px;font-weight: bold;line-height: 24px;padding: 12px 24px;text-align: center;text-decoration: none !important;transition: opacity 0.1s ease-in;color: #001b71 !important;background-color: #d5ff05;font-family: Roboto, Tahoma, sans-serif;\"\n" +
	"                                            href=\"https://www.bilibili.com/video/av%d\">Watch now &gt;</a><![endif]>\n" +
	"                                        <!--[if mso]><p style=\"line-height:0;margin:0;\">&nbsp;</p>\n" +
	"                                    <v:roundrect xmlns:v=\"urn:schemas-microsoft-com:vml\" href=\"https://www.bilibili.com/video/av%d\"\n" +
	"                                                 style=\"width:121px\" arcsize=\"9%%\" fillcolor=\"#D5FF05\" stroke=\"f\">\n" +
	"                                        <v:textbox style=\"mso-fit-shape-to-text:t\" inset=\"0px,11px,0px,11px\">\n" +
	"                                            <center style=\"font-size:14px;line-height:24px;color:#001B71;font-family:Roboto,Tahoma,sans-serif;font-weight:bold;mso-line-height-rule:exactly;mso-text-raise:4px\">\n" +
	"                                                Watch now &gt;\n" +
	"                                            </center>\n" +
	"                                        </v:textbox>\n" +
	"                                    </v:roundrect><![endif]--></div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 35px;font-size: 1px;\">\n" +
	"                                        &nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td>\n" +
	"                        <td class=\"layout__edges\">&nbsp;</td></tr></table><![endif]-->\n" +
	"                        </div>\n" +
	"                    </div>\n" +
	"                </div>"

var htmlHeader = "<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional //EN\"\n" +
	"        \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\"><!--[if IE]>\n" +
	"<html xmlns=\"http://www.w3.org/1999/xhtml\" class=\"ie\"><![endif]--><!--[if !IE]><!-->\n" +
	"<html style=\"margin: 0;padding: 0;\" xmlns=\"http://www.w3.org/1999/xhtml\"><!--<![endif]-->\n" +
	"<head>\n" +
	"    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/>\n" +
	"    <title></title>\n" +
	"    <!--[if !mso]><!-->\n" +
	"    <meta http-equiv=\"X-UA-Compatible\" content=\"IE=edge\"/><!--<![endif]-->\n" +
	"    <meta name=\"viewport\" content=\"width=device-width\"/>\n" +
	"    <style type=\"text/css\">\n" +
	"        @media only screen and (min-width: 620px) {\n" +
	"            .wrapper {\n" +
	"                min-width: 600px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h1 {\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h1 {\n" +
	"                font-size: 28px !important;\n" +
	"                line-height: 36px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h2 {\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h2 {\n" +
	"                font-size: 22px !important;\n" +
	"                line-height: 31px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h3 {\n" +
	"            }\n" +
	"\n" +
	"            .wrapper h3 {\n" +
	"                font-size: 18px !important;\n" +
	"                line-height: 26px !important\n" +
	"            }\n" +
	"\n" +
	"            .column {\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-8 {\n" +
	"                font-size: 8px !important;\n" +
	"                line-height: 14px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-9 {\n" +
	"                font-size: 9px !important;\n" +
	"                line-height: 16px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-10 {\n" +
	"                font-size: 10px !important;\n" +
	"                line-height: 18px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-11 {\n" +
	"                font-size: 11px !important;\n" +
	"                line-height: 19px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-12 {\n" +
	"                font-size: 12px !important;\n" +
	"                line-height: 19px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-13 {\n" +
	"                font-size: 13px !important;\n" +
	"                line-height: 21px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-14 {\n" +
	"                font-size: 14px !important;\n" +
	"                line-height: 21px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-15 {\n" +
	"                font-size: 15px !important;\n" +
	"                line-height: 23px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-16 {\n" +
	"                font-size: 16px !important;\n" +
	"                line-height: 24px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-17 {\n" +
	"                font-size: 17px !important;\n" +
	"                line-height: 26px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-18 {\n" +
	"                font-size: 18px !important;\n" +
	"                line-height: 26px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-20 {\n" +
	"                font-size: 20px !important;\n" +
	"                line-height: 28px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-22 {\n" +
	"                font-size: 22px !important;\n" +
	"                line-height: 31px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-24 {\n" +
	"                font-size: 24px !important;\n" +
	"                line-height: 32px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-26 {\n" +
	"                font-size: 26px !important;\n" +
	"                line-height: 34px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-28 {\n" +
	"                font-size: 28px !important;\n" +
	"                line-height: 36px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-30 {\n" +
	"                font-size: 30px !important;\n" +
	"                line-height: 38px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-32 {\n" +
	"                font-size: 32px !important;\n" +
	"                line-height: 40px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-34 {\n" +
	"                font-size: 34px !important;\n" +
	"                line-height: 43px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-36 {\n" +
	"                font-size: 36px !important;\n" +
	"                line-height: 43px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper\n" +
	"            .size-40 {\n" +
	"                font-size: 40px !important;\n" +
	"                line-height: 47px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-44 {\n" +
	"                font-size: 44px !important;\n" +
	"                line-height: 50px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-48 {\n" +
	"                font-size: 48px !important;\n" +
	"                line-height: 54px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-56 {\n" +
	"                font-size: 56px !important;\n" +
	"                line-height: 60px !important\n" +
	"            }\n" +
	"\n" +
	"            .wrapper .size-64 {\n" +
	"                font-size: 64px !important;\n" +
	"                line-height: 63px !important\n" +
	"            }\n" +
	"        }\n" +
	"    </style>\n" +
	"    <style type=\"text/css\">\n" +
	"        body {\n" +
	"            margin: 0;\n" +
	"            padding: 0;\n" +
	"        }\n" +
	"\n" +
	"        table {\n" +
	"            border-collapse: collapse;\n" +
	"            table-layout: fixed;\n" +
	"        }\n" +
	"\n" +
	"        * {\n" +
	"            line-height: inherit;\n" +
	"        }\n" +
	"\n" +
	"        [x-apple-data-detectors],\n" +
	"        [href^=\"tel\"],\n" +
	"        [href^=\"sms\"] {\n" +
	"            color: inherit !important;\n" +
	"            text-decoration: none !important;\n" +
	"        }\n" +
	"\n" +
	"        .wrapper .footer__share-button a:hover,\n" +
	"        .wrapper .footer__share-button a:focus {\n" +
	"            color: #ffffff !important;\n" +
	"        }\n" +
	"\n" +
	"        .btn a:hover,\n" +
	"        .btn a:focus,\n" +
	"        .footer__share-button a:hover,\n" +
	"        .footer__share-button a:focus,\n" +
	"        .email-footer__links a:hover,\n" +
	"        .email-footer__links a:focus {\n" +
	"            opacity: 0.8;\n" +
	"        }\n" +
	"\n" +
	"        .preheader,\n" +
	"        .header,\n" +
	"        .layout,\n" +
	"        .column {\n" +
	"            transition: width 0.25s ease-in-out, max-width 0.25s ease-in-out;\n" +
	"        }\n" +
	"\n" +
	"        .preheader td {\n" +
	"            padding-bottom: 8px;\n" +
	"        }\n" +
	"\n" +
	"        .layout,\n" +
	"        div.header {\n" +
	"            max-width: 400px !important;\n" +
	"            -fallback-width: 95% !important;\n" +
	"            width: calc(100% - 20px) !important;\n" +
	"        }\n" +
	"\n" +
	"        div.preheader {\n" +
	"            max-width: 360px !important;\n" +
	"            -fallback-width: 90% !important;\n" +
	"            width: calc(100% - 60px) !important;\n" +
	"        }\n" +
	"\n" +
	"        .snippet,\n" +
	"        .webversion {\n" +
	"            Float: none !important;\n" +
	"        }\n" +
	"\n" +
	"        .column {\n" +
	"            max-width: 400px !important;\n" +
	"            width: 100% !important;\n" +
	"        }\n" +
	"\n" +
	"        .fixed-width.has-border {\n" +
	"            max-width: 402px !important;\n" +
	"        }\n" +
	"\n" +
	"        .fixed-width.has-border .layout__inner {\n" +
	"            box-sizing: border-box;\n" +
	"        }\n" +
	"\n" +
	"        .snippet,\n" +
	"        .webversion {\n" +
	"            width: 50% !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .btn {\n" +
	"            width: 100%;\n" +
	"        }\n" +
	"\n" +
	"        [owa] .column div,\n" +
	"        [owa] .column button {\n" +
	"            display: block !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .column,\n" +
	"        [owa] .column,\n" +
	"        .ie .gutter,\n" +
	"        [owa] .gutter {\n" +
	"            display: table-cell;\n" +
	"            float: none !important;\n" +
	"            vertical-align: top;\n" +
	"        }\n" +
	"\n" +
	"        .ie div.preheader,\n" +
	"        [owa] div.preheader,\n" +
	"        .ie .email-footer,\n" +
	"        [owa] .email-footer {\n" +
	"            max-width: 560px !important;\n" +
	"            width: 560px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .snippet,\n" +
	"        [owa] .snippet,\n" +
	"        .ie .webversion,\n" +
	"        [owa] .webversion {\n" +
	"            width: 280px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie div.header,\n" +
	"        [owa] div.header,\n" +
	"        .ie .layout,\n" +
	"        [owa] .layout,\n" +
	"        .ie .one-col .column,\n" +
	"        [owa] .one-col .column {\n" +
	"            max-width: 600px !important;\n" +
	"            width: 600px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .fixed-width.has-border,\n" +
	"        [owa] .fixed-width.has-border,\n" +
	"        .ie .has-gutter.has-border,\n" +
	"        [owa] .has-gutter.has-border {\n" +
	"            max-width: 602px !important;\n" +
	"            width: 602px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .two-col .column,\n" +
	"        [owa] .two-col .column {\n" +
	"            max-width: 300px !important;\n" +
	"            width: 300px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .three-col .column,\n" +
	"        [owa] .three-col .column,\n" +
	"        .ie .narrow,\n" +
	"        [owa] .narrow {\n" +
	"            max-width: 200px !important;\n" +
	"            width: 200px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .wide,\n" +
	"        [owa] .wide {\n" +
	"            width: 400px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .two-col.has-gutter .column,\n" +
	"        [owa] .two-col.x_has-gutter .column {\n" +
	"            max-width: 290px !important;\n" +
	"            width: 290px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .three-col.has-gutter .column,\n" +
	"        [owa] .three-col.x_has-gutter .column,\n" +
	"        .ie .has-gutter .narrow,\n" +
	"        [owa] .has-gutter .narrow {\n" +
	"            max-width: 188px !important;\n" +
	"            width: 188px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .has-gutter .wide,\n" +
	"        [owa] .has-gutter .wide {\n" +
	"            max-width: 394px !important;\n" +
	"            width: 394px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .two-col.has-gutter.has-border .column,\n" +
	"        [owa] .two-col.x_has-gutter.x_has-border .column {\n" +
	"            max-width: 292px !important;\n" +
	"            width: 292px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .three-col.has-gutter.has-border .column,\n" +
	"        [owa] .three-col.x_has-gutter.x_has-border .column,\n" +
	"        .ie .has-gutter.has-border .narrow,\n" +
	"        [owa] .has-gutter.x_has-border .narrow {\n" +
	"            max-width: 190px !important;\n" +
	"            width: 190px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .has-gutter.has-border .wide,\n" +
	"        [owa] .has-gutter.x_has-border .wide {\n" +
	"            max-width: 396px !important;\n" +
	"            width: 396px !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .fixed-width .layout__inner {\n" +
	"            border-left: 0 none white !important;\n" +
	"            border-right: 0 none white !important;\n" +
	"        }\n" +
	"\n" +
	"        .ie .layout__edges {\n" +
	"            display: none;\n" +
	"        }\n" +
	"\n" +
	"        .mso .layout__edges {\n" +
	"            font-size: 0;\n" +
	"        }\n" +
	"\n" +
	"        .layout-fixed-width,\n" +
	"        .mso .layout-full-width {\n" +
	"            background-color: #ffffff;\n" +
	"        }\n" +
	"\n" +
	"        @media only screen and (min-width: 620px) {\n" +
	"            .column,\n" +
	"            .gutter {\n" +
	"                display: table-cell;\n" +
	"                Float: none !important;\n" +
	"                vertical-align: top;\n" +
	"            }\n" +
	"\n" +
	"            div.preheader,\n" +
	"            .email-footer {\n" +
	"                max-width: 560px !important;\n" +
	"                width: 560px !important;\n" +
	"            }\n" +
	"\n" +
	"            .snippet,\n" +
	"            .webversion {\n" +
	"                width: 280px !important;\n" +
	"            }\n" +
	"\n" +
	"            div.header,\n" +
	"            .layout,\n" +
	"            .one-col .column {\n" +
	"                max-width: 600px !important;\n" +
	"                width: 600px !important;\n" +
	"            }\n" +
	"\n" +
	"            .fixed-width.has-border,\n" +
	"            .fixed-width.ecxhas-border,\n" +
	"            .has-gutter.has-border,\n" +
	"            .has-gutter.ecxhas-border {\n" +
	"                max-width: 602px !important;\n" +
	"                width: 602px !important;\n" +
	"            }\n" +
	"\n" +
	"            .two-col .column {\n" +
	"                max-width: 300px !important;\n" +
	"                width: 300px !important;\n" +
	"            }\n" +
	"\n" +
	"            .three-col .column,\n" +
	"            .column.narrow {\n" +
	"                max-width: 200px !important;\n" +
	"                width: 200px !important;\n" +
	"            }\n" +
	"\n" +
	"            .column.wide {\n" +
	"                width: 400px !important;\n" +
	"            }\n" +
	"\n" +
	"            .two-col.has-gutter .column,\n" +
	"            .two-col.ecxhas-gutter .column {\n" +
	"                max-width: 290px !important;\n" +
	"                width: 290px !important;\n" +
	"            }\n" +
	"\n" +
	"            .three-col.has-gutter .column,\n" +
	"            .three-col.ecxhas-gutter .column,\n" +
	"            .has-gutter .narrow {\n" +
	"                max-width: 188px !important;\n" +
	"                width: 188px !important;\n" +
	"            }\n" +
	"\n" +
	"            .has-gutter .wide {\n" +
	"                max-width: 394px !important;\n" +
	"                width: 394px !important;\n" +
	"            }\n" +
	"\n" +
	"            .two-col.has-gutter.has-border .column,\n" +
	"            .two-col.ecxhas-gutter.ecxhas-border .column {\n" +
	"                max-width: 292px !important;\n" +
	"                width: 292px !important;\n" +
	"            }\n" +
	"\n" +
	"            .three-col.has-gutter.has-border .column,\n" +
	"            .three-col.ecxhas-gutter.ecxhas-border .column,\n" +
	"            .has-gutter.has-border .narrow,\n" +
	"            .has-gutter.ecxhas-border .narrow {\n" +
	"                max-width: 190px !important;\n" +
	"                width: 190px !important;\n" +
	"            }\n" +
	"\n" +
	"            .has-gutter.has-border .wide,\n" +
	"            .has-gutter.ecxhas-border .wide {\n" +
	"                max-width: 396px !important;\n" +
	"                width: 396px !important;\n" +
	"            }\n" +
	"        }\n" +
	"\n" +
	"        @media only screen and (-webkit-min-device-pixel-ratio: 2), only screen and (min--moz-device-pixel-ratio: 2), only screen and (-o-min-device-pixel-ratio: 2/1), only screen and (min-device-pixel-ratio: 2), only screen and (min-resolution: 192dpi), only screen and (min-resolution: 2dppx) {\n" +
	"            .fblike {\n" +
	"                background-image: url(https://i7.createsend1.com/static/eb/master/13-the-blueprint-3/images/fblike@2x.png) !important;\n" +
	"            }\n" +
	"\n" +
	"            .tweet {\n" +
	"                background-image: url(https://i8.createsend1.com/static/eb/master/13-the-blueprint-3/images/tweet@2x.png) !important;\n" +
	"            }\n" +
	"\n" +
	"            .linkedinshare {\n" +
	"                background-image: url(https://i10.createsend1.com/static/eb/master/13-the-blueprint-3/images/lishare@2x.png) !important;\n" +
	"            }\n" +
	"\n" +
	"            .forwardtoafriend {\n" +
	"                background-image: url(https://i9.createsend1.com/static/eb/master/13-the-blueprint-3/images/forward@2x.png) !important;\n" +
	"            }\n" +
	"        }\n" +
	"\n" +
	"        @media (max-width: 321px) {\n" +
	"            .fixed-width.has-border .layout__inner {\n" +
	"                border-width: 1px 0 !important;\n" +
	"            }\n" +
	"\n" +
	"            .layout,\n" +
	"            .column {\n" +
	"                min-width: 320px !important;\n" +
	"                width: 320px !important;\n" +
	"            }\n" +
	"\n" +
	"            .border {\n" +
	"                display: none;\n" +
	"            }\n" +
	"        }\n" +
	"\n" +
	"        .mso div {\n" +
	"            border: 0 none white !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w560 .divider {\n" +
	"            Margin-left: 260px !important;\n" +
	"            Margin-right: 260px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w360 .divider {\n" +
	"            Margin-left: 160px !important;\n" +
	"            Margin-right: 160px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w260 .divider {\n" +
	"            Margin-left: 110px !important;\n" +
	"            Margin-right: 110px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w160 .divider {\n" +
	"            Margin-left: 60px !important;\n" +
	"            Margin-right: 60px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w354 .divider {\n" +
	"            Margin-left: 157px !important;\n" +
	"            Margin-right: 157px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w250 .divider {\n" +
	"            Margin-left: 105px !important;\n" +
	"            Margin-right: 105px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .w148 .divider {\n" +
	"            Margin-left: 54px !important;\n" +
	"            Margin-right: 54px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-8,\n" +
	"        .ie .size-8 {\n" +
	"            font-size: 8px !important;\n" +
	"            line-height: 14px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-9,\n" +
	"        .ie .size-9 {\n" +
	"            font-size: 9px !important;\n" +
	"            line-height: 16px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-10,\n" +
	"        .ie .size-10 {\n" +
	"            font-size: 10px !important;\n" +
	"            line-height: 18px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-11,\n" +
	"        .ie .size-11 {\n" +
	"            font-size: 11px !important;\n" +
	"            line-height: 19px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-12,\n" +
	"        .ie .size-12 {\n" +
	"            font-size: 12px !important;\n" +
	"            line-height: 19px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-13,\n" +
	"        .ie .size-13 {\n" +
	"            font-size: 13px !important;\n" +
	"            line-height: 21px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-14,\n" +
	"        .ie .size-14 {\n" +
	"            font-size: 14px !important;\n" +
	"            line-height: 21px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-15,\n" +
	"        .ie .size-15 {\n" +
	"            font-size: 15px !important;\n" +
	"            line-height: 23px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-16,\n" +
	"        .ie .size-16 {\n" +
	"            font-size: 16px !important;\n" +
	"            line-height: 24px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-17,\n" +
	"        .ie .size-17 {\n" +
	"            font-size: 17px !important;\n" +
	"            line-height: 26px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-18,\n" +
	"        .ie .size-18 {\n" +
	"            font-size: 18px !important;\n" +
	"            line-height: 26px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-20,\n" +
	"        .ie .size-20 {\n" +
	"            font-size: 20px !important;\n" +
	"            line-height: 28px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-22,\n" +
	"        .ie .size-22 {\n" +
	"            font-size: 22px !important;\n" +
	"            line-height: 31px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-24,\n" +
	"        .ie .size-24 {\n" +
	"            font-size: 24px !important;\n" +
	"            line-height: 32px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-26,\n" +
	"        .ie .size-26 {\n" +
	"            font-size: 26px !important;\n" +
	"            line-height: 34px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-28,\n" +
	"        .ie .size-28 {\n" +
	"            font-size: 28px !important;\n" +
	"            line-height: 36px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-30,\n" +
	"        .ie .size-30 {\n" +
	"            font-size: 30px !important;\n" +
	"            line-height: 38px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-32,\n" +
	"        .ie .size-32 {\n" +
	"            font-size: 32px !important;\n" +
	"            line-height: 40px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-34,\n" +
	"        .ie .size-34 {\n" +
	"            font-size: 34px !important;\n" +
	"            line-height: 43px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-36,\n" +
	"        .ie .size-36 {\n" +
	"            font-size: 36px !important;\n" +
	"            line-height: 43px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-40,\n" +
	"        .ie .size-40 {\n" +
	"            font-size: 40px !important;\n" +
	"            line-height: 47px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-44,\n" +
	"        .ie .size-44 {\n" +
	"            font-size: 44px !important;\n" +
	"            line-height: 50px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-48,\n" +
	"        .ie .size-48 {\n" +
	"            font-size: 48px !important;\n" +
	"            line-height: 54px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-56,\n" +
	"        .ie .size-56 {\n" +
	"            font-size: 56px !important;\n" +
	"            line-height: 60px !important;\n" +
	"        }\n" +
	"\n" +
	"        .mso .size-64,\n" +
	"        .ie .size-64 {\n" +
	"            font-size: 64px !important;\n" +
	"            line-height: 63px !important;\n" +
	"        }\n" +
	"    </style>\n" +
	"\n" +
	"    <!--[if !mso]><!-->\n" +
	"    <style type=\"text/css\">\n" +
	"        @import url(https://fonts.googleapis.com/css?family=Roboto:400,700,400italic,700italic);\n" +
	"    </style>\n" +
	"    <link href=\"https://fonts.googleapis.com/css?family=Roboto:400,700,400italic,700italic\" rel=\"stylesheet\"\n" +
	"          type=\"text/css\"/><!--<![endif]-->\n" +
	"    <style type=\"text/css\">\n" +
	"        body {\n" +
	"            background-color: #009ff7\n" +
	"        }\n" +
	"\n" +
	"        .logo a:hover, .logo a:focus {\n" +
	"            color: #fff !important\n" +
	"        }\n" +
	"\n" +
	"        .mso .layout-has-border {\n" +
	"            border-top: 1px solid #005d91;\n" +
	"            border-bottom: 1px solid #005d91\n" +
	"        }\n" +
	"\n" +
	"        .mso .layout-has-bottom-border {\n" +
	"            border-bottom: 1px solid #005d91\n" +
	"        }\n" +
	"\n" +
	"        .mso .border, .ie .border {\n" +
	"            background-color: #005d91\n" +
	"        }\n" +
	"\n" +
	"        .mso h1, .ie h1 {\n" +
	"        }\n" +
	"\n" +
	"        .mso h1, .ie h1 {\n" +
	"            font-size: 28px !important;\n" +
	"            line-height: 36px !important\n" +
	"        }\n" +
	"\n" +
	"        .mso h2, .ie h2 {\n" +
	"        }\n" +
	"\n" +
	"        .mso h2, .ie h2 {\n" +
	"            font-size: 22px !important;\n" +
	"            line-height: 31px !important\n" +
	"        }\n" +
	"\n" +
	"        .mso h3, .ie h3 {\n" +
	"        }\n" +
	"\n" +
	"        .mso h3, .ie h3 {\n" +
	"            font-size: 18px !important;\n" +
	"            line-height: 26px !important\n" +
	"        }\n" +
	"\n" +
	"        .mso .layout__inner, .ie .layout__inner {\n" +
	"        }\n" +
	"\n" +
	"        .mso .footer__share-button p {\n" +
	"        }\n" +
	"\n" +
	"        .mso .footer__share-button p {\n" +
	"            font-family: sans-serif\n" +
	"        }\n" +
	"    </style>\n" +
	"    <meta name=\"robots\" content=\"noindex,nofollow\"/>\n" +
	"    <meta property=\"og:title\" content=\"My First Campaign\"/>\n" +
	"</head>\n" +
	"<body class=\"no-padding\" style=\"margin: 0;padding: 0;-webkit-text-size-adjust: 100%;\">\n" +
	"<!--<![endif]-->\n" +
	"<table class=\"wrapper\"\n" +
	"       style=\"border-collapse: collapse;table-layout: fixed;min-width: 320px;width: 100%;background-color: #009ff7;\"\n" +
	"       cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"    <tbody>\n" +
	"    <tr>\n" +
	"        <td>\n" +
	"            <div role=\"banner\">\n" +
	"                <div class=\"preheader\"\n" +
	"                     style=\"Margin: 0 auto;max-width: 560px;min-width: 280px; width: 280px;width: calc(28000% - 167440px);\">\n" +
	"                    <div style=\"border-collapse: collapse;display: table;width: 100%;\">\n" +
	"                        <!--[if (mso)|(IE)]>\n" +
	"                        <table align=\"center\" class=\"preheader\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                            <tr>\n" +
	"                                <td style=\"width: 280px\" valign=\"top\"><![endif]-->\n" +
	"                        <div class=\"snippet\"\n" +
	"                             style=\"display: table-cell;Float: left;font-size: 12px;line-height: 19px;max-width: 280px;min-width: 140px; width: 140px;width: calc(14000% - 78120px);padding: 10px 0 5px 0;color: #fff;font-family: sans-serif;\">\n" +
	"\n" +
	"                        </div>\n" +
	"                        <!--[if (mso)|(IE)]></td>\n" +
	"                    <td style=\"width: 280px\" valign=\"top\"><![endif]-->\n" +
	"                        <div class=\"webversion\"\n" +
	"                             style=\"display: table-cell;Float: left;font-size: 12px;line-height: 19px;max-width: 280px;min-width: 139px; width: 139px;width: calc(14100% - 78680px);padding: 10px 0 5px 0;text-align: right;color: #fff;font-family: sans-serif;\">\n" +
	"                            <p style=\"Margin-top: 0;Margin-bottom: 0;\">No Images?\n" +
	"                                <webversion style=\"text-decoration: underline;\">Click here</webversion>\n" +
	"                            </p>\n" +
	"                        </div>\n" +
	"                        <!--[if (mso)|(IE)]></td></tr></table><![endif]-->\n" +
	"                    </div>\n" +
	"                </div>\n" +
	"                <div class=\"header\"\n" +
	"                     style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);\"\n" +
	"                     id=\"emb-email-header-container\">\n" +
	"                    <!--[if (mso)|(IE)]>\n" +
	"                    <table align=\"center\" class=\"header\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                        <tr>\n" +
	"                            <td style=\"width: 600px\"><![endif]-->\n" +
	"                    <div class=\"logo emb-logo-margin-box\"\n" +
	"                         style=\"font-size: 26px;line-height: 32px;Margin-top: 0px;Margin-bottom: 46px;color: #c3ced9;font-family: Roboto,Tahoma,sans-serif;Margin-left: 20px;Margin-right: 20px;\"\n" +
	"                         align=\"center\">\n" +
	"                        <div class=\"logo-center\" align=\"center\" id=\"emb-email-header\"><img\n" +
	"                                style=\"display: block;height: auto;width: 100%;border: 0;max-width: 390px;\"\n" +
	"                                src=\"https://i.loli.net/2019/03/17/5c8e447949b55.jpg\" alt=\"CCM is CCM Cloud Messaging\" width=\"390\"/></div>\n" +
	"                    </div>\n" +
	"                    <!--[if (mso)|(IE)]></td></tr></table><![endif]-->\n" +
	"                </div>\n" +
	"            </div>\n" +
	"            <div>\n" +
	"                <div style=\"background-color: #ffffff;\">\n" +
	"                    <div class=\"layout one-col\"\n" +
	"                         style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;\">\n" +
	"                        <div class=\"layout__inner\" style=\"border-collapse: collapse;display: table;width: 100%;\">\n" +
	"                            <!--[if (mso)|(IE)]>\n" +
	"                            <table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                                <tr class=\"layout-full-width\" style=\"background-color: #ffffff;\">\n" +
	"                                    <td class=\"layout__edges\">&nbsp;</td>\n" +
	"                                    <td style=\"width: 600px\" class=\"w560\"><![endif]-->\n" +
	"                            <div class=\"column\"\n" +
	"                                 style=\"max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);text-align: left;color: #f6f6f8;font-size: 16px;line-height: 24px;font-family: Roboto,Tahoma,sans-serif;\">\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 50px;font-size: 1px;\">\n" +
	"                                        &nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td>\n" +
	"                        <td class=\"layout__edges\">&nbsp;</td></tr></table><![endif]-->\n" +
	"                        </div>\n" +
	"                    </div>\n" +
	"                </div>"

var htmlFooter = "<div style=\"background-color: #009ff7;background-position: 0px 0px;background-image: url(https://i1.createsend1.com/ei/t/20/316/BD9/235535/csfinal/promo-bg1.png);background-repeat: repeat;\">\n" +
	"                    <div class=\"layout one-col\"\n" +
	"                         style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;\">\n" +
	"                        <div class=\"layout__inner\" style=\"border-collapse: collapse;display: table;width: 100%;\">\n" +
	"                            <!--[if (mso)|(IE)]>\n" +
	"                            <table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                                <tr class=\"layout-full-width\"\n" +
	"                                    style=\"background-color: #009ff7;background-position: 0px 0px;background-image: url(https://i1.createsend1.com/ei/t/20/316/BD9/235535/csfinal/promo-bg1.png);background-repeat: repeat;\">\n" +
	"                                    <td class=\"layout__edges\">&nbsp;</td>\n" +
	"                                    <td style=\"width: 600px\" class=\"w560\"><![endif]-->\n" +
	"                            <div class=\"column\"\n" +
	"                                 style=\"max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);text-align: left;color: #f6f6f8;font-size: 16px;line-height: 24px;font-family: Roboto,Tahoma,sans-serif;\">\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 35px;font-size: 1px;\">\n" +
	"                                        &nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 1px;font-size: 1px;\">&nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;mso-text-raise: 4px;\">\n" +
	"                                        <h1 class=\"size-34\"\n" +
	"                                            style=\"Margin-top: 0;Margin-bottom: 20px;font-style: normal;font-weight: normal;color: #fff;font-size: 30px;line-height: 38px;text-align: center;\"\n" +
	"                                            lang=\"x-size-34\">Make Android Better In China</h1>\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;\">\n" +
	"                                    <div style=\"mso-line-height-rule: exactly;line-height: 1px;font-size: 1px;\">&nbsp;\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td>\n" +
	"                        <td class=\"layout__edges\">&nbsp;</td></tr></table><![endif]-->\n" +
	"                        </div>\n" +
	"                    </div>\n" +
	"                </div>\n" +
	"\n" +
	"\n" +
	"                <div role=\"contentinfo\">\n" +
	"                    <div class=\"layout email-footer\"\n" +
	"                         style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;\">\n" +
	"                        <div class=\"layout__inner\" style=\"border-collapse: collapse;display: table;width: 100%;\">\n" +
	"                            <!--[if (mso)|(IE)]>\n" +
	"                            <table align=\"center\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                                <tr class=\"layout-email-footer\">\n" +
	"                                    <td style=\"width: 400px;\" valign=\"top\" class=\"w360\"><![endif]-->\n" +
	"                            <div class=\"column wide\"\n" +
	"                                 style=\"text-align: left;font-size: 12px;line-height: 19px;color: #fff;font-family: sans-serif;Float: left;max-width: 400px;min-width: 320px; width: 320px;width: calc(8000% - 47600px);\">\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;Margin-top: 10px;Margin-bottom: 10px;\">\n" +
	"                                    <table class=\"email-footer__links emb-web-links\"\n" +
	"                                           style=\"border-collapse: collapse;table-layout: fixed;\" role=\"presentation\">\n" +
	"                                        <tbody>\n" +
	"                                        <tr role=\"navigation\">\n" +
	"                                            <td class=\"emb-web-links\" style=\"padding: 0;width: 26px;\"><a\n" +
	"                                                    style=\"text-decoration: underline;transition: opacity 0.1s ease-in;color: #fff;\"\n" +
	"                                                    href=\"https://github.com/chen622/Spider\"><img style=\"border: 0;\"\n" +
	"                                                                                                  src=\"https://i7.createsend1.com/static/eb/master/13-the-blueprint-3/images/website.png\"\n" +
	"                                                                                                  width=\"26\" height=\"26\"\n" +
	"                                                                                                  alt=\"Website\"/></a>\n" +
	"                                            </td>\n" +
	"                                        </tr>\n" +
	"                                        </tbody>\n" +
	"                                    </table>\n" +
	"                                    <div style=\"font-size: 12px;line-height: 19px;Margin-top: 20px;\">\n" +
	"                                        <div>CCM is CCM Cloud Messaging</div>\n" +
	"                                    </div>\n" +
	"                                    <div style=\"font-size: 12px;line-height: 19px;Margin-top: 18px;\">\n" +
	"\n" +
	"                                    </div>\n" +
	"                                    <!--[if mso]>&nbsp;<![endif]-->\n" +
	"                                </div>\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td>\n" +
	"                        <td style=\"width: 200px;\" valign=\"top\" class=\"w160\"><![endif]-->\n" +
	"                            <div class=\"column narrow\"\n" +
	"                                 style=\"text-align: left;font-size: 12px;line-height: 19px;color: #fff;font-family: sans-serif;Float: left;max-width: 320px;min-width: 200px; width: 320px;width: calc(72200px - 12000%);\">\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;Margin-top: 10px;Margin-bottom: 10px;\">\n" +
	"\n" +
	"                                </div>\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td></tr></table><![endif]-->\n" +
	"                        </div>\n" +
	"                    </div>\n" +
	"                    <div class=\"layout one-col email-footer\"\n" +
	"                         style=\"Margin: 0 auto;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);overflow-wrap: break-word;word-wrap: break-word;word-break: break-word;\">\n" +
	"                        <div class=\"layout__inner\" style=\"border-collapse: collapse;display: table;width: 100%;\">\n" +
	"                            <!--[if (mso)|(IE)]>\n" +
	"                            <table align=\"center\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\">\n" +
	"                                <tr class=\"layout-email-footer\">\n" +
	"                                    <td style=\"width: 600px;\" class=\"w560\"><![endif]-->\n" +
	"                            <div class=\"column\"\n" +
	"                                 style=\"text-align: left;font-size: 12px;line-height: 19px;color: #fff;font-family: sans-serif;max-width: 600px;min-width: 320px; width: 320px;width: calc(28000% - 167400px);\">\n" +
	"                                <div style=\"Margin-left: 20px;Margin-right: 20px;Margin-top: 10px;Margin-bottom: 10px;\">\n" +
	"                                    <div style=\"font-size: 12px;line-height: 19px;\">\n" +
	"                                        <unsubscribe style=\"text-decoration: underline;\">Unsubscribe</unsubscribe>\n" +
	"                                    </div>\n" +
	"                                </div>\n" +
	"                            </div>\n" +
	"                            <!--[if (mso)|(IE)]></td></tr></table><![endif]-->\n" +
	"                        </div>\n" +
	"                    </div>\n" +
	"                </div>\n" +
	"                <div style=\"line-height:40px;font-size:40px;\">&nbsp;</div>\n" +
	"            </div>\n" +
	"        </td>\n" +
	"    </tr>\n" +
	"    </tbody>\n" +
	"</table>\n" +
	"\n" +
	"</body>\n" +
	"</html>\n"
