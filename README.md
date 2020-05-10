# slack-file
Show slack file list, download slack files, delete slack files by channel and period of time.

**Table of Contents**

- [Install](#install)
- [Usage](#usage)
    - [Show slack file list](#show_slack_file_list)
    - [Show slack file list at channel id XXXXXXXXX](#show_slack_file_list_at_channel_id_xxxxxxxxx)
    - [Show slack file list at channel name XXXXXX](#show_slack_file_list_at_channel_name_xxxxxx)
    - [Show slack file that are filtered by type](#show_slack_file_that_are_filtered_by_type)
    - [Show slack files that are more than 30 days old](#show_slack_files_that_are_more_than_30_days_old)
    - [Show slack files that are more than the end of last month old](#show_slack_files_that_are_more_than_the_end_of_last_month_old)
    - [Show slack files that are more than the end of two months old](#show_slack_files_that_are_more_than_the_end_of_two_months_old)
    - [Download slack files](#download_slack_files)
    - [Delete slack files](#delete_slack_files)
    - [Show, download, delete slack files](#show_download,_delete_slack_files)
    - [Set options by .env file](#set_options_by_env_file)
    - [Set options by specified .env file](#set_options_by_specified_env_file)
- [Flags](#flags)


## Install

Require Golang environment and setup GOPATH.

```
$ go get github.com/tmyymmt/slack-file
```

## Usage

Set your slack api token.

### Show slack file list
```
$ slack-file --list --token xxxxxx-xxxxxxxxx
```

### Show slack file list at channel id XXXXXXXXX
```
$ slack-file --list --channel-id=XXXXXXXXX --token xxxxxx-xxxxxxxxx
```

### Show slack file list at channel name XXXXXX
```
$ slack-file --list --channel-name=XXXXXX --token xxxxxx-xxxxxxxxx
```

### Show slack file that are filtered by type

See https://api.slack.com/methods/files.list
```
$ slack-file --list --types=images --token xxxxxx-xxxxxxxxx
```

### Show slack files that are more than 30 days old
```
$ slack-file --list --before-days 30 --token xxxxxx-xxxxxxxxx
```

### Show slack files that are more than the end of last month old
```
$ slack-file --list --before-end-of-month=1 --token xxxxxx-xxxxxxxxx
```

### Show slack files that are more than the end of two months old
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
./downloads_20200510/
```

### Delete slack files
```
$ slack-file --delete 
```

### Show, download, delete slack files

Show, download, delete slack files at channel name XXXXXX that are more than more than the end of last month old.
```
$ slack-file --list --download --delete --channel-name=XXXXXX --before-end-of-month=1 --token xxxxxx-xxxxxxxxx
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

### Set options by specified .env file

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
| channel-id | Filter files by channel id | all |  |
| channel-name | Filter files by channel name | all |  |
| types | Filter files by type | all |  |
| before-days | Filter files by more than ? days old | now |  |
| before-end-of-month | Filter files by more than the end of ? month ago | now |  |
| to | Specified download folder |  | false |
| to-with-date | Add date info to download folder |  | false |
| env | Specitied .env file | .env |  |
