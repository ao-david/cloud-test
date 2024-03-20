/*
LICENSE
  Copyright (C) 2024 the Australian Ocean Lab (AusOcean)

  This file is part of Ocean Cron. Ocean Cron is free software: you can
  redistribute it and/or modify it under the terms of the GNU
  General Public License as published by the Free Software
  Foundation, either version 3 of the License, or (at your option)
  any later version.

  Ocean Cron is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with NetReceiver in gpl.txt. If not, see
  <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"math"
	"testing"

	"bitbucket.org/ausocean/iotsvc/iotds"
)

var cronSpecTests = []struct {
	cron     *iotds.Cron
	lat, lon float64
	want     string
	wantErr  error
}{
	{
		cron:    &iotds.Cron{},
		want:    "",
		wantErr: nil,
	},
	{
		cron:    &iotds.Cron{Enabled: true},
		want:    "",
		wantErr: errNoTimeSpec,
	},
	{
		cron: &iotds.Cron{TOD: "@sunrise"},
		lat:  1, lon: 1,
		want:    "",
		wantErr: nil,
	},
	{
		cron: &iotds.Cron{TOD: "@sunrise", Enabled: true},
		lat:  math.NaN(), lon: math.NaN(),
		want:    "",
		wantErr: errNoLocation,
	},
	{
		cron: &iotds.Cron{TOD: "@sunrise", Enabled: true},
		lat:  1, lon: 1,
		want:    "@sunrise 1 1",
		wantErr: nil,
	},
	{
		cron: &iotds.Cron{TOD: "@sunrise+1h", Enabled: true},
		lat:  1, lon: 1,
		want:    "@sunrise+1h 1 1",
		wantErr: nil,
	},
	{
		cron: &iotds.Cron{TOD: "@noon", Enabled: true},
		lat:  1, lon: 1,
		want:    "@noon 1 1",
		wantErr: nil,
	},
	{
		cron: &iotds.Cron{TOD: "@midnight", Enabled: true},
		lat:  1, lon: 1,
		want:    "@midnight",
		wantErr: nil,
	},
}

func TestCronSpec(t *testing.T) {
	for _, test := range cronSpecTests {
		got, err := cronSpec(test.cron, test.lat, test.lon)
		if fmt.Sprint(err) != fmt.Sprint(test.wantErr) {
			t.Errorf("unexpected error: got:%v want:%v", err, test.wantErr)
		}
		if err != nil {
			continue
		}
		if got != test.want {
			t.Errorf("unexpected cron spec: got:%s want:%s", got, test.want)
		}
	}
}
