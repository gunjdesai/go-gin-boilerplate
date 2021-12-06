package helpers

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetEnv() string {

	var env string

	switch os.Getenv(ENV_VAR) {

	case DEVELOPMENT:
		env = DEVELOPMENT
	case STAGING:
		env = STAGING
	case PRODUCTION:
		env = PRODUCTION
	default:
		env = PRODUCTION

	}

	return env

}

func DatesAreEqual(t1 time.Time, t2 time.Time) bool {

	return t1.Truncate(24 * time.Hour).Equal(t2.Truncate(24 * time.Hour))

}

func DatesLiesWithinRange(t time.Time, lb time.Time, ub time.Time) bool {

	return lb.Truncate(24*time.Hour).After(t.Truncate(24*time.Hour)) &&
		ub.Truncate(24*time.Hour).Before(t.Truncate(24*time.Hour)) ||
		t.Truncate(24*time.Hour).Equal(lb.Truncate(24*time.Hour)) ||
		t.Truncate(24*time.Hour).Equal(ub.Truncate(24*time.Hour))

}
