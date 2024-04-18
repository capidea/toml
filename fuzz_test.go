package toml

//import 
import (
	"bytes"
	"testing"
)

func FuzzDecode(f *testing.F) {
	buf := make([]byte, 0, 2048)

	f.Add(`
# This is an example TOML document which shows most of its features.

desc = """
An example TOML document. \
"""

# Array with integers and floats in the various allowed formats.
integers = [42, 0x42, 0o42, 0b0110]
floats   = [1.42, 1e-02]

# Array with supported datetime formats.
times = [
	2021-11-09T15:16:17+01:00,  # datetime with timezone.
	2021-11-09T15:16:17Z,       # UTC datetime.
	2021-11-09T15:16:17,        # local datetime.
	2021-11-09,                 # local date.
	15:16:17,                   # local time.
]

# Durations.
duration = ["4m49s", "8m03s", "1231h15m55s"]

# Table with inline tables.
distros = [
	{name = "Arch Linux", packages = "pacman"},
	{name = "Void Linux", packages = "xbps"},
	{name = "Debian",     packages = "apt"},
]

# Create new table; note the "servers" table is created implicitly.
[servers.alpha]
	# You can indent as you please, tabs or spaces.
	ip        = '10.0.0.1'
	hostname  = 'server1'
	enabled   = false
[servers.beta]
	ip        = '10.0.0.2'
	hostname  = 'server2'
	enabled   = true

# Start a new table array; note that the "characters" table is created implicitly.
[[characters.star-trek]]
	name = "James Kirk"
	rank = "Captain\u0012 \t"
[[characters.star-trek]]
	name = "Spock"
	rank = "Science officer"

[undecoded] # To show the MetaData.Undecoded() feature.
	key = "This table intentionally left undecoded"
`)
	f.Fuzz(func(t *testing.T, file string) {
		var m map[string]any
		_, err := Decode(file, &m)
		if err != nil {
			t.Skip()
		}

		NewEncoder(bytes.NewBuffer(buf)).Encode(m)

		// TODO: should check if the output is equal to the input, too, but some
		// TODO: should check if the output is equal to the input, too, but some
		// information is lost when encoding.
	})
}
