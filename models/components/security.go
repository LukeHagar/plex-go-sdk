// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Security struct {
	AccessToken *string `security:"scheme,type=apiKey,subtype=header,name=X-Plex-Token"`
}

func (o *Security) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}
