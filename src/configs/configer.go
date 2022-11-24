package configs

import(
	"runtime"
	"path/filepath"
)

var (
	_, f, _, _ = runtime.Caller(0)
	ProjectConfigrePath = filepath.Join(f, "../")
)