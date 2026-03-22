package pokeapi

import (
	// "encoding/json"
	// "net/http"
)

// https://pokeapi.co/api/v2

$ curl --request GET https://pokeapi.co/api/v2/location/ | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  1551    0  1551    0     0   7952      0 --:--:-- --:--:-- --:--:--  7953
type pokeapiResponse struct{
  Count int "json:`count`"
  Next string "json:`next`"
  Previous string "json:`previous`"
  Results []struct{
  	Name string "json:`name`"
  	Url  string "json:`url`"
  } "json:`results`"
}