# GetTokenByPinIDResponseBody

Bad Request response when the X-Plex-Client-Identifier is missing


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Errors`                                                                             | [][sdkerrors.GetTokenByPinIDErrors](../../models/sdkerrors/gettokenbypiniderrors.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `RawResponse`                                                                        | [*http.Response](https://pkg.go.dev/net/http#Response)                               | :heavy_minus_sign:                                                                   | Raw HTTP response; suitable for custom response parsing                              |