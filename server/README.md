# Server

## Overview

This server handles requests for the splendor game. It exposes apis to create &
join games, update game state (which tracks ongoing games as well as event-logs
for all games,) and player maintenance.

## Lifecycle of a Game

### Glossary

#### Player

A player is a human entity that interfaces with a game. They are agents that
initiate actions in a game.

#### Player's Hand

A [Player's](####Player) hand refers to all of the resources a player currently
holds during a game. These resources can be meaningfully divvied up into coins,
cards, and nobles. A card may be either owned or reserved. A reserved card may
be shown or hidden. A player may only hold a max of 3 reserved cards.

#### Game

The entity an agent interacts with and the core component of splendor

#### Game Id

Entity that will be used to generate a game format <low>-<med>-<high>-<noble>
aprox. "XXXX-XXXX-XXXX-XXX"

#### Card

An entity within a game that a has a value (>=0), a relative value (low, med,
high), a gem type (blue, red, green, brown, white) and a cost ([0-5].)

#### Coin

An entity within a game that has a value (1), a coin type (blue, red, green,
brown, white, ALL). These can only be used to purchase cards.

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

<!-- TODO(GikuyuNderitu): Elaborate Gameplay -->
