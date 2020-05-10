# slack-file

Show slack file list, download slack files, delete slack files by channel and period of time.

**Table of Contents**

- [Install](#install)
- [Usage](#usage)
    - [Show slack file list](#show-slack-file-list)
    - [Download slack files](#download-slack-files)
    - [Delete slack files](#delete-slack-files)
    - [Show, download, delete slack files](#show-download-delete-slack-files)
    - [Set options by .env file](#set-options-by-env-file)
- [Flags](#flags)


## Install

Require Golang environment and setup GOPATH.

```
$ go get github.com/tmyymmt/slack-file
```

## Usage

Set your slack api token.


### Show slack file list

Show slack file list.
```
$ slack-file --list --token xxxxxx-xxxxxxxxx
```

Show slack file list at channel id XXXXXXXXX.
```
$ slack-file --list --channel-id=XXXXXXXXX --token xxxxxx-xxxxxxxxx
```

Show slack file list at channel name XXXXXX.
```
$ slack-file --list --channel-name=XXXXXX --token xxxxxx-xxxxxxxxx
```

Show slack file that are filtered by type. See https://api.slack.com/methods/files.list
```
$ slack-file --list --types=images --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than 30 days old.
```
$ slack-file --list --before-days 30 --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than the end of last month old.
```
$ slack-file --list --before-end-of-month=1 --token xxxxxx-xxxxxxxxx
```

Show slack files that are more than the end of two months old.
```
$ slack-file --list --before-end-of-month=2 --token xxxxxx-xxxxxxxxx
```

### Download slack files

Default folder is `./downloads`
```
$ slack-file --download --token xxxxxx-xxxxxxxxx
```

If you want to specified download folder then set `--to` option`.
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
$ slack-file --list --download --delete --token xxxxxx-xxxxxxxxx
```

Show, download, delete slack files at channel name to downloads_YYYYMMDD folder XXXXXX that are more than more than the end of last month old.
```
$ slack-file --list --download --delete --channel-name=XXXXXX --before-end-of-month=1 --to-with-date --token xxxxxx-xxxxxxxxx
```

### Set options by .env file

.env
```
token=xxxxxx-xxxxxxxxx
channel-name=XXXXXX
before-end-of-month=1
```

Show slack file list, download slack files, delete slack files at channel name XXXXXX that are more than the end of last month old.
```
$ slack-file --list --download --delete
```

### Set options by specified .env file.

.env.prod
```
token=xxxxxx-xxxxxxxxx
channel-name=XXXXXX
before-end-of-month=1
```

Show slack file list, download slack files, delete slack files at channel name XXXXXX that are before the end of last month.
```
$ slack-file --list --download --delete --env .env.prod
```

## Flags

| name | description | default | require |
| :--- | :---------- | :-----: | :-----: |
| token | Your slack api token |  | true |
| list | Show slack file list | true | false |
| download | Downloaded file from slack | false | false |
| delete | Delete slack files | false | false |
| channel-id | Filter files by channel id | all | false |
| channel-name | Filter files by channel name | all | false |
| types | Filter files by type | all | false |
| before-days | Filter files by more than ? days old | now | false |
| before-end-of-month | Filter files by more than the end of ? month ago | now | false |
| to | Specified download folder | downloads | false |
| to-with-date | Add date info to download folder | false | false |
| env | Specitied .env file | .env | false |
