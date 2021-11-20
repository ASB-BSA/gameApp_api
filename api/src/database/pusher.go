package database

import "github.com/pusher/pusher-http-go"

var PusherClient = pusher.Client{
	AppID:   "1300788",
	Key:     "fcbd0fe292a544abd4dd",
	Secret:  "5fecd811e3a105208a5b",
	Cluster: "ap3",
	Secure:  true,
}
