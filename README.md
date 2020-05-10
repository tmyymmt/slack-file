# slack-file

Show slack file list, download slack files, delete slack files by channel and period of time.

**Table of Contents**

- [Install](#install)
- [Usage](#usage)
    - [Show slack files](#show-slack-files)
    - [Download slack files](#download-slack-files)
    - [Delete slack files](#delete-slack-files)
    - [Show, download, delete slack files](#show-download-delete-slack-files)
    - [Set options by .env file](#set-options-by-env-file)
- [Options](#options)


## Install

Require Golang environment and setup GOPATH.

```
$ go get github.com/tmyymmt/slack-file
```

## Usage

Set your slack api token.


### Show slack files

Show slack files.
```
$ slack-file --show --token xxxxxx-xxxxxxxxx
```

Show slack files at channel id XXXXXXXXX.
```
$ slack-file --show --channel-id=XXXXXXXXX --token xxxxxx-xxxxxxxxx
```

Show slack files that are filtered by type. See https://api.slack.com/methods/files.list
```
$ slack-file --show --types=images --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than 30 days old.
```
$ slack-file --show --before-days 30 --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than the end of last month old.
```
$ slack-file --show --before-end-of-month=1 --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than the end of two months old.
```
$ slack-file --show --before-end-of-month=2 --token xxxxxx-xxxxxxxxx
```

### Download slack files

Default folder is `./downloads`
```
$ slack-file --download --token xxxxxx-xxxxxxxxx
```

If you want to download slack files to specified download folder then set `--to` option`.
```
$ slack-file --download --to ./folder1 --token xxxxxx-xxxxxxxxx
```

If you want to add date info to download folder then set `--to-with-date` option`.
```
$ slack-file --download --to-with-date --token xxxxxx-xxxxxxxxx
```
```
./downloads_YYYYMMDD/
```

### Delete slack files

Delete slack files.
```
$ slack-file --delete 
```

### Show, download, delete slack files

Show, download, delete slack files.
```
$ slack-file --show --download --delete --token xxxxxx-xxxxxxxxx
```

Show, download, delete slack files at channel id to downloads_YYYYMMDD folder XXXXXX that are more than more than the end of last month old.
```
$ slack-file --show --download --delete --channel-id=XXXXXX --before-end-of-month=1 --to-with-date --token xxxxxx-xxxxxxxxx
```

### Set options by .env file

Set options capitalized and replaced '-' to '_'.

.env
```
TOKEN=xxxxxx-xxxxxxxxx
CHANNEL_ID=XXXXXX
BEFORE_END_OF_MONTH=1
```

```
$ slack-file --list --download --delete
```

## Options

| name | description | default | require |
| :--- | :---------- | :-----: | :-----: |
| token | Your slack api token |  | true |
| show | Show slack files | true | false |
| download | Download file from slack | false | false |
| delete | Delete slack files | false | false |
| channel-id | Filter files by channel id | all | false |
| exclude-channel-ids | Filter files by excluded channel ids | all | false |
| types | Filter files by type | all | false |
| before-timestamp | Filter files by before the timestamp | now | false |
| before-days | Filter files by more than ? days old | now | false |
| before-end-of-month | Filter files by more than the end of ? month(s) ago | now | false |
| to | Download slack files to specified download folder | downloads | false |
| to-with-date | Add date info to download folder name | false | false |
