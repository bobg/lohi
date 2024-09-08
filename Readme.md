# Lohi - Library and command-line tool for processing Google Location History data

This is lohi, a software library and command-line tool for processing Google Location History data.

## Installation and usage

For library usage,
please see [the Godoc](https://pkg.go.dev/github.com/bobg/lohi).

To install the command-line tool:

```sh
go install github.com/bobg/lohi/cmd/lohi@latest
```

Command-line usage:

```sh
lohi [-trips] [-creds FILE] [-token FILE] [-cache FILE] [-qps RATE] < Timeline.json
```

The flags and their meanings are:

| Flag         | Meaning                                                                                                       |
|--------------|---------------------------------------------------------------------------------------------------------------|
| -trips       | Show “trips” from the input rather than individual place visits.                                              |
| -creds FILE  | The file holding Google API auth credentials. The default is creds.json. See [Places API](#places-api) below. |
| -token FILE  | The file holding an OAuth token. The default is token.json. See [Places API](#places-api) below.              |
| -cache FILE  | The file that caches results from Google Place API lookups. The default is places.db.                         |
| -qps RATE    | The maximum number of queries per second to send to the Google Places API. The default is 1.                  |

JSON timeline data is read from standard input.

The output is a human-readable summary of place-visit data from the parsed input,
or trip data if the `-trips` flag is specified.

## Background

Google’s Location History feature records the position of capable mobile devices when enabled.
It correlates those saved positions with locations in Google Maps,
and makes educated guesses about how the device got from one location to the next
(“walking,” “in a passenger vehicle,” etc).
The Timeline feature in Google Maps allowed users to browse this history data,
and could also remind users when looking at some map location,
“You visited here six years ago.”

In mid-2024 Google announced that the Timeline feature would no longer be available in the desktop version of Google Maps,
and that Location History data would be removed from Google’s cloud.
Affected users could opt in to transfering that data to their mobile devices instead.
If they did, then the Timeline feature would still be available, but only in the Google Maps app on mobile.

That data can be exported in JSON format.
On an Android phone,
open Settings,
then navigate to Location>Location Services>Timeline>Export Timeline data.
This will create a file for you named Timeline.json.
That’s the data that this package can parse and process.

## Places API

Google’s JSON timeline data represents locations with a short alphanumeric [place ID](https://developers.google.com/maps/documentation/places/web-service/place-id).
To convert these to details such as address, business name, etc.,
it is necessary to query [the Google Places API service](https://developers.google.com/maps/documentation/places/web-service/overview).
In order to do that,
you must supply _authentication credentials_ to lohi so that it can contact the service on your behalf.
This is done with the `-creds` and `-token` flags.

A full discussion of how to obtain the necessary credentials is beyond the scope of this document.
For more information on this topic,
please see [Using OAuth 2.0 to Access Google APIs](https://developers.google.com/identity/protocols/oauth2).

The file for `-creds` must exist, but the file for `-token` normally does not when you first run lohi.
Instead, lohi will read the credentials file and open your web browser,
prompting you to authorize the application.
If you consent, an “authorization token” is written to the `-token` file
and used in calls to the Google Places API.

Usage of the Google Places API is not free;
see [Places API Usage and Billing](https://developers.google.com/maps/documentation/places/web-service/usage-and-billing).
So it’s desirable both to limit how often the service gets queried
(to prevent runaway costs)
and to _cache_ the results of Google Places API lookups,
so that the second and subsequent times a given place ID is needed,
it can be obtained from the cache
(a local file)
rather than from the Google service.
The cache file is named by the `-cache` flag and created if needed.
The rate limit is specified with the `-qps` flag and defaults to one query per second.
