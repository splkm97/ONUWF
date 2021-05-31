package main

import embed "github.com/clinet/discordgo-embed"

type roleWerewolf struct {
	role
}

func (wf *roleWerewolf) Action(tar targetObject, player *user, g *game) {
	if disRoleIdx != -1 {
		target := wf.getRoleFromDiscard(disRoleIdx, g)
	} else {
		wolves := wf.getRoleUser(wf, g)
		wolves = append(wolves, wf.getRoleUser(roleMisticwolf{}, g)...)
		wolves = append(wolves, wf.getRoleUser(roleAlphawolf{}, g)...)
		dreams = wf.getRoleUser(roleDreamwolf{}, g)
		var wolflist string
		for _, item := range wolves {
			wolflist += "`" + item.nick + "` "
		}
		for _, item := range dreams {
			dreamlist += "`" + item.nick + "` "
		}
		wolflist += dreamlist
		msg := "당신의 동료 늑대인간은\n"
		msg += wolflist
		msg += "\n ... 입니다."
		s.ChannelMessageSendEmbed(embed.NewGenericEmbed("동료 늑대인간 확인", msg))
		if len(dreams) == 0 {

		}
	}
}

func (wf *roleWerewolf) String() {
	return "늑대인간"
}
