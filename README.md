# spidercm


```sh
$ touch 1.txt

$ mkdir something/

$ touch something/2.bin

$ spidercm init

$ ls -lha
total 5.4M
drwxrwxr-x 5 fmn fmn 4.0K Apr  2 01:19 .
drwxrwxr-x 6 fmn fmn 4.0K Apr  2 01:19 ..
-rw-rw-r-- 1 fmn fmn    0 Apr  2 01:19 1.txt
-rw-r--r-- 1 fmn fmn 224K Apr  2 01:19 .fossil.sqlite3
-rw-r--r-- 1 fmn fmn  32K Apr  2 01:19 .fslckout
drwxrwxr-x 7 fmn fmn 4.0K Apr  2 01:19 .git
drwxrwxr-x 5 fmn fmn 4.0K Apr  2 01:19 .hg
drwxrwxr-x 2 fmn fmn 4.0K Apr  2 01:19 something

$ spidercm add -p 1.txt something/2.bin 

$ spidercm commit -m 'something something'

$ git log
commit d9439453d9731e2d7e4a877928c01ea56f89f1a8 (HEAD -> master)
Author: <REDACTED> 
Date:   Tue Apr 2 01:21:06 2024 +0200

    something something

$ hg log
changeset:   0:e8c3c975ba8e
tag:         tip
user:        <REDACTED>
date:        Tue Apr 02 01:07:28 2024 +0200
summary:     something something

$ fossil timeline
=== 2024-04-01 ===
23:21:06 [0bbb654d46] *CURRENT* something something (user: <REDACTED> tags: trunk)
23:19:34 [e054bc9935] initial empty check-in (user: <REDACTED> tags: trunk)
+++ no more data (2) +++
```
