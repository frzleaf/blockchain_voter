package types

func NewMsgCreatePool(creator string, title string, options []string) *MsgCreatePool {
	return &MsgCreatePool{
		Creator: creator,
		Title:   title,
		Options: options,
	}
}
