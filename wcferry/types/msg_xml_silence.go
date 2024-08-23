package types

// 转账、收款
// @msg.type 1
type MsgXmlSilence struct {
	Msgsource struct {
		Membercount string `xml:"membercount"`
		Signature   string `xml:"signature"`
		Silence     string `xml:"silence"`
		TmpNode     struct {
			PublisherId string `xml:"publisher-id"`
		}
	} `xml:"msgsource"`
}
