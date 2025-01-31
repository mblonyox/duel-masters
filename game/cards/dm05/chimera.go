package dm05

import (
	"fmt"

	"duel-masters/game/civ"
	"duel-masters/game/family"
	"duel-masters/game/fx"
	"duel-masters/game/match"
	"duel-masters/game/cnd"
)

func Gigazoul(c *match.Card) {

	c.Name = "Gigazoul"
	c.Power = 3000
	c.Civ = civ.Darkness
	c.Family = family.Chimera
	c.ManaCost = 3
	c.ManaRequirement = []string{civ.Darkness}

	c.Use(fx.Creature, func(card *match.Card, ctx *match.Context) {

		if _, ok := ctx.Event.(*match.UntapStep); ok {

			opponentShields := fx.Find(ctx.Match.Opponent(card.Player), match.SHIELDZONE)
			n := len(opponentShields)

			if n == 0 {
				card.AddCondition(cnd.CantAttackPlayers, true, card.ID)
				card.AddCondition(cnd.CantAttackCreatures, true, card.ID)
			} else {
				card.RemoveCondition(cnd.CantAttackPlayers)
				card.RemoveCondition(cnd.CantAttackCreatures)
			}

		}

		if event, ok := ctx.Event.(*match.CardMoved); ok {

			if event.From == match.SHIELDZONE || event.To == match.SHIELDZONE {

				opponentShields := fx.Find(ctx.Match.Opponent(card.Player), match.SHIELDZONE)
				n := len(opponentShields)

				if n == 0 {
					card.AddCondition(cnd.CantAttackPlayers, true, card.ID)
					card.AddCondition(cnd.CantAttackCreatures, true, card.ID)
				} else {
					card.RemoveCondition(cnd.CantAttackPlayers)
					card.RemoveCondition(cnd.CantAttackCreatures)
				}

			}

		}

		if event, ok := ctx.Event.(*match.AttackPlayer); ok {

			// Is this event for me or someone else?
			if event.CardID != card.ID || !card.HasCondition(cnd.CantAttackPlayers) {
				return
			}

			ctx.Match.WarnPlayer(card.Player, fmt.Sprintf("%s can't attack players", card.Name))

			ctx.InterruptFlow()

		}

		if event, ok := ctx.Event.(*match.AttackCreature); ok {

			// Is this event for me or someone else?
			if event.CardID != card.ID || !card.HasCondition(cnd.CantAttackCreatures) {
				return
			}

			ctx.Match.WarnPlayer(card.Player, fmt.Sprintf("%s can't attack creatures", card.Name))

			ctx.InterruptFlow()

		}

	})

}
