zback-find
==========

Navigate up the folder tree to find a file with a given file name.

Why ?
-----
For scripting purpose mostly. That's a little like `git` works, actually. It detects that there is a `.git`
folder somewhere in a upper directory and allow you to use `git` commands. If the folder does not exist, that simply means
that you are not inside a `git` tracked folder, and you won't be able to perform most of `git` actions.

I often have this use case when I need to find a specific file (or check if it exists) in a upper directory for script purpose.
Additionally, I wanted an easy-to-use solution instead of relying on the same bash script block that only make script
harder to read and manage.

Why Go ?
--------

This is not my main language, but I wanted something simple, as a single executable (with a reasonable size), without having to manage a lot of dependencies,
without worrying how to make it run on different systems, and capable of cross compilation (so I don't have to rely on agents with a specific OS).

Giving all of this, `go` was the more reasonable choice. I narrowed it down to `go` or `rust`, but the cross compilation was the decisive
element for the final choice.