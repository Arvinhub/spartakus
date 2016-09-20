package collector

import "github.com/coreos/pkg/capnslog"

var log = capnslog.NewPackageLogger("k8s.io/spartakus", "collector")

type logWriter struct {
	log   *capnslog.PackageLogger
	level capnslog.LogLevel
}

func (l *logWriter) Write(d []byte) (int, error) {
	l.log.Log(l.level, string(d))
	return len(d), nil
}
