package handler

import (
	"context"

	"github.com/sirupsen/logrus"
)

func log(ctx context.Context) logrus.FieldLogger {
	return logrus.WithContext(ctx).WithField("pkg", "handler")
}
