# Donezo

[![Go Report Card](https://goreportcard.com/badge/github.com/GaruGaru/donezo)](https://goreportcard.com/report/github.com/GaruGaru/donezo)


Donezo is a cli helper for devs that needs to work instead of getting micromanaged

## Installation

```bash

go get -u github.com/GaruGaru/donezo

```

## Usage


Add task 

```bash

donezo add "today i've fixed the p=np problem!"

```

List current task

```bash

donezo list today

```

List yesterday's task

```bash

donezo list yesterday

```

List specific date tasks

```bash

donezo list day 01-02-2006

```

## Build locally


```bash

git clone https://github.com/GaruGaru/donezo.git

cd donezo/

make install

```


