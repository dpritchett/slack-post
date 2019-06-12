# slack-post

Post messages from ARGV or STDIN to a slack channel

## Usage

Set two environment variables first:
```sh
export SLACK_POSTER_API_TOKEN=redacted
export SLACK_POSTER_DESTINATION="#bot-tests"
```

- Post from STDIN: `echo hi mom | slack-post` 
- Post from ARGV: `slack-post hi mom`
