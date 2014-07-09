# jott

####A note-taking program for your terminal.

Planned features (in the form of command line args):

* list | list x:        list the latest 5/x jots

* new  | new text-following:    insert a new jot

* cp x | copy x:         copy the jot with canonical ID x

* del x | rm x:           delete the jot with canonical ID x

* purge:                 purge the entire jot database from disk (with [y/N] confirmation)

Makes use of [tiedot](http://github.com/HouzuoGuo/tideot) for data storage and retrieval.
Places .jott directory in the user's $HOME.

Installation:

1. Install [Go](http://golang.org) and follow [this guide](http://golang.org/doc/code.html) on setting up your Go workspace.

2. Execute the following in your terminal:
    ```
    go get github.com/tsums/jott
    go install github.com/tsums/jott
    ```
3. Enjoy running __jott__!
