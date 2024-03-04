# tChat

Terminal chat application for developers. Share texts and files between users (or servers) using end-to-end encryption.

## Features

- [x] Send text messages
- [x] Send files
- [ ] Secure authentication
- [x] End-to-end encryption
- [ ] Live chat

## Working

### Public Key Sharing

![Public Key Sharing Sequence Diagram](./assets/pubkey-sharing-seq-diagram.png)

### End-to-end Encryption

![End-to-end Encryption Sequence Diagram](./assets/e2ee-seq-diagram.png)

## Installation

Git clone the repository and run the following commands:

1. For Mac and Linux:

```bash
sudo sh install.sh
```

This will build the application and copy the executable to /usr/local/bin directory.

2. For Windows: (using PowerShell)

```bash
.\install.ps1
```

## Usage

To start the application, run the following command:

```bash
tchat
```

Please run the following command to start a new connection:

```bash
tchat conn
```

## Commands

| Command | Description              | Example                                 |
| ------- | ------------------------ | --------------------------------------- |
| `conn`  | Start or save connection | `tchat conn`                            |
| `msg`   | Receive messages         | `tchat msg`                             |
| `send`  | Send messages            | `tchat send -r <receiver> -m <message>` |

### Message Command

| Option | Description      | Example            |
| ------ | ---------------- | ------------------ |
| -c     | Clear the chat   | `tchat msg -c all` |
| -d     | Display the chat | `tchat m -d 10`    |

### Send Command

| Option | Description | Example                                 |
| ------ | ----------- | --------------------------------------- |
| -r     | Receiver    | `tchat send -r <receiver> -m <message>` |
| -m     | Message     | `tchat send -r <receiver> -m <message>` |
| -f     | File        | `tchat send -r <receiver> -f <file>`    |
