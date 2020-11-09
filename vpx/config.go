package vpx

import "C"

func (c *CodecEncCfg) Load() {
	c.ref37e25db9.g_w = C.uint(c.GW)
	c.ref37e25db9.g_h = C.uint(c.GH)
	c.ref37e25db9.g_timebase.num = C.int(c.GTimebase.Num)
	c.ref37e25db9.g_timebase.den = C.int(c.GTimebase.Den)
	c.ref37e25db9.rc_target_bitrate = C.uint(c.RcTargetBitrate)
	c.ref37e25db9.g_error_resilient = 1
}
