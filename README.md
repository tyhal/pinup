# pinup

Github has the functionality in this tool: <https://dependabot.com/>

However it doesn't upgrade the dependencies installed on RUN lines and doesn't run on Gitlab :(

* * *

## Run Me

```bash
    script/bootstrap
    script/test
```

* * *

Upgrade your docker pinned dependencies

Currently it is going over all the relevant lines

    Go pkg manager... no check
    ...
    Line 42 can use: alpine:3.9.4
    Alpine pkg manager check
    ...
    Line 45 can use: alpine:3.9.4
    Alpine pkg manager check

Since I have this info, I can:

-   Search dockerhub for a newer images (need to think of a upgrade policy for versions)
-   Search alpines package manager for versions

This will involve possibly making a GO api for each repo
