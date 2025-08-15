package types

func NewMsgCastVote(creator string, pollId string, option string) *MsgCastVote {
	return &MsgCastVote{
		Creator: creator,
		PollId:  pollId,
		Option:  option,
	}
}
