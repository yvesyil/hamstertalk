# Hamstertalk 🐹

## About

Hamstertalk is a simple chat server state machine.

## Protocol

### Message Format

```txt
(UTC_TIMESTAMP)~HOUSE_ID~:HAMSTER_ID:@HAMSTER_NICKNAME@ :: MESSAGE_BODY
```

### Message Body

```txt
COMMAND [OPTION[=OPT_VALUE]] ... -- [ARG] ...
```

## Commands

### msg

Send an encoded message. Default encoding is utf-8.

```txt
msg [en=ENCODING] -- MESSAGE
```

### bytes

Send a message in bytes.

```txt
bytes -- MESSAGE
```

### set

Sets an option for a hamster.

```txt
set -- [OPTION[=OPT_VALUE]]
```

### goto

Join the house with the given ID.

```txt
goto -- HOUSE_ID
```

### list

List the available houses or hamsters.

```txt
list -- 
list h=HOUSE_ID -- 
```

### quit

Quit the application.

```txt
quit -- 
```
