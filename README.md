# cdropp
## ram cache dropper for linux
----


I recently decided to play gta5 and found that after about 40-50 minutes the game crashes with the error "Out-Of-Memory". Although availible ram still more than 5gb. After a while I realized that gta is overloading ram cache and kernel has no time to take availible memory from other applications which leads to crashes.

Since I'm a bit of a developer I decided to write a daemon which will clear the cache during the game.


# How to Install?
### 1) install go
### 2) clone repo
```fish
git clone https://github.com/ssleert/cdropp.go
```
### 3) make
```fish
make build
sudo make install
```

# How to uninstall?
```fish
sudo make uninstall
```

or

```fish
sudo rm /usr/local/bin/cdropp
```

# How to use?

### Check version.
```fish
~> cdropp -v
cdropp - 0.0.1
```

### Get help message.
```fish
~> cdropp -h
cdropp - simple deamon for dropping caches in ram
 -d --daemon  | start main loop
 -c --check   | check current ram usage
 -v --version | print program version
 -h --help    | print help message
```

### Check ram stats.
```fish
~> cdropp -c
ram info:
 total     - 15883 mb
 free      - 3364 mb
 availible - 4252 mb
```

### Start in deamon mode.
```fish
~> cdropp -d # with debug=1
2023/01/06 18:25:44 cache dropped | panic: 0
2023/01/06 18:25:44 cache checked | panic: 0
2023/01/06 18:25:44 cache dropped | panic: 0
2023/01/06 18:25:44 cache checked | panic: 0
2023/01/06 18:25:44 cache dropped | panic: 0
2023/01/06 18:25:44 cache checked | panic: 0
```

# How to configure?
By default cdropp reads `/etc/cdropp/conf.ini`.

but config file location can be redefined with `CDROPP_CONF_PATH` environment variable
```fish
~> CDROPP_CONF_PATH=./examples/conf.ini ./cdropp -d
```

### Default values of /etc/cdropp/conf.ini
```ini
[cdropp]

# minimal free ram trigger in mb
trigger = 256

# time between checks in ms
timer = 12000

# strength of cache drops 1,2,3
strength = 1

# debug 0,1
debug = 0
```

<br>

-----
### have a nice day!
cdropp made with go and love by sfome)