# Hamstertalk 🐹

## About

Hamstertalk is a chat server. It implements its own minimal protocol.

## The Hamster Protocol

It sits on top of HTTP. Everything is expressed using events.

### Event Format

```txt
(UTC_TIMESTAMP)~HOUSE_ID~:HAMSTER_ID:@HAMSTER_NICKNAME@ :: EVENT_BODY
```

### Event Body

```txt
ACTION [OPTION[=OPT_VALUE]] ... -- [ACTION_BODY] ...
```

## Actions

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
