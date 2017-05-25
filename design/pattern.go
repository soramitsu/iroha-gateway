package design

const (
	patternBase64      = `^[0-9a-zA-Z+/=]+$`
	patternTimestamp   = `^[0-9]{1,18}$`
	patternURLEncoded  = `^[0-9a-zA-Z-_.~]+$`
	patternAccountUUID = `^[0-9a-f-]{32}$`
	patternDomainURI   = `^(?:[0-9a-zA-Z-._~]|%[0-9][0-9])+$` // following: RFC3986
	patternAccountUsername    = `^[a-zA-Z0-9\_\-\.]{4,32}$`
)
