package base64_captcha

import (
	"context"
	"crmeb_go/utils/izap"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"image/color"
)

/*
使用内存方式
// GCLimitNumber 生成的验证码数量，用于触发默认存储空间的垃圾回收
GCLimitNumber = 10240
// Expiration 验证码过期时间
Expiration = 10 * time.Minute
*/

func GetCaptcha(RedisClient redis.UniversalClient, ctx context.Context) (error, string, string) {
	var store = NewRedisStore(RedisClient, ctx)

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
	driver := driverStringConfig.ConvertFonts()

	// 实例化一个captcha结构体
	cap := base64Captcha.NewCaptcha(driver, store)
	// 生成id,content,answer
	id, content, answer := cap.Driver.GenerateIdQuestionAnswer()

	// 生成图片
	item, err := cap.Driver.DrawCaptcha(content)
	if err != nil {
		izap.Log.Error(" cap.Driver.DrawCaptcha:", zap.Error(err))

		return err, "", ""
	}

	err = store.Set(id, answer)
	if err != nil {
		izap.Log.Error("设置验证码密码错误:", zap.String("id为", id), zap.Error(err))

		return err, "", ""
	}

	// 转为base64编码string格式
	return nil, item.EncodeB64string(), id
}
