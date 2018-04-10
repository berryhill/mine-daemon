# Mine Daemon

## Up daemon with mock mine

*Tested with Go 10.0*

In one terminal window, start the mine daemon
```
$ go get github.com/berryhill/mine-daemon
$ cd $GOPATH/src/github.com/berryhill/mine-daemon
$ go get 
$ go run main.go -addr=amqp://ltvdoacc:urQws_KDYLQbcK0mOidy48snnJQMsr7Z@wombat.rmq.cloudamqp.com/ltvdoacc
```

In another terminal window, start the mock mine
``` 
$ cd $GOPATH/src/github.com/berryhill/mine-daemon
$ ./mock-mine
# if using linux with amd64
$ ./mock-mine-linux-amd64
```

Go to UI and see the associated hashrates updating [here](http://35.225.59.241/#/rigs)

*let mock mine run for about 30 seconds before hashrates start changing*

Also find the redis JSON state [here](http://35.226.250.99:5051/)
