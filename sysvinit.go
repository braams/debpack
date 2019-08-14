package debpack

//inspire by docker init script

var Sysvinit = `#! /bin/sh
### BEGIN INIT INFO
# Provides:          {{.Name}}
# Required-Start:    $remote_fs $network
# Required-Stop:     $remote_fs $network
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: {{.Description}}
# Description:       {{.Description}}
### END INIT INFO

PATH="/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin"
NAME="{{.Name}}"
DAEMON="/usr/sbin/$NAME"
DAEMON_OPTS="-config /etc/$NAME/$NAME.conf"
DESC="{{.Description}}"
PIDFILE="/var/run/$NAME.pid"
BINARY="/usr/sbin/$NAME"
USER="{{.User}}"
LOGFILE="/var/log/{{.Name}}/{{.Name}}.log"

test -x $DAEMON || exit 0

. /lib/lsb/init-functions

# Include defaults if available
if [ -f /etc/default/$NAME ] ; then
. /etc/default/$NAME
fi

start() {
start-stop-daemon --start --quiet --pidfile $PIDFILE --oknodo \
--background --exec $BINARY --startas $DAEMON \
--make-pidfile --chuid $USER --no-close \
-- $DAEMON_OPTS >> $LOGFILE 2>&1 || return 1
return 0
}

stop() {
start-stop-daemon --stop --quiet --pidfile $PIDFILE \
--oknodo --exec $BINARY || return 1
rm -f $PIDFILE
return 0
}

case "$1" in
start)
log_begin_msg "Starting $DESC: $NAME"
start
log_end_msg $?
;;
stop)
log_begin_msg "Stopping $DESC: $NAME"
stop
log_end_msg $?
;;
#reload)
#
#   If the daemon can reload its config files on the fly
#   for example by sending it SIGHUP, do it here.
#
#   If the daemon responds to changes in its config file
#   directly anyway, make this a do-nothing entry.
#
# echo "Reloading $DESC configuration files."
# start-stop-daemon --stop --signal 1 --quiet --pidfile \
#   /var/run/$NAME.pid --exec $DAEMON
#;;
restart|force-reload)
#
#   If the "reload" option is implemented, move the "force-reload"
#   option to the "reload" entry above. If not, "force-reload" is
#   just the same as "restart".
#
log_begin_msg "Restarting $DESC: $NAME"
stop && sleep 1 && start
log_end_msg $?
;;
status)
status_of_proc -p "$PIDFILE" "$DAEMON" "$NAME" && exit 0 || exit $?
;;
*)
# echo "Usage: $0 {start|stop|restart|reload|force-reload}" >&2
echo "Usage: $0 {start|stop|restart|status}" >&2
exit 1
;;
esac

exit 0
`

var Sysvinit2 = `#! /bin/sh
### BEGIN INIT INFO
# Provides:          {{.Name}}
# Required-Start:    $remote_fs $network
# Required-Stop:     $remote_fs $network
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: {{.Description}}
# Description:       {{.Description}}
### END INIT INFO

PATH="/usr/local/sbin:/usr/local/bin:/sbin:/bin:/usr/sbin:/usr/bin"
NAME="{{.Name}}"
DAEMON="/usr/sbin/$NAME"
DAEMON_OPTS="-config /etc/$NAME.conf"
DESC="{{.Description}}"
PIDFILE="/var/run/$NAME.pid"
BINARY="/usr/sbin/$NAME"
USER="{.User}"
LOGFILE="/var/log/{{.Name}}.log"

test -x $DAEMON || exit 0

. /lib/lsb/init-functions

# Include defaults if available
if [ -f /etc/default/$NAME ] ; then
. /etc/default/$NAME
fi

start() {
start-stop-daemon --start --quiet --pidfile $PIDFILE --oknodo \
--background --exec $BINARY --startas $DAEMON \
--make-pidfile --chuid $USER --no-close \
-- $DAEMON_OPTS >> $LOGFILE 2>&1 || return 1
return 0
}

stop() {
start-stop-daemon --stop --quiet --pidfile $PIDFILE \
--oknodo --exec $BINARY || return 1
rm -f $PIDFILE
return 0
}

case "$1" in
start)
log_begin_msg "Starting $DESC: $NAME"
start
log_end_msg $?
;;
stop)
log_begin_msg "Stopping $DESC: $NAME"
stop
log_end_msg $?
;;
restart|force-reload)
log_begin_msg "Restarting $DESC: $NAME"
stop && sleep 1 && start
log_end_msg $?
;;
status)
status_of_proc -p "$PIDFILE" "$DAEMON" "$NAME" && exit 0 || exit $?
;;
*)
echo "Usage: $0 {start|stop|restart|status}" >&2
exit 1
;;
esac

exit 0
`
