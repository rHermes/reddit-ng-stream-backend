# `reddit-ng-stream-backend`

The backend for my reddit stream project. It converts [the pushshift.io stream](http://stream.pushshift.io) 
from SSE to websockets. The reason this is needed is that there is currently a SSE bug in firefox
which prevents the connection on startup.

## Credits

The websocket parts are shamelessly stolen from the excent "chat" example from the gorilla
websockets library.
