package captcha

import (
	"context"
	"crmeb_go/utils/izap"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"image/color"
	"strings"
)

type CaptchaClient struct {
	captcha *base64Captcha.Captcha
}

func NewCaptchaClient(redisClient redis.UniversalClient, ctx context.Context) *CaptchaClient {
	driverStringConfig := &base64Captcha.DriverString{
		// 验证码图片的高度
		Height: 50,
		// 验证码图片的宽度
		Width: 120,
		// 验证码图片中随机噪点的数量
		NoiseCount: 2,
		// 控制显示在验证码图片中的线条的选项
		ShowLineOptions: 2 | 4,
		// 验证码的长度，即验证码中字符的数量
		Length: 4,
		// 验证码的字符源，用于生成验证码的字符。在这个例子中，使用数字和小写字母作为字符源
		Source: "1234567890abcdefghijklmnopqrstuvwxyz",
		BgColor: &color.RGBA{
			//验证码图片的背景颜色
			R: 100,
			G: 200,
			B: 100,
			A: 125,
		},
		//用于绘制验证码文本的字体文件。使用"wqy-microhei.ttc"的字体文件
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 将driverString中指定的字体文件转换为驱动程序所需的字体格式,这个步骤是为了将字体文件转换为正确的格式，以便在生成验证码时使用正确的字体
	return &CaptchaClient{
		captcha: base64Captcha.NewCaptcha(driverStringConfig.ConvertFonts(), NewRedisStore(redisClient, ctx)),
	}
}

func (c *CaptchaClient) GetCaptcha() (error, string, string) {
	id, b64s, _, err := c.captcha.Generate()
	if err != nil {
		izap.Log.Error("c.captcha.Generate:", zap.Error(err))

		return err, "", ""
	}

	return nil, b64s, id
}

func (c *CaptchaClient) Verify(id, answer string) bool {
	answer = strings.ToLower(answer)
	return c.captcha.Verify(id, answer, true)
}
