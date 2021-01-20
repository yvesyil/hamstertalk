# Hamstertalk 🐹

## A unique kind of chatting platform that dwells happily in the console

### About

Hamstertalk is similar to IRC clients. But it has a unique feature called tunnels. Tunnels allow hamsters (users) to discover and explore other houses (chat rooms) with sequence. You can use a tunnel and quickly go to the next house that's linked to that tunnel.

### Commands

- `!set nickname <nickname>` sets a nickname for you. If a hamster doesn't have a nickname, they will be displayed as "anon".
- `!use tunnel <tunnelid>` join to a tunnel! Explore other houses.
- `!stepto next | previous` moves to the next or previous house that is connected to the currently joined tunnel.
- `!set <houseid> | <tunnelid> private | public` sets the given house to public or pirvate.
- `!hopto <houseid>` joins to the specified house.
- `!list [tunnels | hamsters]` lists other available tunnels or other fellow hamsters in the current house.
- `!squeakto | sq <nickname> <message>` privately sends a message to the hamster with the specified nickname.
- `!exit [tunnelid]` if no tunnel is specified, exits from the current house, otherwise disconnects from the given tunnel.
- `!quit` close the application, bye bye.
