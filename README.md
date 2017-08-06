# meetup-api

Api wrapper for [meetup API v3](https://www.meetup.com/meetup_api/) written in Go.

# Usage

```
  c, err := meetup.NewClient(MEETUP_API_KEY)
  if err != nil {
    log.Fatal(err)
  }

  c.GETEvents(GROUP_NAME) - WIP
```

# Requirement

Meetup API key, can be generated here:
https://secure.meetup.com/meetup_api/key/
