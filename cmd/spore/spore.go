package spore

import (
	. "github.com/l0k18/spore/pkg/log"
	"github.com/l0k18/spore/pkg/util"
	"path/filepath"
	"runtime"
)

type Shell struct {
	dataDir string
	config  string
}

func New() *Shell {
	dd := util.Dir("spore", false)
	s := &Shell{
		dataDir: dd,
		config:  filepath.Join(dd, "spore.json"),
	}
	util.EnsureDir(s.config)
	var err error
	Debug(runtime.GOOS, runtime.GOARCH)
	if info, ok := goVersions[runtime.GOOS][runtime.GOARCH]; ok {
		Debug("downloading go", info)
		var wf string
		if wf, err = util.DownloadFile(s.dataDir, info.url, info.hash); Check(err) {
		}
		Debug("download completed", wf)
	}
	return s
}

type downloadInfo struct {
	url, hash string
}

var goVersions = map[string]map[string]downloadInfo{
	"darwin": {
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.darwin-amd64.tar.gz",
			"50a64d6a7ef85510321f0cbcd64e7c72f7e82e27c22f0ba475b9b6b6213f136e",
		},
	},
	"linux": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-386.tar.gz",
			"b08e088ba99134035782c71aeaf139f36d2306eb88eddc22c1278b8b446f157e",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-amd64.tar.gz",
			"6f1354c9040d65d1622b451f43c324c1e5197aa9242d00c5a117d0e2625f3e0d",
		},
		"arm64": downloadInfo{
			"https://golang.org/dl/go1.14.14.linux-arm64.tar.gz",
			"511d764197121f212d130724afb9c296f0cb4a22424e5ae956a5cc043b0f4a29",
		},
		"armv6l": downloadInfo{
			"https://golang.org/dl/go1.15.7.linux-armv6l.tar.gz",
			"e4d614c23b77a367becaeac3032cf4911793363a33efa299d29440be3d66234b",
		},
	},
	"windows": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.windows-386.zip",
			"60ebb9f44549f4827bd29bab822ad881cec6d0f83fff49bda7ad20e69b7b4e7b",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.windows-amd64.zip",
			"88e6be798902d802481b83015e23f6e587cbe0e58766dfa7959d1032865f6bab",
		},
	},
	"freebsd": {
		"386": downloadInfo{
			"https://golang.org/dl/go1.14.14.freebsd-386.tar.gz",
			"7865dffe01499e5e26a40ebc15e068e683e64a2f2edff7440fc9802b02f122bb",
		},
		"amd64": downloadInfo{
			"https://golang.org/dl/go1.14.14.freebsd-amd64.tar.gz",
			"a4fab9549523eefe4cdb4d1334144cb51825db2cfe7993497773f5c9349f6647",
		},
	},
}

/*
go1.14.14.darwin-amd64.tar.gz 	Archive 	macOS 	x86-64 	120MB 	50a64d6a7ef85510321f0cbcd64e7c72f7e82e27c22f0ba475b9b6b6213f136e
go1.14.14.linux-386.tar.gz 	Archive 	Linux 	x86 	100MB 	b08e088ba99134035782c71aeaf139f36d2306eb88eddc22c1278b8b446f157e
go1.14.14.linux-amd64.tar.gz 	Archive 	Linux 	x86-64 	118MB 	6f1354c9040d65d1622b451f43c324c1e5197aa9242d00c5a117d0e2625f3e0d
go1.14.14.linux-arm64.tar.gz 	Archive 	Linux 	ARMv8 	97MB 	511d764197121f212d130724afb9c296f0cb4a22424e5ae956a5cc043b0f4a29
go1.14.14.linux-armv6l.tar.gz 	Archive 	Linux 	ARMv6 	97MB 	e4d614c23b77a367becaeac3032cf4911793363a33efa299d29440be3d66234b
go1.14.14.windows-386.zip 	Archive 	Windows 	x86 	113MB 	60ebb9f44549f4827bd29bab822ad881cec6d0f83fff49bda7ad20e69b7b4e7b
go1.14.14.windows-amd64.zip 	Archive 	Windows 	x86-64 	132MB 	88e6be798902d802481b83015e23f6e587cbe0e58766dfa7959d1032865f6bab
go1.14.14.freebsd-386.tar.gz 	Archive 	FreeBSD 	x86 	100MB 	7865dffe01499e5e26a40ebc15e068e683e64a2f2edff7440fc9802b02f122bb
go1.14.14.freebsd-amd64.tar.gz 	Archive 	FreeBSD 	x86-64 	118MB 	a4fab9549523eefe4cdb4d1334144cb51825db2cfe7993497773f5c9349f6647

*/
