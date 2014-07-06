# jott

A note-taking program for the terminal that will allow you to jot quick text and retrieve them from a user-based database.

Planned features (in the form of command line args):

list | list x:        list the latest 5/x jots
new  | new "text"     insert a new jot
cp x | copy x         copy the jot with canonical ID x
del x| rm x           delete the jot with canonical ID x
purge                 purge the entire jot database from disk (with [y/N] confirmation)

This program is mainly an exercise in initially exploring golang. Don't expect it to work well or be efficient at the moment.
