// +build linux,amd64,go1.15,!cgo

package role

import embed "github.com/clinet/discordgo-embed"

type RoleWerewolf struct {
	role
}

func (wf RoleWerewolf) Action(tar targetObject, player *user, g *game) {
	switch tar.actionType {
	case 1:
		//recvRole := g.getDisRole(tar.disRoleIdx)
	case 0:
		wolves := g.getRoleUsers(wf)
		//wolves = append(wolves, g.getRoleUsers(roleMisticwolf{})...)
		//wolves = append(wolves, g.getRoleUsers(roleAlphawolf{})...)
		//dreams := g.getRoleUsers(roleDreamwolf{})
		var wolflist string
		//var dreamlist string
		for _, item := range wolves {
			wolflist += "`" + item.nick + "` "
		}
		//for _, item := range dreams {
		//	dreamlist += "`" + item.nick + "` "
		//}
		//wolflist += dreamlist
		msg := "당신의 동료 늑대인간은\n"
		msg += wolflist
		msg += "\n ... 입니다."
		//if len(dreams) == 0 {
		//	msg += "\n\n"
		//	msg += dreamlist + "는 잠에 빠져 서로를 확인하지 못하였지만,"
		//	msg += "당신의 동료 늑대인간입니다."
		//}
		curSession.ChannelMessageSendEmbed(player.dmChanID, embed.NewGenericEmbed("동료 늑대인간 확인", msg))
	default:
	}
}

func (wf RoleWerewolf) String() string {
	return "늑대인간"
}
