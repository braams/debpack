package debpack

var Logrotate = `{{.LogFileName}} {
	su {{.User}} {{.User}}
	daily
	missingok
	rotate 7
	notifempty
	nocreate
	copytruncate
}
`
