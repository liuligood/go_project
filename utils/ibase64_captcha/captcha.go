package ibase64_captcha

import (
	"context"
	"crmeb_go/utils/izap"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strings"
)

type CaptchaClient struct {
	redisStore *RedisStore
	captcha    *base64Captcha.Captcha
}

func NewCaptchaClient(redisClient redis.UniversalClient, ctx context.Context) CaptchaClient {

	return CaptchaClient{
		captcha: base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, NewRedisStore(redisClient, ctx)),
	}
}

func (c CaptchaClient) GetCaptcha() (error, string, string) {
	id, b64s, _, err := c.captcha.Generate()
	if err != nil {
		izap.Log.Error("c.captcha.Generate:", zap.Error(err))

		return err, "", ""
	}

	return nil, b64s, id
}

func (c CaptchaClient) Verify(id, answer string) bool {
	answer = strings.ToLower(answer)
	return c.captcha.Verify(id, answer, true)
}
