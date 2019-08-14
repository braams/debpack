package debpack

var Postinst = `#!/bin/bash
echo "Postinst"
chmod 755 {{.InitFileName}}
mkdir -p /var/log/{{.Name}}
touch {{.LogFileName}}
chown -R {{.User}}:{{.User}} /var/log/{{.Name}}
update-rc.d {{.Name}} defaults
`

var Preinst = `#!/bin/bash
echo "Preinst"
if [ -z "$(getent passwd {{.User}})" ]; then
useradd --system --shell "/bin/false" --user-group {{.User}}
fi`

var Prerm = `#!/bin/bash
echo "Prerm"
`

var Postrm = `#!/bin/bash
echo "Postrm"
update-rc.d -f {{.Name}} remove
`
