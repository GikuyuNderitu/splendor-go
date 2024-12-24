# Server

## Overview

This server handles requests for the splendor game. It exposes apis to create &
join games, update game state (which tracks ongoing games as well as event-logs
for all games,) and player maintenance.

## Glossary

### Player

A player is a human entity that interfaces with a game. They are agents that
initiate actions in a game.

### Player's Hand

A [Player's](#player) hand refers to all of the resources a player currently
holds during a game. These resources can be meaningfully divvied up into coins,
cards, and nobles. A card may be either owned or reserved. A reserved card may
be shown or hidden. A player may only hold a max of 3 reserved cards.

### Game

The entity an agent interacts with and the core component of splendor

### Round

Combination of each unique player taking 1 turn.

### Action

A discrete thing an agent can do during its turn during a round. This can be
one of:

- Taking 3 coins of different colors
- Taking 2 coins of the same color. A player may only take 2 coins (of the same
  type) when there are 4 or more coins of that type.
- Reserving a card with a gold coin
- Purchasing a card

#### Selecting Coins

A player can never hold more than 10 tokens at the end of their turn, including
gold coins. If the player reaches the end of their turn and they have more than
10 coins, they can return all or some of those **they just drew**. A player's
tokens are known to all other players.

#### Reserving a Card

To reserve a card, a player may take a face-up card from the table. They may
also reserve the top card from one of the three decks **without** displaying
it to the other players. A gold coin is used to reserve a card (this is the
only way to obtain a gold coin.) If there are no more gold coins, a player
_can_ still reserve their desired card, but they will not be able to obtain a
gold coin.

A player only displays the reserved cards that were on faceup on
the table, cards reserved when they were still in the deck are hidden.

#### Purchasing a Card

To purchase a card, a player must spend the number of coins listed on the card.
A gold coin can be used for a stand in for any of the colors. The cards must be
arranged such that their bonus and value are visible.

Players will use the associated bonus value (color of the card) to discount
future card purchases and accquire nobles.

### Game Id

Entity that will be used to generate a game format <low>-<med>-<high>-<noble>
aprox. "XXXX-XXXX-XXXX-XXX"

### Card

An entity within a game that a has a value (>=0), a relative value (low, med,
high), a bonus type (blue, red, green, brown, white) and a cost ([0-5].) The
bonus type indicates that the corresponding coin type is discounted by 1.

### Coin

An entity within a game that has a value (1), a coin type (blue, red, green,
brown, white, ALL). These can only be used to purchase cards.

### Noble

Visible at all times on the table. They have a point value and multiple bonus
values. A noble visits a player at the end of the player's turn if the bonus
type conditions are met. If a player is eligible for more than 1 noble, they
will chose which noble they would like to visit them.

## Lifecycle of a Game

### Game Creation

A game must first be created by a player. Even though the game is created, this
does not mean it has yet started. Upon creation, the player may provide a valid
id. If no id is provided, an id will be generated. the game will be assigned
its own unique-id. This id should be able to deterministically dictate the order
the different cards are sorted, and the order that the nobles are displayed.

The game may be **joined** by up to 4 players, (typically, one of the players
that join is the creator. They have no special privleges.) Once all expected
players have joined, the game may begin.

<!-- TODO(GikuyuNderitu): Add section for allowing non-player agents to join -->

### Game Starting

Once all players have joined a game, the game may **start**. The initial game
state will be generated depending on the number of agents a game has.

<!-- TODO(GikuyuNderitu): Elaborate on the different initial states. -->

The cards and nobles will be generated based on the game id. The cards will be
generated/randomized the same every time. The nobles will also be randomized
the same, (ie. all nobles will be considered during randomization), but only
the first N nobles will be chosen after randomization (N = # players + 1.)

### Gameplay

The first player is chosen at random. They mark the beginning and end of a
"[Round](#round)".

Each player takes an [action](#action).

When taking coins, the coins chosen this round will be held in an "escrow"
state until the next player advances. This facilitates the ability for a player
to return coins if they have over-capped their coin collection when their
action was "reserving" or "taking coins".

After they take their action, a [noble](#noble) decides if they will visit the
player that just took an action.

### Win Condition

At the end of each player's action, the game checks to see if the player is in a
winning state. If the player is in a winning state (15 points), the game will
attempt to finish out the round so that each player has an opportunity to
perform the same number of actions.

After all players have taken the same number of actions, no more actions may be
taken. The player with the most points is considered the winner. If there is a
tie, the player with the fewest number of cards is considered the winner.
