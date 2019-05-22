# pinup

Upgrade your docker pinned dependencies

Currently it is going over all the relevant lines

```
Go pkg manager... no check
Line 42 can use: latest
Alpine pkg manager check
Line 45 can use: latest
Alpine pkg manager check
```

Since I have this info, I can:

- Search dockerhub for a newer images (need to think of a upgrade policy for versions)
 - Search alpines package manager for versions

This will involve possibly making a GO api for each repo