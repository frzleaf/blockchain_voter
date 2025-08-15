package types

func NewMsgCreatePoll(creator string, title string, options []string) *MsgCreatePoll {
	return &MsgCreatePoll{
		Creator: creator,
		Title:   title,
		Options: options,
	}
}
