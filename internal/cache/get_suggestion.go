/*
  Copyright (C) 2019 - 2020 MWSOFT
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
package cache

import (
	"github.com/go-redis/redis"
	"github.com/superhero-match/superhero-get-match/internal/cache/model"
)

// GetSuggestions fetches suggestions from cache.
func (c *Cache) GetMatch(key string) (*model.Superhero, error) {
	res, err := c.Redis.Get(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	var suggestion model.Superhero

	if err := suggestion.UnmarshalBinary([]byte(res)); err != nil {
		return nil, err
	}

	return &suggestion, nil
}