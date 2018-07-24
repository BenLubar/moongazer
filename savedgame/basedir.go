package savedgame

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

var ErrUnknownBaseDir = errors.New("savedgame: cannot determine base directory for MoonQuest save games")

func FindBaseDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", errors.Wrap(err, "cannot determine current user")
	}
	for _, distributor := range []string{"steam", "itch", "standalone"} {
		if runtime.GOOS == "windows" {
			name := filepath.Join(u.HomeDir, "Saved Games", "moonquest", distributor, "saves")
			if fi, err := os.Stat(name); err == nil && fi.IsDir() {
				return filepath.Dir(name), nil
			}
		}
	}

	return "", ErrUnknownBaseDir
}
