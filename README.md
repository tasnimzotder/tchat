# tChat

Terminal chat application for developers

## Features

- [x] Send text messages 
- [ ] Send files and images
- [ ] Secure authentication
- [ ] End-to-end encryption
- [ ] Live chat

## Installation

Git clone the repository and run the following commands:

```bash
sudo sh install.sh
```

This will build the application and copy the executable to /usr/local/bin directory.

## Usage

To start the application, run the following command:

```bash
tchat
```

Please run the following command to register your username:

```bash
tchat init -u <username>
```

## Commands

| Command | Description            | Example                                  |
|---------|------------------------|------------------------------------------|
| init    | Register your username | `tchat init -u <username>`               |
| s       | Send a message         | `tchat send -r <recipient> -m <message>` |
| m       | Message                | `tchat m`                                |

### Message Command

| Option | Description      | Example          |
|--------|------------------|------------------|
| -c     | Clear the chat   | `tchat m -c all` |
| -d     | Display the chat | `tchat m -d 10`  |
