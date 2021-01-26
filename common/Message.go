/*
Zahhak2, a Golang multiplayer console game.
Copyright (C) 2016 Aryo Pehlewan feylikurds@gmail.com
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package common

import (
	"encoding/json"
)

type Message struct {
	Class       string
	ID          string
	Action      string
	Params      map[string]string
	MultiParams map[string][]string
}

func NewMessage(class, id, action string) *Message {
	return &Message{
		Class:       class,
		ID:          id,
		Action:      action,
		Params:      map[string]string{},
		MultiParams: map[string][]string{},
	}
}

func (m *Message) JSON() ([]byte, error) {
	return json.Marshal(m)
}