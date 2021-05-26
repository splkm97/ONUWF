package main

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

// role werewolf struct
type werewolf struct {
	// role을 상속받는 구조체.
	role
}

// 늑대인간을 반환하는 함수.
func (wf *werewolf) String() string {
	return "늑대인간"
}

// 늑대인간의 능력을 사용하는 함수.
func (wf *werewolf) Action(uid1, uid2 string, disRole int, player *user, g *game, s *discordgo.Session) {
	if disRole != -1 {
		// 늑대인간이 유일한 경우.
		target := wf.getDiscard(disRole, g)
		msg := "당신은 `" + target.String() + "`이(가) 버려진 것을 확인했습니다."
		s.ChannelMessageSendEmbed(embed.NewGenericEmbed("버려진 카드 1장 확인", msg))
		g.appendLog("`" + user.nick + "` 은(는) 유일한 늑대인간이었습니다.")
		g.appendLog("`" + user.nick + "` 은(는) 버려진 직업 `" + target.String() + "`을(를) 확인했습니다.")
	} else {
		// 늑대인간이 여럿인 경우.
		wolves := wf.getRoleUser(wf, g)
		mistic := wf.getRoleUser(misticWolf{}, g)
		alpha := wf.getRoleUser(alphaWolf{}, g)
		dream := wf.getRoleUser(dreamWolf{}, g)
		var wolflist string
		var dreamlist string
		for _, item := range wolves {
			wolflist += "`" + item.nick + "` "
		}
		for _, item := range mistic {
			wolflist += "`" + item.nick + "` "
		}
		for _, item := range alpha {
			wolflist += "`" + item.nick + "` "
		}
		for _, item := range dream {
			dreamlist += "`" + item.nick + "` "
		}
		msg := "당신의 동료 늑대인간은\n"
		msg += wolflist + dreamlist
		msg += "\n... 입니다."
		s.ChannelMessageSendEmbed(embed.NewGenericEmbed("동료 늑대인간 확인", msg))
		if dream != nil {
			g.appendLog("잠자는 늑대 " + dreamlist + "을(를) 제외한 모든 늑대인간들은 서로를 확인했습니다.")
		} else {
			g.appendLog("늑대인간들은 서로를 확인했습니다.")
		}
		g.appendLog("늑대인간들은 그 시점에서 " + wolflist + dreamlist + "였습니다.")
	}
}
