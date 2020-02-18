package config

import "time"

const MainAppPort = 7777
const DBHostname = "postgres"

const Database = "test"
const User = "postgres"
const Password = "postgres"

// Time allowed to write a message to the peer.
const WriteWait = 10 * time.Second

// Time allowed to read the next pong message from the peer.
const PongWait = 50 * time.Second

// Send pings to peer with this period. Must be less than pongWait.
const PingPeriod = (PongWait * 9) / 10
const MaxMessageSize = 512

const PublicDir = "/static"