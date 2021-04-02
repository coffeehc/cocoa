package foundation

type URLRequestCachePolicy uint

const (
	URLRequestUseProtocolCachePolicy                URLRequestCachePolicy = 0
	URLRequestReloadIgnoringLocalCacheData          URLRequestCachePolicy = 1
	URLRequestReloadIgnoringLocalAndRemoteCacheData URLRequestCachePolicy = 4
	URLRequestReloadIgnoringCacheData               URLRequestCachePolicy = URLRequestReloadIgnoringLocalCacheData
	URLRequestReturnCacheDataElseLoad               URLRequestCachePolicy = 2
	URLRequestReturnCacheDataDontLoad               URLRequestCachePolicy = 3
	URLRequestReloadRevalidatingCacheData           URLRequestCachePolicy = 5
)
