### Watch Twitch

## Overview

Small cli program to view your followed streamers who are currently live.
Prompts user for if they want to open just the stream, chat, or both.
Uses streamlink, mpv, and chatterino so you can enjoy twitch.  No need for the browser.

## Requirements

Download and install the following tools:
* streamlink
* mpv
* chatterino

## Env Vars

https://streamlink.github.io/cli/plugins/twitch.html#authentication

USER_ID -- Account ID of streamer
USER_ACCESS_TOKEN -- OAUTH token w/ user:read:follows permissions used for getting followers
BROWSER_AUTH_TOKEN -- Used for watching authenticated video feed.  Used to skip adds if subbed to the stream you are watching or if you have turbo. See above link on how to acquire this token.
CLIENT_ID -- client ID of twitch dev app
CLIENT_SECRET -- client secret of twitch dev app

