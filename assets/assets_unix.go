//go:build linux || freebsd || openbsd || netbsd
// +build linux freebsd openbsd netbsd

package assets

import (
	_ "embed"
)

// -----------------------------------------------------------------------------

//go:embed wifi1.png
var Wifi1Image []byte

//go:embed wifi2.png
var Wifi2Image []byte

//go:embed wifi3.png
var Wifi3Image []byte

//go:embed sad_dizzy.png
var SadDizzyImage []byte

//go:embed check_circle.png
var CheckCircleImage []byte
