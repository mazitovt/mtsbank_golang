package geocoding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mtsbank_golang/lesson5/distance/points"
	"net/http"
)

type WebGeocoder struct {
	client *http.Client
	url    string
	token  string
}

func (g WebGeocoder) Geocode(address string) (point points.PointOnSphere, err error) {
	panic("implement me")
}

func (g WebGeocoder) ReverseGeocode(point points.PointOnSphere) (data GeocodeData, err error) {
	jsonRequest, _ := json.Marshal(map[string]string{"lat": fmt.Sprintf("%f", point.Lat()), "lon": fmt.Sprintf("%f", point.Lon()), "count": "1"})
	req, _ := http.NewRequest("POST", g.url, bytes.NewBuffer(jsonRequest))
	req.Header.Add("Authorization", "Token "+g.token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	response, err := g.client.Do(req)

	if err != nil {
		return
	}

	//buf, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(buf))

	var res map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&res)
	suggestions := res["suggestions"].([]interface{})
	sugElem := suggestions[0].(map[string]interface{})
	data1 := sugElem["data"].(map[string]interface{})
	postalCode := data1["postal_code"].(string)
	fmt.Println("postalcode ", postalCode)

	return
}

/*
curl -X POST \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-H "Authorization: Token 11704dfd376beff80fa6b77d454026c9ddd025f0" \
-d '{ "lat": 55.878, "lon": 37.653 }' \
https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address
*/

func NewGeocoder() *WebGeocoder {
	return &WebGeocoder{client: &http.Client{}, url: "https://suggestions.dadata.ru/suggestions/api/4_1/rs/geolocate/address", token: "11704dfd376beff80fa6b77d454026c9ddd025f0"}
}
