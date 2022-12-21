package main

var ACL = new(Credentials)

type Credentials struct {
	Username *string
	Password *string
}

func (provider Credentials) Hydrate() {
	if username, valid := Environment("AWS_REDIS_USERNAME"); valid || !(valid) {
		ACL.Username = Declare("AWS_REDIS_USERNAME", username, "default")
	}

	if password, valid := Environment("AWS_REDIS_PASSWORD"); valid || !(valid) {
		ACL.Password = Declare("AWS_REDIS_PASSWORD", password, "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")
	}
}
