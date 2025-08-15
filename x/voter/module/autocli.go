package voter

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"voter/x/voter/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePool",
					Use:            "create-pool [title] [options]",
					Short:          "Send a create-pool tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "title"}, {ProtoField: "options", Varargs: true}},
				},
				{
					RpcMethod:      "CastVote",
					Use:            "cast-vote [poll-id] [option]",
					Short:          "Send a cast-vote tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "poll_id"}, {ProtoField: "option"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
