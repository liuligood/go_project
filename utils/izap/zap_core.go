package izap

import (
	"crmeb_go/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
	zap *config.Zap
}

func NewZapCore(za *config.Zap, level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level, zap: za}

	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})

	entity.Core = zapcore.NewCore(za.Encoder(), syncer, levelEnabler)

	return entity
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := NewCutter(
		z.zap.Director,
		z.level.String(),
		z.zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)

	if z.zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}

	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(z.zap.Encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
